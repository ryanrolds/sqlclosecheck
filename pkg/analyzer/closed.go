package analyzer

import (
	"flag"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/buildssa"
)

type closedAnalyzer struct{}

func NewClosedAnalyzer() *analysis.Analyzer {
	analyzer := &closedAnalyzer{}
	flags := flag.NewFlagSet("closedAnalyzer", flag.ExitOnError)
	return newAnalyzer(analyzer.Run, flags)
}

// Run implements the main analysis pass
func (a *closedAnalyzer) Run(pass *analysis.Pass) (interface{}, error) {
	pssa, ok := pass.ResultOf[buildssa.Analyzer].(*buildssa.SSA)
	if !ok {
		return nil, nil
	}

	// Build list of types we are looking for
	targetTypes := findTypesNeedingClosed(pssa, sqlPackages)

	// If non of the types are found, skip
	if len(targetTypes) == 0 {
		return nil, nil
	}

	funcs := pssa.SrcFuncs
	for _, f := range funcs {
		for _, b := range f.Blocks {
			for i := range b.Instrs {
				// Check if instruction is call that returns a target pointer type
				targetValues := findResultsNeedingClosed(b, i, targetTypes)
				if len(targetValues) == 0 {
					continue
				}

				// For each found target check if they are closed and deferred
				for _, targetValue := range targetValues {
					refs := targetValue.value.Referrers()
					isClosed := checkClosed(refs, targetTypes)
					if !isClosed {
						pass.Reportf((targetValue.instr).Pos(), "Rows/Stmt/NamedStmt was not closed")
					}
				}
			}
		}
	}

	return nil, nil
}
