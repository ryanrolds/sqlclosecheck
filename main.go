package main

import (
	"github.com/ryanrolds/sqlclosecheck/pkg/analyzer"
	"golang.org/x/tools/go/analysis/singlechecker"
)

const (
	defaultMode = analyzer.ConfigurableAnalyzerDeferOnly
)

func main() {
	checker := analyzer.NewConfigurableAnalyzer(defaultMode)
	singlechecker.Main(checker)
}
