package analyzer

import (
	"go/types"

	"golang.org/x/tools/go/analysis/passes/buildssa"
	"golang.org/x/tools/go/ssa"
)

const (
	// types of interest
	TypeRowsName      = "Rows"
	TypeStmtName      = "Stmt"
	TypeNamedStmtName = "NamedStmt"

	// methods of interest
	MethodCloseName = "Close"
	MethodNextName  = "Next"
	MethodErrName   = "Err"
)

var (
	sqlPackages = []string{
		"database/sql",
		"github.com/jmoiron/sqlx",
		"github.com/jackc/pgx/v5",
		"github.com/jackc/pgx/v5/pgxpool",
	}
)

// findTypesNeedingClosed returns a list of types that we will analayze
// List is `package` * `method` * `type (pointer, non-pointer)`. This is brute
// force and may produce pakage+method+type combinations that do not exist
//
// TODO append only combinations that exist
// TODO think about if this list should be abstracted or enriched
func findTypesNeedingClosed(pssa *buildssa.SSA, targetPackages []string) []types.Type {
	targets := []types.Type{}

	for _, sqlPkg := range targetPackages {
		pkg := pssa.Pkg.Prog.ImportedPackage(sqlPkg)
		if pkg == nil {
			// the SQL package being checked isn't imported
			continue
		}

		rowsPtrType := getTypePointerFromName(pkg, TypeRowsName)
		if rowsPtrType != nil {
			targets = append(targets, rowsPtrType)
		}

		rowsType := getTypeFromName(pkg, TypeRowsName)
		if rowsType != nil {
			targets = append(targets, rowsType)
		}

		stmtType := getTypePointerFromName(pkg, TypeStmtName)
		if stmtType != nil {
			targets = append(targets, stmtType)
		}

		namedStmtType := getTypePointerFromName(pkg, TypeNamedStmtName)
		if namedStmtType != nil {
			targets = append(targets, namedStmtType)
		}
	}

	return targets
}

func getTypePointerFromName(pkg *ssa.Package, name string) *types.Pointer {
	pkgType := pkg.Type(name)
	if pkgType == nil {
		// this package does not use Rows/Stmt/NamedStmt
		return nil
	}

	obj := pkgType.Object()
	named, ok := obj.Type().(*types.Named)
	if !ok {
		return nil
	}

	return types.NewPointer(named)
}

func getTypeFromName(pkg *ssa.Package, name string) *types.Named {
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

	return named
}

func isTargetType(t types.Type, targetTypes []types.Type) bool {
	for _, targetType := range targetTypes {
		switch tt := targetType.(type) {
		case *types.Pointer:
			if types.Identical(t, tt) {
				return true
			}
		case *types.Named:
			if types.Identical(t, tt) {
				return true
			}
		}
	}

	return false
}

type targetValue struct {
	instr ssa.Instruction
	value ssa.Value
}

// findResultsNeedingClosed returns a list of values that are of the target types
func findResultsNeedingClosed(b *ssa.BasicBlock, i int, targetTypes []types.Type) []targetValue {
	targetValues := []targetValue{}

	instr := b.Instrs[i]
	call, ok := instr.(*ssa.Call)
	// we are only interested in calls
	if !ok {
		return targetValues
	}

	signature := call.Call.Signature()
	results := signature.Results()
	// iterate through the return values of the call
	for i := 0; i < results.Len(); i++ {
		v := results.At(i)
		varType := v.Type()

		// iterate through the target types
		for _, targetType := range targetTypes {
			var tt types.Type

			switch t := targetType.(type) {
			case *types.Pointer:
				tt = t
			case *types.Named:
				tt = t
			default:
				continue
			}

			if !types.Identical(varType, tt) {
				continue
			}

			for _, cRef := range *call.Referrers() {
				switch instr := cRef.(type) {
				case *ssa.Call:
					if len(instr.Call.Args) >= 1 && types.Identical(instr.Call.Args[0].Type(), tt) {
						targetValues = append(targetValues, targetValue{
							value: instr.Call.Args[0],
							instr: call,
						})
					}
				case ssa.Value:
					if types.Identical(instr.Type(), tt) {
						targetValues = append(targetValues, targetValue{
							value: instr,
							instr: call,
						})
					}
				}
			}
		}
	}

	return targetValues
}
