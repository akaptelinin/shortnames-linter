# shortnames-linter

[![Tests](https://github.com/akaptelinin/shortnames-linter/actions/workflows/test.yml/badge.svg)](https://github.com/akaptelinin/shortnames-linter/actions/workflows/test.yml)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

Go linter that warns about short identifier names (<3 chars) to improve code readability.

## What it checks

| Check | Example bad | Example good |
|-------|-------------|--------------|
| Receiver names | `func (h *Handler)` | `func (handler *Handler)` |
| Parameter names | `func Process(c ctx, r req)` | `func Process(ctx context.Context, req *Request)` |
| Import aliases | `import c "context"` | `import ctx "context"` |
| Named returns | `func f() (s string)` | `func f() (result string)` |

## Whitelist

These short names are allowed by default:

`err`, `ok`, `ctx`, `i`, `j`, `k`, `n`, `id`, `ip`, `db`, `tx`, `mu`, `wg`, `rw`, `fn`, `cb`, `ch`, `v`, `t`, `b`

## Installation

### Standalone

```bash
go install github.com/akaptelinin/shortnames-linter@latest
```

Run:
```bash
shortnames ./...
```

### With golangci-lint

Build plugin:
```bash
go build -buildmode=plugin -o shortnames.so ./plugin
```

Add to `.golangci.yml`:
```yaml
linters-settings:
  custom:
    shortnames:
      path: /path/to/shortnames.so

linters:
  enable:
    - shortnames
```

## Flags

| Flag | Default | Description |
|------|---------|-------------|
| `-whitelist` | (empty) | Comma-separated additional allowed names |
| `-severity` | warning | `warning` or `error` |

Example:
```bash
shortnames -whitelist="io,fs,tt" ./...
```

## Example output

```
main.go:4:2:  import alias "x" is too short (<3 chars), consider a descriptive name
main.go:13:7: receiver name "h" is too short (<3 chars), use full type name like "handler"
main.go:13:27: parameter name "a" is too short (<3 chars), use descriptive name
```
