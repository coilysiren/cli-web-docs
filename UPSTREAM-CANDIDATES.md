# Upstream candidates for urfave/cli

After the refactor onto [urfave/cli-docs/v3](https://github.com/urfave/cli-docs) and [yuin/goldmark](https://github.com/yuin/goldmark), this repo is a thin layout layer. Most of what would have been an upstream candidate now lives in cli-docs already.

## Candidates this repo would still benefit from upstreaming

- **`cli.Command.Walk(func(*cli.Command))`** - the recursive tree walker. Used by `buildNavHTML` and `renderPerPage`. Same shape in cli-mcp and cli-web-ops.
- **`cli.Command.Path() []string`** - path from root to a command. Used here for nav slugs and crumbs. cli-docs likely has the same need internally; cross-package consolidation would help.

## Candidates this repo originates

- **HTML layout convention for cli-docs Markdown output.** The layout shell here (dark-mode CSS, sidebar nav, per-page slug scheme) could become a sister artifact to cli-docs, possibly as `cli-docs/v3/html` once the contract is stable.
- **Goldmark extension defaults for CLI documentation.** GFM + autolinks + fenced code is the right baseline; if cli-docs upstream adopts HTML rendering, it should reach for the same defaults.

## Closed by the refactor

The previous version of this file listed flag-to-HTML mapping, command introspection, and most rendering primitives as upstream candidates. Those are all subsumed by cli-docs now. This is what the four-package experiment is supposed to surface: copy-pasted helpers across consumers, decide what's stable, promote then.

The shape: pin patterns here, copy into each consumer, watch what's stable across consumers, promote then.
