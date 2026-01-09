# shortnames-linter

[![Tests](https://github.com/akaptelinin/shortnames-linter/actions/workflows/test.yml/badge.svg)](https://github.com/akaptelinin/shortnames-linter/actions/workflows/test.yml)
[![codecov](https://codecov.io/gh/akaptelinin/shortnames-linter/graph/badge.svg)](https://codecov.io/gh/akaptelinin/shortnames-linter)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

Go linter that warns about short identifier names (<3 chars) to improve code readability.

## Why?

Why did you decide that single-letter names are correct approach? Is typing a paid feature on your keyboard? You think >2 letter variables are too long for Go, which is famously known for being as compact as possible?

Anyway, this plugin fixes goslop for you.

## Why no "scope distance"?

Some linters allow short names if a variable is used within N lines ("scope distance"). We don't.

```go
// Other linters: "scope distance = 5, looks good!"
func handler(c *gin.Context) {
    u, err := getUser(c)        // line 1: ok, remembering u=user
    if err != nil {             // line 2: thinking about err now
        c.JSON(500, err)        // line 3: wait, what's c again?
        return                  // line 4
    }
    c.JSON(200, u)              // line 5: what's u? ah, user...
}
```

**5 lines is a lot.** It's a DB query, error check, response, logging. Your brain doesn't cache single-letter variables across `if` statements and function calls.

```go
// shortnames-linter: "no. just no."
func handler(ctx *gin.Context) {
    user, err := getUser(ctx)
    if err != nil {
        ctx.JSON(500, err)
        return
    }
    ctx.JSON(200, user)         // instantly readable
}
```

**Every short name = mental lookup.** We optimize for readers, not writers.

Also: your monitor might scroll the function header off-screen while you might be eating a burger 🍔. You might be reviewing a PR at 11pm, drunk. In all these cases, `user.Process(ctx)` saves you,  `u.P(c)` — does not.

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
| `-disable-default-whitelist` | false | Disable default whitelist, only use custom |

Example:
```bash
shortnames -whitelist="io,fs,tt" ./...
shortnames -disable-default-whitelist ./...
```

## Example output

```
main.go:4:2:  import alias "x" is too short (<3 chars), consider a descriptive name
main.go:13:7: receiver name "h" is too short (<3 chars), use full type name like "handler"
main.go:13:27: parameter name "a" is too short (<3 chars), use descriptive name
```
