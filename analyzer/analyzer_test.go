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

func TestWhitelist(t *testing.T) {
	testdata := filepath.Join(os.Getenv("PWD"), "testdata")
	a := analyzer.Analyzer
	if err := a.Flags.Set("whitelist", "ab,xy"); err != nil {
		t.Fatal(err)
	}
	analysistest.Run(t, testdata, a, "whitelist")
	_ = a.Flags.Set("whitelist", "")
}

func TestDisableDefaultWhitelist(t *testing.T) {
	testdata := filepath.Join(os.Getenv("PWD"), "testdata")
	a := analyzer.Analyzer
	if err := a.Flags.Set("disable-default-whitelist", "true"); err != nil {
		t.Fatal(err)
	}
	analysistest.Run(t, testdata, a, "nodefault")
	_ = a.Flags.Set("disable-default-whitelist", "false")
}
