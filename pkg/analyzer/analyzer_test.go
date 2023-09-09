package analyzer_test

import (
	"testing"

	"github.com/ryanrolds/sqlclosecheck/pkg/analyzer"
	"golang.org/x/tools/go/analysis/analysistest"
)

func TestAnalyzer(t *testing.T) {
	t.Parallel()

	testdata := analysistest.TestData()
	checker := analyzer.NewAnalyzer()

	packages := []string{
		"github.com/ryanrolds/sqlclosecheck/pkg/analyzer/testdata/rows",
		"github.com/ryanrolds/sqlclosecheck/pkg/analyzer/testdata/stmt",
		"github.com/ryanrolds/sqlclosecheck/pkg/analyzer/testdata/pgx",
	}

	for _, pkg := range packages {
		pkg := pkg

		t.Run(pkg, func(t *testing.T) {
			t.Parallel()

			analysistest.Run(t, testdata, checker, pkg)
		})
	}
}
