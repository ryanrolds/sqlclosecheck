package analyzer

import (
	"flag"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/buildssa"
)

// NewAnalyzer returns a non-configurable analyzer that defaults to the defer-only mode.
// Deprecated, this will be removed in v1.0.0.
func NewAnalyzer(extraPackages []string) *analysis.Analyzer {
	flags := flag.NewFlagSet("analyzer", flag.ExitOnError)
	return newAnalyzer(func(pass *analysis.Pass) (interface{}, error) {
		return run(pass, extraPackages)
	}, flags)
}

func run(pass *analysis.Pass, extraPackages []string) (interface{}, error) {
	opinionatedAnalyzer := &deferOnlyAnalyzer{
		extraPackages: extraPackages,
	}
	return opinionatedAnalyzer.Run(pass)
}

// newAnalyzer returns a new analyzer with the given run function, should be used by all analyzers.
func newAnalyzer(
	r func(pass *analysis.Pass) (interface{}, error),
	flags *flag.FlagSet,
) *analysis.Analyzer {
	return &analysis.Analyzer{
		Name: "sqlclosecheck",
		Doc:  "Checks that sql.Rows, sql.Stmt, sqlx.NamedStmt, pgx.Query are closed.",
		Run:  r,
		Requires: []*analysis.Analyzer{
			buildssa.Analyzer,
		},
	}
}
