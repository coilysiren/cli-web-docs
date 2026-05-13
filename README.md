# cli-web-docs

Static HTML documentation generator for [urfave/cli](https://github.com/urfave/cli) v3 command trees.

Pure Go, no JavaScript on the generated site, minimal CSS. Output is a single self-contained `index.html` (default) or a per-command file tree. Independent of the rest of the cli-* family; works on any `*cli.Command`.

Companion to [urfave/cli-docs](https://github.com/urfave/cli-docs) (Markdown). When you want a browser-renderable artifact you can host on GitHub Pages or serve from any static host, reach here.

Part of the four-package cli-* family intended for the urfave/cli ecosystem:

- [cli-guard](https://github.com/coilysiren/cli-guard) - scope tokens, audit log, argv validation
- [cli-mcp](https://github.com/coilysiren/cli-mcp) - project a command tree as an MCP server
- **cli-web-docs** (this repo) - static HTML documentation
- [cli-web-ops](https://github.com/coilysiren/cli-web-ops) - mobile-first web executor over Tailscale

## Usage

```go
import (
    webdocs "github.com/coilysiren/cli-web-docs"
    "github.com/urfave/cli/v3"
)

func main() {
    app := &cli.Command{Name: "myapp", /* ... */}
    if err := webdocs.Render(app, webdocs.Options{
        OutputDir: "./site",
        Title:     "myapp docs",
    }); err != nil {
        log.Fatal(err)
    }
}
```

See `examples/render/` for a runnable example that emits docs for a tiny CLI.

## Output shape

- `index.html` - landing page with the command tree and root usage
- One section per command, anchored by command path (`#hello`, `#sub-foo-bar`)
- Per-command pages (`Options.PerPage = true`) at `./<path>.html`

## Composition

Independent of cli-guard - operates on any `*cli.Command`. A Guard-wrapped command renders the same as a bare one. Annotations on `cli.Command.Metadata` are surfaced in the generated docs if present (see `Options.MetadataKeys`).

## Status

v0. API will firm up under second-consumer pressure.

## License

MIT.
