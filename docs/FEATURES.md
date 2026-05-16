# cli-web-docs features

Inventory of what cli-web-docs does today. Scope changes should land in the same commit that touches code, so this file stays a faithful mirror of the public API.

## Rendering

- **`webdocs.Render(*cli.Command, Options) error`** - Walk a urfave/cli v3 command tree, generate canonical Markdown via [urfave/cli-docs](https://github.com/urfave/cli-docs), render to HTML via [goldmark](https://github.com/yuin/goldmark), wrap in the shared layout shell. Writes HTML files to `Options.OutputDir`.
- **Single-page mode** (`Options.PerPage = false`) - One `index.html` with every command inlined under section anchors.
- **Multi-page mode** (`Options.PerPage = true`) - One HTML file per visible subcommand at `<path>.html`, plus an index that lists the tree.
- **Custom themes** - `Options.CSS` overrides the default stylesheet wholesale. Layout shell targets stable class names (`nav.sidebar`, `.muted`, `.crumb`, etc).
- **Hidden commands skipped** - `cli.Command.Hidden = true` excludes a command from both nav and rendered pages.

## Layout subpackage

- **`github.com/coilysiren/cli-web-docs/layout`** - Dep-free HTML shell consumed by both this repo and [cli-web-ops](https://github.com/coilysiren/cli-web-ops). Exports `Page`, `Render`, `NavItem`, `BuildNavList`, and `DefaultCSS`. Operators get visually consistent docs + ops pages by importing this subpackage.

## Examples (`examples/`)

- **render/** - Basic case: tiny CLI, single-page.
- **multi-page/** - Same shape with `Options.PerPage = true`.
- **custom-theme/** - Operator-supplied CSS demonstrating wholesale rebrand.
- **deep-tree/** - Three-level nested command tree with realistic flag surface.

## Hosting

- **Static output** - Pure HTML, no JavaScript. Anything that serves files works.
- **`deploy/Caddyfile.example`** - Recommended posture: Tailscale-bound private docs site, with a public-docs alternative commented out.

## Repo development

- `.agent-guard/agent-guard.yaml` declares local dev verbs.
- `Makefile` is the source of truth.
- `coily lint` validates the yaml/Makefile contract on every CI run.
- `.golangci.yaml`, `staticcheck.conf` mirror urfave/cli.
- GitHub Actions CI runs vet/build/test/lint.

## See also

- [README.md](../README.md) - human-facing intro.
- [AGENTS.md](../AGENTS.md) - agent-facing operating rules.
- [.agent-guard/agent-guard.yaml](../.agent-guard/agent-guard.yaml) - allowlisted commands.

Cross-reference convention from [coilysiren/agentic-os#59](https://github.com/coilysiren/agentic-os/issues/59).
