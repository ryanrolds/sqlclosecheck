package analyzer_test

import (
	"testing"

	"github.com/ryanrolds/sqlclosecheck/pkg/analyzer"
	"golang.org/x/tools/go/analysis/analysistest"
)

func XTestAnalyzer(t *testing.T) {
	t.Parallel()

	testdata := analysistest.TestData()
	checker := analyzer.NewAnalyzer()

	for _, pkg := range legacyTestPackages {
		pkg := pkg

		t.Run(pkg, func(t *testing.T) {
			t.Parallel()

			analysistest.Run(t, testdata, checker, pkg)
		})
	}
}
