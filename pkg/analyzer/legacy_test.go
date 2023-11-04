package analyzer_test

import (
	"testing"

	"github.com/ryanrolds/sqlclosecheck/pkg/analyzer"
	"golang.org/x/tools/go/analysis/analysistest"
)

var legacyTestPackages = []string{
	//"github.com/ryanrolds/sqlclosecheck/pkg/analyzer/testdata/defer_close",
	//"github.com/ryanrolds/sqlclosecheck/pkg/analyzer/testdata/defer_only",
	//"github.com/ryanrolds/sqlclosecheck/pkg/analyzer/testdata/negative",
	"github.com/ryanrolds/sqlclosecheck/pkg/analyzer/testdata/passed",
	//"github.com/ryanrolds/sqlclosecheck/pkg/analyzer/testdata/returned",
}

func TestLegacyAnalyzer(t *testing.T) {
	t.Parallel()

	testdata := analysistest.TestData()
	checker := analyzer.NewLegacyAnalyzer()

	for _, pkg := range legacyTestPackages {
		pkg := pkg

		t.Run(pkg, func(t *testing.T) {
			t.Parallel()

			analysistest.Run(t, testdata, checker, pkg)
		})
	}
}
