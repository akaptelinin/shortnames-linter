package analyzer

import (
	"go/ast"
	"strings"

	"golang.org/x/tools/go/analysis"
)

const minNameLength = 3

var defaultWhitelist = map[string]bool{
	"err": true,
	"ok":  true,
	"ctx": true,
	"i":   true,
	"j":   true,
	"k":   true,
	"n":   true,
	"id":  true,
	"ip":  true,
	"db":  true,
	"tx":  true,
	"mu":  true,
	"wg":  true,
	"rw":  true,
	"fn":  true,
	"cb":  true,
	"ch":  true,
	"v":   true,
	"t":   true,
	"b":   true,
}

var userWhitelist = make(map[string]bool)

var Analyzer = &analysis.Analyzer{
	Name: "shortnames",
	Doc:  "warns about short receiver names, parameter names, and import aliases",
	Run:  run,
}

var (
	flagWhitelist        string
	flagSeverity         string
	flagDisableDefault   bool
)

func init() {
	Analyzer.Flags.StringVar(&flagWhitelist, "whitelist", "",
		"comma-separated list of additional allowed short names")
	Analyzer.Flags.StringVar(&flagSeverity, "severity", "warning",
		"severity level: warning or error")
	Analyzer.Flags.BoolVar(&flagDisableDefault, "disable-default-whitelist", false,
		"disable the default whitelist, only use custom whitelist")
}

func parseUserWhitelist() {
	if flagWhitelist == "" {
		return
	}
	for _, name := range strings.Split(flagWhitelist, ",") {
		name = strings.TrimSpace(name)
		if name != "" {
			userWhitelist[name] = true
		}
	}
}

func isWhitelisted(name string) bool {
	if len(name) >= minNameLength {
		return true
	}
	if !flagDisableDefault && defaultWhitelist[name] {
		return true
	}
	if userWhitelist[name] {
		return true
	}
	return false
}

func run(pass *analysis.Pass) (interface{}, error) {
	parseUserWhitelist()

	for _, file := range pass.Files {
		for _, imp := range file.Imports {
			if imp.Name != nil {
				name := imp.Name.Name
				if name != "_" && name != "." && !isWhitelisted(name) {
					pass.Reportf(imp.Name.Pos(),
						"import alias %q is too short (<%d chars), consider a descriptive name",
						name, minNameLength)
				}
			}
		}

		ast.Inspect(file, func(n ast.Node) bool {
			fn, ok := n.(*ast.FuncDecl)
			if !ok {
				return true
			}

			if fn.Recv != nil {
				for _, field := range fn.Recv.List {
					for _, name := range field.Names {
						if !isWhitelisted(name.Name) {
							pass.Reportf(name.Pos(),
								"receiver name %q is too short (<%d chars), use full type name like %q",
								name.Name, minNameLength, suggestReceiverName(field.Type))
						}
					}
				}
			}

			if fn.Type.Params != nil {
				for _, field := range fn.Type.Params.List {
					for _, name := range field.Names {
						if !isWhitelisted(name.Name) {
							pass.Reportf(name.Pos(),
								"parameter name %q is too short (<%d chars), use descriptive name",
								name.Name, minNameLength)
						}
					}
				}
			}

			if fn.Type.Results != nil {
				for _, field := range fn.Type.Results.List {
					for _, name := range field.Names {
						if !isWhitelisted(name.Name) {
							pass.Reportf(name.Pos(),
								"named return %q is too short (<%d chars), use descriptive name",
								name.Name, minNameLength)
						}
					}
				}
			}

			return true
		})
	}

	return nil, nil
}

func suggestReceiverName(expr ast.Expr) string {
	switch t := expr.(type) {
	case *ast.StarExpr:
		return suggestReceiverName(t.X)
	case *ast.Ident:
		name := t.Name
		if len(name) == 0 {
			return "self"
		}
		return strings.ToLower(name[:1]) + name[1:]
	default:
		return "self"
	}
}
