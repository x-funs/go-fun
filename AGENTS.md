# Repository Guidelines

## Project Structure & Module Organization

This repository is the Go module `github.com/x-funs/go-fun`. Most public utility APIs live in the root package as paired `*.go` and `*_test.go` files, for example `string.go` with `string_test.go`, `http.go` with `http_test.go`, and `aes.go` with `aes_test.go`.

Subpackages are limited and purpose-specific:

- `alias/`: JSON and SQL helper time types.
- `strtotime/`: PHP-style date/time parser used by root datetime helpers.
- `tree/tire/`: trie-like keyword matching implementation.

Documentation and project assets are in `README.md`, `README_zh.md`, `LICENSE`, and `banner.txt`.

## Build, Test, and Development Commands

- `go test ./...`: runs all tests. Note that some existing HTTP tests depend on external sites or `localhost:8080`, so failures may be environment-related.
- `go test ./... -run '^$'`: compile-checks all packages without running tests.
- `go vet ./...`: runs Go's standard static analysis.
- `go test -race ./...`: checks for data races; use it for concurrency-sensitive changes.
- `gofmt -w <files>`: formats edited Go files before commit.

## Coding Style & Naming Conventions

Use standard Go formatting and idioms. Keep exported functions and types documented with concise comments beginning with the exported name. Prefer small, focused helper functions over adding more large variadic APIs unless compatibility requires it.

Tests should follow Go naming conventions: `TestName`, `BenchmarkName`, and file names ending in `_test.go`. Keep package names lowercase. Be aware that existing API names are part of compatibility, even if imperfect.

## Testing Guidelines

The project uses Go's built-in `testing` package plus `github.com/stretchr/testify/assert`. New tests should assert behavior rather than only logging output. Avoid network-dependent tests for new code; prefer `httptest.Server` for HTTP behavior and deterministic fixtures for time, random, and encoding cases.

Cover normal behavior, error returns, nil inputs, empty inputs, and panic-prone boundaries.

## Commit & Pull Request Guidelines

Recent commit messages use short prefixes such as `update:` and `bugfix:` with concise Chinese descriptions. Follow that style unless the maintainers specify otherwise.

Pull requests should include a short problem statement, the behavior changed, test commands run, and any compatibility risks. Link related issues when available.

## Security & Configuration Tips

HTTP and crypto helpers are security-sensitive. Do not add insecure defaults, silent error swallowing, or external network dependencies without documenting the tradeoff clearly.
