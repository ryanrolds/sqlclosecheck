package analyzer_test

import (
	"testing"

	"github.com/ryanrolds/sqlclosecheck/pkg/analyzer"
	"golang.org/x/tools/go/analysis/analysistest"
)

func TestConfigurableAnalyzer(t *testing.T) {
	testdata := analysistest.TestData()

	checker := analyzer.NewConfigurableAnalyzer(analyzer.ConfigurableAnalyzerDeferOnly)
	analysistest.Run(t, testdata, checker, "rows", "stmt")
}
