package analyzer

import (
	"flag"
	"go/types"
	"log"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/buildssa"
	"golang.org/x/tools/go/ssa"
)

type action uint8

const (
	actionUnhandled action = iota
	actionHandled
	actionReturned
	actionPassed
	actionClosed
	actionUnvaluedCall
	actionUnvaluedDefer
	actionNoOp
)

// deferOnlyAnalyzer is an analyzer that requires a close to be deferred immediately after creation
type deferOnlyAnalyzer struct{}

// NewDeferOnlyAnalyzer returns a new deferOnlyAnalyzer
func NewDeferOnlyAnalyzer() *analysis.Analyzer {
	analyzer := &deferOnlyAnalyzer{}
	flags := flag.NewFlagSet("deferOnlyAnalyzer", flag.ExitOnError)
	return newAnalyzer(analyzer.Run, flags)
}

// Run implements the main analysis pass
//
// # TODO write description of how this analyzer works
//
// TODO is there anyway we can cut down on the number of iterations?
func (a *deferOnlyAnalyzer) Run(pass *analysis.Pass) (interface{}, error) {
	pssa, ok := pass.ResultOf[buildssa.Analyzer].(*buildssa.SSA)
	if !ok {
		return nil, nil
	}

	typesNeedingClosed := findTypesNeedingClosed(pssa, sqlPackages)
	if len(typesNeedingClosed) == 0 {
		return nil, nil
	}

	// Iterate over all functions, blocks, and instructions looking for calls to target methods
	funcs := pssa.SrcFuncs
	for _, f := range funcs {
		for _, b := range f.Blocks {
			for i := range b.Instrs {
				// Check if each instruction is a call that returns a value that needs closed
				resultsNeedingClosed := findResultsNeedingClosed(b, i, typesNeedingClosed)
				if len(resultsNeedingClosed) == 0 {
					continue
				}

				// For each found target check if they are closed and deferred
				for _, result := range resultsNeedingClosed {
					refs := result.value.Referrers()

					// isClosed := checkClosed(refs, typesNeedingClosed)
					// if !isClosed {
					// 	pass.Reportf((result.instr).Pos(), "Rows/Stmt/NamedStmt was not closed")
					// }

					reportImproperHandling(pass, result.instr, refs, typesNeedingClosed, false)
				}
			}
		}
	}

	return nil, nil
}

func reportImproperHandling(
	pass *analysis.Pass,
	instr ssa.Instruction,
	refs *[]ssa.Instruction,
	targetTypes []types.Type,
	inDefer bool,
) {
	for _, refInstr := range *refs {
		log.Printf("instr: %v", refInstr)

		switch instr := refInstr.(type) {
		// check for defer (.e.g. `defer rows.Close()`)
		case *ssa.Defer:
			log.Printf("instr.Call: %v", instr.Call)

			if instr.Call.Value != nil && instr.Call.Value.Name() == MethodCloseName {
				return
			}

			if instr.Call.Method != nil && instr.Call.Method.Name() == MethodCloseName {
				return
			}
		// check for function call (.e.g. `rows.Close()`)
		case *ssa.Call:
			if instr.Call.Value != nil && instr.Call.Value.Name() == MethodCloseName {
				// if we are not inside of a defer, then we should report that a defer wasn't used
				if !inDefer {
					pass.Reportf(instr.Pos(), "Close should use defer")
					return
				}

				// in a defer and was closed, so we are good
				return
			}
			// case *ssa.Store:
			// 	if len(*instr.Addr.Referrers()) == 0 {
			// 		return
			// 	}

			// 	for _, aRef := range *instr.Addr.Referrers() {
			// 		if c, ok := aRef.(*ssa.MakeClosure); ok {
			// 			if f, ok := c.Fn.(*ssa.Function); ok {
			// 				for _, b := range f.Blocks {
			// 					reportImproperHandling(pass, instr, &b.Instrs, targetTypes, true)
			// 					return
			// 				}
			// 			}
			// 		}
			// 	}
			// case *ssa.UnOp:
			// 	instrType := instr.Type()
			// 	for _, targetType := range targetTypes {
			// 		var tt types.Type

			// 		switch t := targetType.(type) {
			// 		case *types.Pointer:
			// 			tt = t
			// 		case *types.Named:
			// 			tt = t
			// 		default:
			// 			continue
			// 		}

			// 		if types.Identical(instrType, tt) {
			// 			reportImproperHandling(pass, instr, instr.Referrers(), targetTypes, inDefer)
			// 			return
			// 		}
			// 	}
			// case *ssa.FieldAddr:
			// 	reportImproperHandling(pass, instr, instr.Referrers(), targetTypes, inDefer)
			// 	return
		}
	}

	log.Printf("not closed: %v", instr)
	pass.Reportf(instr.Pos(), "Rows/Stmt/NamedStmt was not closed")
}

func checkClosed(refs *[]ssa.Instruction, targetTypes []types.Type) bool {
	numInstrs := len(*refs)
	for idx, ref := range *refs {
		action := getAction(ref, targetTypes)
		switch action {
		case actionClosed, actionReturned, actionHandled:
			return true
		case actionPassed:
			// Passed and not used after
			if numInstrs == idx+1 {
				return true
			}
		}
	}

	return false
}

func getAction(instr ssa.Instruction, targetTypes []types.Type) action {
	switch instr := instr.(type) {
	case *ssa.Defer:
		if instr.Call.Value != nil {
			name := instr.Call.Value.Name()
			if name == MethodCloseName {
				return actionClosed
			}
		}

		if instr.Call.Method != nil {
			name := instr.Call.Method.Name()
			if name == MethodCloseName {
				return actionClosed
			}
		}

		return actionUnvaluedDefer
	case *ssa.Call:
		if instr.Call.Value == nil {
			return actionUnvaluedCall
		}

		isTarget := false
		staticCallee := instr.Call.StaticCallee()
		if staticCallee != nil {
			receiver := instr.Call.StaticCallee().Signature.Recv()
			if receiver != nil {
				isTarget = isTargetType(receiver.Type(), targetTypes)
			}
		}

		name := instr.Call.Value.Name()
		if isTarget && name == MethodCloseName {
			return actionClosed
		}

		if !isTarget {
			return actionPassed
		}
	case *ssa.Phi:
		return actionPassed
	case *ssa.MakeInterface:
		return actionPassed
	case *ssa.Store:
		// A Row/Stmt is stored in a struct, which may be closed later
		// by a different flow.
		if _, ok := instr.Addr.(*ssa.FieldAddr); ok {
			return actionReturned
		}

		if len(*instr.Addr.Referrers()) == 0 {
			return actionNoOp
		}

		for _, aRef := range *instr.Addr.Referrers() {
			if c, ok := aRef.(*ssa.MakeClosure); ok {
				if f, ok := c.Fn.(*ssa.Function); ok {
					for _, b := range f.Blocks {
						if checkClosed(&b.Instrs, targetTypes) {
							return actionHandled
						}
					}
				}
			}
		}
	case *ssa.UnOp:
		instrType := instr.Type()
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

			if types.Identical(instrType, tt) {
				if checkClosed(instr.Referrers(), targetTypes) {
					return actionHandled
				}
			}
		}
	case *ssa.FieldAddr:
		if checkClosed(instr.Referrers(), targetTypes) {
			return actionHandled
		}
	case *ssa.Return:
		return actionReturned
	}

	return actionUnhandled
}
