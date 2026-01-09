# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [0.1.0] - 2025-01-09

### Added
- Initial release
- Lint checks for:
  - Short receiver names (<3 chars)
  - Short parameter names (<3 chars)
  - Short import aliases (<3 chars)
  - Short named return values (<3 chars)
- Default whitelist: `err`, `ok`, `ctx`, `i`, `j`, `k`, `n`, `id`, `ip`, `db`, `tx`, `mu`, `wg`, `rw`, `fn`, `cb`, `ch`, `v`, `t`, `b`
- Custom whitelist via `--whitelist` flag
- Standalone binary and golangci-lint plugin support
- Comprehensive test coverage (~25 cases)
