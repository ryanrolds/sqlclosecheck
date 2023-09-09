package analyzer_test

import (
	"testing"

	"github.com/ryanrolds/sqlclosecheck/pkg/analyzer"
	"golang.org/x/tools/go/analysis/analysistest"
)

func TestDeferOnlyAnalyzer(t *testing.T) {
	testdata := analysistest.TestData()

	analyzer := analyzer.NewDeferOnlyAnalyzer()
	analysistest.Run(t, testdata, analyzer, "rows", "stmt")
}
