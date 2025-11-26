Interpreter (Go)

Planned layout
- `cmd/repl/` — REPL entry point (will contain `main` later).
- `internal/token/` — Token types and keyword handling.
- `internal/lexer/` — Lexer for turning input into tokens.
- `internal/ast/` — AST node definitions.
- `internal/parser/` — Parser that builds the AST from tokens.
- `internal/object/` — Runtime objects/values.
- `internal/evaluator/` — Evaluator to execute AST nodes.

Next steps
- Add parser, AST, object system, and evaluator incrementally.
