package analyzer

import (
	"go/types"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/buildssa"
	"golang.org/x/tools/go/ssa"
)

const (
	rowsName    = "Rows"
	stmtName    = "Stmt"
	closeMethod = "Close"
)

var (
	sqlPackages = []string{
		"database/sql",
	}
)

func NewAnalyzer() *analysis.Analyzer {
	return &analysis.Analyzer{
		Name: "sqlclosecheck",
		Doc:  "Checks that sql.Rows and sql.Stmt are closed.",
		Run:  NewRun(sqlPackages),
		Requires: []*analysis.Analyzer{
			buildssa.Analyzer,
		},
	}
}

func NewRun(sqlPkgs []string) func(pass *analysis.Pass) (interface{}, error) {
	return func(pass *analysis.Pass) (interface{}, error) {
		for _, sqlPkg := range sqlPkgs {
			_, err := run(pass, sqlPkg)
			if err != nil {
				return nil, err
			}
		}

		return nil, nil
	}
}

func run(pass *analysis.Pass, sqlPkg string) (interface{}, error) {
	pssa := pass.ResultOf[buildssa.Analyzer].(*buildssa.SSA)

	// Check for our target types and build list
	targetTypes := getTargetTypes(pssa, sqlPkg)

	// If non of the types are found, skip
	if len(targetTypes) == 0 {
		return nil, nil
	}

	funcs := pssa.SrcFuncs
	for _, f := range funcs {
		for _, b := range f.Blocks {
			for i := range b.Instrs {
				// Check if instruction is call that returns a target type
				targetValues := getTargetTypesValues(b, i, targetTypes)
				if len(targetValues) == 0 {
					continue
				}

				for _, targetValue := range targetValues {
					//pass.Reportf(targetValues.Pos(), "found values of the target type")

					refs := (*targetValue.value).Referrers()
					if len(*refs) == 0 {
						continue
					}

					isClosed := checkClosed(refs, targetTypes)
					if !isClosed {
						pass.Reportf((targetValue.instr).Pos(), "Rows/Stmt was not closed")
					}

					checkDeferred(pass, refs, targetTypes, false)					
				}
			}
		}
	}

	return nil, nil
}

func getTargetTypes(pssa *buildssa.SSA, sqlPkg string) []*types.Pointer {
	targets := []*types.Pointer{}

	pkg := pssa.Pkg.Prog.ImportedPackage(sqlPkg)
	if pkg == nil {
		// the SQL package being checked isn't imported
		return targets
	}

	rowsType := getTypePointerFromName(pkg, rowsName)
	if rowsType != nil {
		targets = append(targets, rowsType)
	}

	stmtType := getTypePointerFromName(pkg, stmtName)
	if stmtType != nil {
		targets = append(targets, stmtType)
	}

	return targets
}

func getTypePointerFromName(pkg *ssa.Package, name string) *types.Pointer {
	pkgType := pkg.Type(name)
	if pkgType == nil {
		// this package does not use Rows/Stmt
		return nil
	}

	obj := pkgType.Object()
	named, ok := obj.Type().(*types.Named)
	if !ok {
		return nil
	}

	return types.NewPointer(named)
}

type targetValue struct {
	value *ssa.Value
	instr ssa.Instruction
}

func getTargetTypesValues(b *ssa.BasicBlock, i int, targetTypes []*types.Pointer) []targetValue {
	targetValues := []targetValue{}

	instr := b.Instrs[i]
	call, ok := instr.(*ssa.Call)
	if !ok {
		return targetValues
	}

	signature := call.Call.Signature()
	results := signature.Results()
	for i := 0; i < results.Len(); i++ {
		for _, targetType := range targetTypes {
			v := results.At(i)
			varType := v.Type()
			if types.Identical(varType, targetType) {
				for _, cRef := range *call.Referrers() {
					switch instr := cRef.(type) {
					case *ssa.Call:
						if len(instr.Call.Args) == 1 && types.Identical(instr.Call.Args[0].Type(), targetType) {
							targetValues = append(targetValues, targetValue{
								value: &instr.Call.Args[0],
								instr: call,
							})
						}
					case ssa.Value:
						if types.Identical(instr.Type(), targetType) {
							targetValues = append(targetValues, targetValue{
								value: &instr,
								instr: call,
							})
						}
					}
				}
			}
		}
	}

	return targetValues
}

func checkClosed(refs *[]ssa.Instruction, targetTypes []*types.Pointer) bool {
	isClosed := false

	for _, refs := range *refs {
		if isCloseCall(refs, targetTypes) {
			isClosed = true
		}
	}

	return isClosed
}

func isCloseCall(instr ssa.Instruction, targetTypes []*types.Pointer) bool {
	switch instr := instr.(type) {
	case *ssa.Defer:
		if instr.Call.Value != nil && instr.Call.Value.Name() == closeMethod {
			return true
		}
	case *ssa.Call:
		if instr.Call.Value != nil && instr.Call.Value.Name() == closeMethod {
			return true
		}
	case *ssa.Store:
		if len(*instr.Addr.Referrers()) == 0 {
			return false
		}

		for _, aRef := range *instr.Addr.Referrers() {
			if c, ok := aRef.(*ssa.MakeClosure); ok {
				f := c.Fn.(*ssa.Function)

				for _, b := range f.Blocks {
					for _, innerInstr := range b.Instrs {
						switch innerInstr := innerInstr.(type) {
						case *ssa.UnOp:
							instrType := innerInstr.Type()
							for _, targetType := range targetTypes {
								if types.Identical(instrType, targetType) {
									isClosed := checkClosed(innerInstr.Referrers(), targetTypes)
									if isClosed {
										return true
									}
								}
							}
						}
					}
				}
			}
		}
	}

	return false
}

func checkDeferred(pass *analysis.Pass, instrs *[]ssa.Instruction, targetTypes []*types.Pointer, inDefer bool) {
	for _, instr := range *instrs {
		switch instr := instr.(type) {
		case *ssa.Defer:
			if instr.Call.Value != nil && instr.Call.Value.Name() == closeMethod {
				return
			}
		case *ssa.Call:
			if instr.Call.Value != nil && instr.Call.Value.Name() == closeMethod {
				if !inDefer {
					pass.Reportf(instr.Pos(), "Close should use defer")
				}
	
				return
			}
		case *ssa.Store:
			if len(*instr.Addr.Referrers()) == 0 {
				return
			}
	
			for _, aRef := range *instr.Addr.Referrers() {
				if c, ok := aRef.(*ssa.MakeClosure); ok {
					f := c.Fn.(*ssa.Function)
	
					for _, b := range f.Blocks {
						for _, innerInstr := range b.Instrs {
							switch innerInstr := innerInstr.(type) {
							case *ssa.UnOp:
								instrType := innerInstr.Type()
								for _, targetType := range targetTypes {
									if types.Identical(instrType, targetType) {
										checkDeferred(pass, innerInstr.Referrers(), targetTypes, true)
									}
								}
							}
						}
					}
				}
			}
		}
	}
}
