# cli-web-docs examples

| Example | Demonstrates |
| ------- | ------------ |
| [`render/`](render/) | The basic case. Tiny CLI, single-page output. Start here. |
| [`multi-page/`](multi-page/) | `Options.PerPage = true`. One HTML file per visible subcommand. |
| [`custom-theme/`](custom-theme/) | Operator-supplied `Options.CSS` to brand the output. |
| [`deep-tree/`](deep-tree/) | Three-level nested command tree with realistic flag surface. |

Each example is a runnable `main.go` and a short README. From the repo root:

```
go run ./examples/render ./site
open site/index.html
```

Pair with [`deploy/Caddyfile.example`](../deploy/Caddyfile.example) when you want a real host.
