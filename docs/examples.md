# Examples

Four runnable examples under [`examples/`](https://github.com/coilysiren/cli-web-docs/tree/main/examples), each demonstrating a different dimension of cli-web-docs.

| Example | Demonstrates |
| ------- | ------------ |
| [`render/`](https://github.com/coilysiren/cli-web-docs/tree/main/examples/render) | The basic case. Tiny CLI, single-page output. Start here. |
| [`multi-page/`](https://github.com/coilysiren/cli-web-docs/tree/main/examples/multi-page) | `Options.PerPage = true`. One HTML file per visible subcommand. |
| [`custom-theme/`](https://github.com/coilysiren/cli-web-docs/tree/main/examples/custom-theme) | Operator-supplied `Options.CSS` to brand the output. |
| [`deep-tree/`](https://github.com/coilysiren/cli-web-docs/tree/main/examples/deep-tree) | Three-level nested command tree with realistic flag surface. |

## Running

From the cli-web-docs repo root:

```bash
go run ./examples/render ./site
open site/index.html
```

Pair with [`deploy/Caddyfile.example`](https://github.com/coilysiren/cli-web-docs/blob/main/deploy/Caddyfile.example) when you want a real host.
