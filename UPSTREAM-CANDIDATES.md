# Upstream candidates for urfave/cli

## Candidates this repo would benefit from upstreaming

- **`cli.Command.Walk(func(*cli.Command))`** - recursive tree walker. Currently reimplemented in `webdocs.buildNode`. Same shape in cli-mcp's `registerTree` and cli-web-ops's `collectFavorites`.
- **`cli.Command.Path() []string`** - path from root to this command. Used here for anchor and slug generation.
- **Flag → DocGenerationFlag interface coverage**. urfave's `DocGenerationFlag` exposes `GetUsage()` and `GetDefaultText()` for most built-in flag types, but coverage across third-party flag types is uneven.

## Candidates this repo originates

- **HTML escape rules for help text**. The default `html/template` autoescape is the right call; documenting the convention upstream would help other generators avoid raw-HTML embedding pitfalls.

The shape: pin patterns here, copy into each consumer, watch what's stable across consumers, promote then.
