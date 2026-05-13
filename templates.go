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
a { color: var(--accent); text-decoration: none; }
a:hover { text-decoration: underline; }
.muted { color: var(--muted); }
.flag-name { font-weight: 600; }
ul.flags { list-style: none; padding-left: 0; }
ul.flags li { margin-bottom: .5rem; }
nav.tree ul { list-style: none; padding-left: 1rem; }
nav.tree { border: 1px solid var(--rule); border-radius: 6px; padding: 1rem; margin: 1rem 0; }
.path-segment { color: var(--muted); }
.metadata { font-size: .9rem; color: var(--muted); margin: .5rem 0; }
.metadata dt { font-weight: 600; display: inline; }
.metadata dd { display: inline; margin: 0 1rem 0 .25rem; }
`

const singlePageTpl = `<!doctype html>
<html lang="en">
<head>
  <meta charset="utf-8">
  <meta name="viewport" content="width=device-width, initial-scale=1">
  <title>{{.Title}}</title>
  <style>{{.CSS}}</style>
</head>
<body>
  <h1>{{.Title}}</h1>
  {{if .Root.Usage}}<p class="muted">{{.Root.Usage}}</p>{{end}}
  {{if .Root.Description}}<p>{{.Root.Description}}</p>{{end}}

  <nav class="tree">
    <strong>Commands</strong>
    {{template "subtree" .Root}}
  </nav>

  {{template "node" dict "Node" .Root "MetadataKeys" .MetadataKeys}}
</body>
</html>
{{define "subtree"}}
  {{if .Subcommands}}
  <ul>
    {{range .Subcommands}}
      <li><a href="#{{.Anchor}}"><code>{{join .Path " "}}</code></a> {{if .Usage}}<span class="muted">{{.Usage}}</span>{{end}}
        {{template "subtree" .}}
      </li>
    {{end}}
  </ul>
  {{end}}
{{end}}
{{define "node"}}
  {{$n := .Node}}{{$keys := .MetadataKeys}}
  <section id="{{$n.Anchor}}">
    <h2><code>{{join $n.Path " "}}</code></h2>
    {{if $n.Usage}}<p>{{$n.Usage}}</p>{{end}}
    {{if $n.Description}}<p>{{$n.Description}}</p>{{end}}
    {{if $n.ArgsUsage}}<p><strong>Usage:</strong> <code>{{join $n.Path " "}} {{$n.ArgsUsage}}</code></p>{{end}}

    {{if hasMetadata $n.Metadata $keys}}
      <dl class="metadata">
        {{range $keys}}{{if metaValue $n.Metadata .}}<dt>{{.}}:</dt><dd>{{metaValue $n.Metadata .}}</dd>{{end}}{{end}}
      </dl>
    {{end}}

    {{if $n.Flags}}
      <h3>Flags</h3>
      <ul class="flags">
        {{range $n.Flags}}
          <li>
            <span class="flag-name">--{{.Name}}</span>{{range .Aliases}} / <span class="flag-name">--{{.}}</span>{{end}}
            {{if .Default}}<span class="muted"> (default: <code>{{.Default}}</code>)</span>{{end}}
            {{if .Usage}}<br><span class="muted">{{.Usage}}</span>{{end}}
          </li>
        {{end}}
      </ul>
    {{end}}
  </section>
  {{range $n.Subcommands}}{{template "node" dict "Node" . "MetadataKeys" $keys}}{{end}}
{{end}}
`

const indexTpl = `<!doctype html>
<html lang="en">
<head>
  <meta charset="utf-8">
  <meta name="viewport" content="width=device-width, initial-scale=1">
  <title>{{.Title}}</title>
  <style>{{.CSS}}</style>
</head>
<body>
  <h1>{{.Title}}</h1>
  {{if .Root.Usage}}<p class="muted">{{.Root.Usage}}</p>{{end}}
  {{if .Root.Description}}<p>{{.Root.Description}}</p>{{end}}

  <nav class="tree">
    <strong>Commands</strong>
    {{template "subtree" .Root}}
  </nav>
</body>
</html>
{{define "subtree"}}
  {{if .Subcommands}}
  <ul>
    {{range .Subcommands}}
      <li><a href="{{.Slug}}"><code>{{join .Path " "}}</code></a> {{if .Usage}}<span class="muted">{{.Usage}}</span>{{end}}
        {{template "subtree" .}}
      </li>
    {{end}}
  </ul>
  {{end}}
{{end}}
`

const perPageTpl = `<!doctype html>
<html lang="en">
<head>
  <meta charset="utf-8">
  <meta name="viewport" content="width=device-width, initial-scale=1">
  <title>{{join .Current.Path " "}} - {{.Title}}</title>
  <style>{{.CSS}}</style>
</head>
<body>
  <p class="muted"><a href="index.html">{{.Root.Name}}</a> / <span class="path-segment">{{join .Current.Path " "}}</span></p>
  <h1><code>{{join .Current.Path " "}}</code></h1>
  {{$n := .Current}}{{$keys := .MetadataKeys}}
  {{if $n.Usage}}<p>{{$n.Usage}}</p>{{end}}
  {{if $n.Description}}<p>{{$n.Description}}</p>{{end}}
  {{if $n.ArgsUsage}}<p><strong>Usage:</strong> <code>{{join $n.Path " "}} {{$n.ArgsUsage}}</code></p>{{end}}

  {{if hasMetadata $n.Metadata $keys}}
    <dl class="metadata">
      {{range $keys}}{{if metaValue $n.Metadata .}}<dt>{{.}}:</dt><dd>{{metaValue $n.Metadata .}}</dd>{{end}}{{end}}
    </dl>
  {{end}}

  {{if $n.Flags}}
    <h2>Flags</h2>
    <ul class="flags">
      {{range $n.Flags}}
        <li>
          <span class="flag-name">--{{.Name}}</span>{{range .Aliases}} / <span class="flag-name">--{{.}}</span>{{end}}
          {{if .Default}}<span class="muted"> (default: <code>{{.Default}}</code>)</span>{{end}}
          {{if .Usage}}<br><span class="muted">{{.Usage}}</span>{{end}}
        </li>
      {{end}}
    </ul>
  {{end}}

  {{if $n.Subcommands}}
    <h2>Subcommands</h2>
    <ul>
      {{range $n.Subcommands}}
        <li><a href="{{.Slug}}"><code>{{join .Path " "}}</code></a> {{if .Usage}}<span class="muted">- {{.Usage}}</span>{{end}}</li>
      {{end}}
    </ul>
  {{end}}
</body>
</html>
`
