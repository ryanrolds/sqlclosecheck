package analyzer_test

import (
	"testing"

	"github.com/ryanrolds/sqlclosecheck/pkg/analyzer"
	"golang.org/x/tools/go/analysis/analysistest"
)

func XTestClosedAnalyzer(t *testing.T) {
	testdata := analysistest.TestData()

	checker := analyzer.NewClosedAnalyzer()
	analysistest.Run(t, testdata, checker, "closed")
}
