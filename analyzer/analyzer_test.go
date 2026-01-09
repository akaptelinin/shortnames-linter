package analyzer_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/akaptelinin/shortnames-linter/analyzer"
	"golang.org/x/tools/go/analysis/analysistest"
)

func TestAnalyzer(t *testing.T) {
	testdata := filepath.Join(os.Getenv("PWD"), "testdata")
	analysistest.Run(t, testdata, analyzer.Analyzer, "a")
}
