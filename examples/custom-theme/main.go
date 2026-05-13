// Command custom-theme demonstrates overriding the default stylesheet
// with operator-supplied CSS. The page structure stays the same; only
// colors, typography, and spacing differ. Pair with custom Options.Title
// to produce branded documentation.
package main

import (
	"fmt"
	"os"

	webdocs "github.com/coilysiren/cli-web-docs"
	"github.com/urfave/cli/v3"
)

// retroCSS is a deliberately-not-the-default look. Replace with your
// company's design tokens to brand a docs site.
const retroCSS = `
:root {
  --fg: #00ff66;
  --muted: #4d7d52;
  --bg: #0a0e0a;
  --accent: #ff00aa;
  --rule: #1a3a1a;
  --code-bg: #0e1610;
}
body {
  background: var(--bg);
  color: var(--fg);
  font: 16px/1.6 ui-monospace, "Cascadia Code", Menlo, monospace;
  margin: 0 auto;
  padding: 2rem;
  max-width: 52rem;
}
h1, h2, h3 { text-transform: uppercase; letter-spacing: .08em; }
h1 { color: var(--accent); border-bottom: 2px solid var(--accent); padding-bottom: .5rem; }
h2 { color: var(--accent); border-bottom: 1px dashed var(--rule); padding-bottom: .3rem; margin-top: 2.5rem; }
a { color: var(--accent); text-decoration: underline; }
a:hover { background: var(--accent); color: var(--bg); }
code, kbd, pre {
  font-family: ui-monospace, "Cascadia Code", Menlo, monospace;
  background: var(--code-bg);
  border: 1px solid var(--rule);
}
code, kbd { padding: .15em .4em; }
pre { padding: 1rem; overflow-x: auto; }
.muted { color: var(--muted); }
nav.sidebar {
  border: 1px solid var(--accent);
  padding: 1rem;
  margin: 1rem 0 2rem;
  background: var(--code-bg);
}
nav.sidebar ul { list-style: ">> " inside; padding-left: 0; margin: 0; }
nav.sidebar li { margin: .25rem 0; }
nav.sidebar strong { color: var(--accent); font-size: .85rem; }
table { border-collapse: collapse; }
th, td { border: 1px solid var(--rule); padding: .4rem .8rem; }
th { background: var(--accent); color: var(--bg); }
`

func main() {
	app := &cli.Command{
		Name:  "neon",
		Usage: "branded CLI - custom theme docs demo",
		Commands: []*cli.Command{
			{Name: "init", Usage: "scaffold a new project"},
			{Name: "build", Usage: "build the project"},
			{Name: "ship", Usage: "deploy the build"},
		},
	}

	out := "site"
	if len(os.Args) > 1 {
		out = os.Args[1]
	}

	if err := webdocs.Render(app, webdocs.Options{
		OutputDir: out,
		Title:     "NEON CLI",
		CSS:       retroCSS,
	}); err != nil {
		fmt.Fprintln(os.Stderr, "render:", err)
		os.Exit(1)
	}
	fmt.Println("wrote", out)
}
