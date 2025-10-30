Interpreter (Go)

Scaffold-only repository for a Go-based interpreter. This commit includes only directory structure and placeholder README files. No project code has been added yet.

Planned layout
- `cmd/repl/` — REPL entry point (will contain `main` later).
- `internal/token/` — Token types and keyword handling.
- `internal/lexer/` — Lexer for turning input into tokens.
- `internal/ast/` — AST node definitions.
- `internal/parser/` — Parser that builds the AST from tokens.
- `internal/object/` — Runtime objects/values.
- `internal/evaluator/` — Evaluator to execute AST nodes.

Notes
- Requires Go 1.21+ once implementation starts.
- Replace the module path in `go.mod` as needed for your environment.

Next steps (when ready)
- Add minimal `cmd/repl/main.go` wiring.
- Implement `internal/token` and `internal/lexer` first.
- Add parser, AST, object system, and evaluator incrementally.
