package webdocs

const defaultCSS = `
:root {
  --fg: #1f2328;
  --muted: #57606a;
  --bg: #ffffff;
  --accent: #0969da;
  --rule: #d0d7de;
  --code-bg: #f6f8fa;
}
@media (prefers-color-scheme: dark) {
  :root {
    --fg: #e6edf3;
    --muted: #8d96a0;
    --bg: #0d1117;
    --accent: #2f81f7;
    --rule: #30363d;
    --code-bg: #161b22;
  }
}
body {
  background: var(--bg);
  color: var(--fg);
  font: 16px/1.5 -apple-system, BlinkMacSystemFont, "Segoe UI", Helvetica, Arial, sans-serif;
  margin: 0;
  padding: 1.5rem;
  max-width: 56rem;
  margin-left: auto;
  margin-right: auto;
}
h1, h2, h3, h4 { line-height: 1.25; }
h1 { font-size: 1.8rem; margin-top: 0; }
h2 { border-bottom: 1px solid var(--rule); padding-bottom: .3rem; margin-top: 2.5rem; }
h3 { margin-top: 2rem; }
code, kbd, pre {
  font: 14px/1.45 ui-monospace, SFMono-Regular, Menlo, monospace;
  background: var(--code-bg);
}
code, kbd { padding: .15em .35em; border-radius: 3px; }
pre { padding: .75rem 1rem; border-radius: 6px; overflow-x: auto; }
pre code { background: none; padding: 0; }
a { color: var(--accent); text-decoration: none; }
a:hover { text-decoration: underline; }
.muted { color: var(--muted); }
nav.sidebar {
  border: 1px solid var(--rule);
  border-radius: 6px;
  padding: 1rem;
  margin: 1rem 0 2rem;
  background: var(--code-bg);
}
nav.sidebar ul { list-style: none; padding-left: 1rem; margin: 0; }
nav.sidebar > ul { padding-left: 0; }
nav.sidebar li { margin: .25rem 0; }
.path-segment { color: var(--muted); }
.crumb { font-size: .9rem; margin-bottom: 1rem; }
table { border-collapse: collapse; margin: 1rem 0; }
th, td { border: 1px solid var(--rule); padding: .4rem .6rem; text-align: left; }
th { background: var(--code-bg); }
`

const layoutTpl = `<!doctype html>
<html lang="en">
<head>
  <meta charset="utf-8">
  <meta name="viewport" content="width=device-width, initial-scale=1">
  <title>{{.Title}}</title>
  <style>{{.CSS}}</style>
</head>
<body>
  {{if .Crumb}}<p class="muted crumb">{{.Crumb}}</p>{{end}}
  <h1>{{.Heading}}</h1>

  <nav class="sidebar">
    <strong>Commands</strong>
    {{.Nav}}
  </nav>

  {{.Body}}
</body>
</html>
`
