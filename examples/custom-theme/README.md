# custom-theme example

Pass `Options.CSS` to override the default stylesheet wholesale. The page structure (nav, headings, code blocks, tables) stays the same so any CSS that targets `body`, `h1`, `h2`, `nav.sidebar`, `code`, `pre`, `table`, `.muted` will land cleanly.

```
$ go run ./examples/custom-theme ./site
$ open ./site/index.html
```

This example uses a deliberately-loud retro green-and-magenta theme. Replace `retroCSS` with your company's design tokens (or import a stylesheet via `Options.CSS` containing `@import url(...)`).
