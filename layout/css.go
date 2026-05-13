package layout

// DefaultCSS is the shared stylesheet for the cli-web-* family. Dark-mode
// aware via prefers-color-scheme. Mobile-friendly (max-width plus
// reasonable touch targets). Both cli-web-docs and cli-web-ops use this
// by default; consumers can override via Page.CSS.
const DefaultCSS = `
:root {
  --fg: #1f2328;
  --muted: #57606a;
  --bg: #ffffff;
  --accent: #0969da;
  --rule: #d0d7de;
  --danger: #cf222e;
  --ok: #1a7f37;
  --code-bg: #f6f8fa;
}
@media (prefers-color-scheme: dark) {
  :root {
    --fg: #e6edf3;
    --muted: #8d96a0;
    --bg: #0d1117;
    --accent: #2f81f7;
    --rule: #30363d;
    --danger: #f85149;
    --ok: #3fb950;
    --code-bg: #161b22;
  }
}
* { box-sizing: border-box; }
html, body { margin: 0; padding: 0; }
body {
  background: var(--bg);
  color: var(--fg);
  font: 16px/1.5 -apple-system, BlinkMacSystemFont, "Segoe UI", Helvetica, Arial, sans-serif;
  margin: 0 auto;
  padding: 1.5rem;
  max-width: 56rem;
}
@media (max-width: 30rem) {
  body { padding: 1rem; font-size: 17px; }
}
h1, h2, h3, h4 { line-height: 1.25; }
h1 { font-size: 1.7rem; margin-top: 0; }
h2 { border-bottom: 1px solid var(--rule); padding-bottom: .3rem; margin-top: 2.5rem; }
h3 { margin-top: 2rem; }
.muted { color: var(--muted); }
.subtitle { margin-top: -0.5rem; font-size: 1rem; }
.crumb { font-size: .9rem; margin-bottom: 1rem; }
.path-segment { color: var(--muted); }
a { color: var(--accent); text-decoration: none; }
a:hover { text-decoration: underline; }
code, kbd, pre {
  font: 14px/1.45 ui-monospace, SFMono-Regular, Menlo, monospace;
  background: var(--code-bg);
}
code, kbd { padding: .15em .35em; border-radius: 3px; }
pre { padding: .75rem 1rem; border-radius: 6px; overflow-x: auto; }
pre code { background: none; padding: 0; }
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
nav.sidebar strong { font-size: .85rem; text-transform: uppercase; letter-spacing: .04em; color: var(--muted); display: block; margin-bottom: .5rem; }
main { line-height: 1.6; }
table { border-collapse: collapse; margin: 1rem 0; }
th, td { border: 1px solid var(--rule); padding: .4rem .6rem; text-align: left; }
th { background: var(--code-bg); }

/* Buttons (cli-web-ops). Inline-style so docs pages without buttons
   incur the rules but never see them rendered. */
.btn, button, a.btn {
  display: block; width: 100%;
  background: var(--accent); color: white;
  border: none; border-radius: 12px;
  padding: 1rem 1.25rem;
  font-size: 1.1rem; font-weight: 600;
  text-align: left; text-decoration: none;
  margin-bottom: .75rem;
  cursor: pointer;
  -webkit-tap-highlight-color: transparent;
}
.btn.confirm { background: var(--danger); }
.btn .sub { display: block; font-size: .85rem; font-weight: 400; opacity: .9; margin-top: .15rem; }
.btn:hover { text-decoration: none; opacity: .92; }

/* Form (cli-web-ops). */
form .field { margin-bottom: 1rem; }
form label { display: block; font-weight: 600; margin-bottom: .25rem; }
form .field .desc { color: var(--muted); font-size: .85rem; font-weight: 400; margin-bottom: .35rem; }
form input[type=text], form input[type=number], form select {
  width: 100%; padding: .65rem;
  font-size: 1rem;
  background: var(--bg); color: var(--fg);
  border: 1px solid var(--rule); border-radius: 8px;
}
form input[type=checkbox] { width: 1.25rem; height: 1.25rem; }

/* Output stream (cli-web-ops). */
pre.log {
  background: #000; color: #e6edf3;
  border-radius: 8px; padding: 1rem;
  font: 13px/1.5 ui-monospace, SFMono-Regular, Menlo, monospace;
  white-space: pre-wrap; word-wrap: break-word;
  min-height: 8rem; max-height: 60vh; overflow-y: auto;
  margin-top: 1rem;
}
.log .err  { color: #ff7b72; }
.log .ok   { color: #3fb950; }
.log .meta { color: #8d96a0; }

/* Markdown rendered content (descriptions, etc). */
.desc-md p { margin: .25rem 0; }
.desc-md code { background: var(--code-bg); padding: .1em .3em; border-radius: 3px; font-size: .9em; }

/* Tabs - used on tool pages that combine docs + run. */
.tabs { display: flex; gap: .5rem; margin: 1rem 0 .5rem; border-bottom: 1px solid var(--rule); }
.tabs a { padding: .5rem .9rem; border-radius: 6px 6px 0 0; color: var(--muted); }
.tabs a.active { color: var(--fg); background: var(--code-bg); border: 1px solid var(--rule); border-bottom-color: var(--bg); margin-bottom: -1px; }
`
