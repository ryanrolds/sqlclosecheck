package analyzer_test

import (
	"testing"

	"github.com/ryanrolds/sqlclosecheck/pkg/analyzer"
	"golang.org/x/tools/go/analysis/analysistest"
)

var closedTestPackages = []string{
	"github.com/ryanrolds/sqlclosecheck/pkg/analyzer/testdata/closed",
}

func XTestClosedAnalyzer(t *testing.T) {
	t.Parallel()

	testdata := analysistest.TestData()
	checker := analyzer.NewClosedAnalyzer()

	for _, pkg := range closedTestPackages {
		pkg := pkg

		t.Run(pkg, func(t *testing.T) {
			t.Parallel()

			analysistest.Run(t, testdata, checker, pkg)
		})
	}
}
