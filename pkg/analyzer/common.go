package analyzer

import (
	"go/types"
	"log"

	"golang.org/x/tools/go/analysis/passes/buildssa"
	"golang.org/x/tools/go/ssa"
)

type closeableValue struct {
	value *ssa.Value
	instr ssa.Instruction
}

// getCloseableValues returns a list of values that should be closed.
// Given a list  of types, the function loops over the blocks instructions. If
// a value of a target type is found, it is added to the list of values to be
// ensured is closed.
func getClosableValues(b *ssa.BasicBlock, i int, targetTypes []any) []closeableValue {
	closeableValues := []closeableValue{}

	instr := b.Instrs[i]

	// if instruction is not a call, then it can't return a value of the target type
	call, ok := instr.(*ssa.Call)
	if !ok {
		return closeableValues
	}

	// check the signature of the function for a value of the target types
	signature := call.Call.Signature()
	results := signature.Results()
	for i := 0; i < results.Len(); i++ {
		v := results.At(i)
		varType := v.Type()

		// iterate over the target types and check if result type is the same
		for _, targetType := range targetTypes {
			var tt types.Type

			// if pointer or named type, then check if it is the same as the target type
			switch t := targetType.(type) {
			case *types.Pointer:
				tt = t
			case *types.Named:
				tt = t
			default:
				// not a pointer or named type, skip
				continue
			}

			// if the result type is not the same as the current target, skip
			if !types.Identical(varType, tt) {
				continue
			}

			// check the referers of the call
			for _, cRef := range *call.Referrers() {
				switch instr := cRef.(type) {
				case *ssa.Call:
					// if reference is a call, confirm the first argument is the same as the
					// target type
					if len(instr.Call.Args) >= 1 && types.Identical(instr.Call.Args[0].Type(), tt) {
						// add the value to the list of values to be closed
						log.Printf("found closeable call: %s %s", instr.String(), instr.Type().String())
						closeableValues = append(closeableValues, closeableValue{
							value: &instr.Call.Args[0],
							instr: call,
						})
					}
				case ssa.Value: // is this a case that can happen? what is it's use case?
					if types.Identical(instr.Type(), tt) {
						log.Printf("found closeable value: %s %s", instr.String(), instr.Type().String())
						closeableValues = append(closeableValues, closeableValue{
							value: &instr,
							instr: call,
						})
					}
				}
			}
		}
	}

	return closeableValues
}

func getTargetTypes(pssa *buildssa.SSA, targetPackages []string) []any {
	targets := []any{}

	for _, sqlPkg := range targetPackages {
		pkg := pssa.Pkg.Prog.ImportedPackage(sqlPkg)
		if pkg == nil {
			// the SQL package being checked isn't imported
			continue
		}

		rowsPtrType := getTypePointerFromName(pkg, rowsName)
		if rowsPtrType != nil {
			targets = append(targets, rowsPtrType)
		}

		rowsType := getTypeFromName(pkg, rowsName)
		if rowsType != nil {
			targets = append(targets, rowsType)
		}

		stmtType := getTypePointerFromName(pkg, stmtName)
		if stmtType != nil {
			targets = append(targets, stmtType)
		}

		namedStmtType := getTypePointerFromName(pkg, namedStmtName)
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

func isTargetType(t types.Type, targetTypes []any) bool {
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
