# cli-web-docs

cli-web-docs is a static HTML documentation generator for [urfave/cli](https://github.com/urfave/cli) v3 command trees.

Pipeline: `*cli.Command` → canonical Markdown from [urfave/cli-docs](https://github.com/urfave/cli-docs) → HTML via [goldmark](https://github.com/yuin/goldmark) → shared layout shell. Pure Go, no JavaScript in the generated output, dark-mode-aware default stylesheet.

The layout subpackage is also consumed by [cli-web-ops](https://github.com/coilysiren/cli-web-ops), so the documentation site and the live operator console share visual identity.

## Where to go next

- **[Features](FEATURES.md)** - feature inventory.
- **[Examples](examples.md)** - render, multi-page, custom-theme, deep-tree.
- **[Source on GitHub](https://github.com/coilysiren/cli-web-docs)** - issues, releases, code.

cli-web-docs is part of the cli-* family: [cli-guard](https://github.com/coilysiren/cli-guard), [cli-mcp](https://github.com/coilysiren/cli-mcp), [cli-web-ops](https://github.com/coilysiren/cli-web-ops).
