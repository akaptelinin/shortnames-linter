package main

import (
	"github.com/akaptelinin/shortnames-linter/analyzer"

	"golang.org/x/tools/go/analysis"
)

type AnalyzerPlugin struct{}

func (AnalyzerPlugin) GetAnalyzers() []*analysis.Analyzer {
	return []*analysis.Analyzer{analyzer.Analyzer}
}

var AnalyzerPluginVar AnalyzerPlugin
