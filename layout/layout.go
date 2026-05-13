// Package layout is the shared HTML layout shell for the cli-web-* family
// (cli-web-docs and cli-web-ops). It exposes a small surface so both
// packages produce visually consistent output: same CSS, same nav,
// same dark-mode handling, same crumb shape.
//
// Why a subpackage: cli-web-ops needs the layout without dragging in
// urfave/cli-docs and the goldmark integration cli-web-docs uses. This
// package is dep-free beyond the standard library.
package layout

import (
	"bytes"
	"fmt"
	"html/template"
	"io"
)

// Page is the input to Render. Zero value is usable for a minimal page;
// every field is optional except Body.
type Page struct {
	// Title is the <title> and the visible <h1>.
	Title string

	// Subtitle, if set, is rendered under the title as a muted line.
	Subtitle string

	// Crumb, if set, is rendered above the title (e.g. a "back" link
	// or breadcrumb trail). Caller is responsible for escaping.
	Crumb template.HTML

	// Nav, if set, is rendered in a sidebar block under the title.
	// Caller is responsible for escaping. Use the helpers in this
	// package (e.g. BuildNavList) to produce well-formed nav HTML.
	Nav template.HTML

	// Body is the page content. Caller is responsible for escaping.
	Body template.HTML

	// ExtraHead is appended verbatim inside <head>. Optional script
	// or stylesheet additions per consumer.
	ExtraHead template.HTML

	// CSS overrides the default stylesheet. Empty means use the
	// shared DefaultCSS.
	CSS string
}

// Render writes a complete HTML page to w.
func Render(w io.Writer, p Page) error {
	t, err := template.New("layout").Parse(layoutTpl)
	if err != nil {
		return fmt.Errorf("layout: parse: %w", err)
	}
	css := p.CSS
	if css == "" {
		css = DefaultCSS
	}
	return t.Execute(w, struct {
		Title     string
		Subtitle  string
		Crumb     template.HTML
		Nav       template.HTML
		Body      template.HTML
		ExtraHead template.HTML
		CSS       template.CSS
	}{
		Title:     p.Title,
		Subtitle:  p.Subtitle,
		Crumb:     p.Crumb,
		Nav:       p.Nav,
		Body:      p.Body,
		ExtraHead: p.ExtraHead,
		CSS:       template.CSS(css),
	})
}

// RenderString is the bytes.Buffer flavor of Render. Convenient when the
// caller wants to embed the rendered page inside another structure or
// inspect it before writing.
func RenderString(p Page) (string, error) {
	var b bytes.Buffer
	if err := Render(&b, p); err != nil {
		return "", err
	}
	return b.String(), nil
}

// NavItem is one entry in a sidebar nav tree. Children are rendered
// nested under the item.
type NavItem struct {
	Label    string
	Href     string
	Subtitle string
	Children []NavItem
}

// BuildNavList renders a flat or nested list of nav items as HTML
// suitable for the Page.Nav field. The output uses <ul> / <li> elements
// styled by DefaultCSS.
func BuildNavList(items []NavItem) template.HTML {
	var b bytes.Buffer
	writeList(&b, items)
	return template.HTML(b.String()) //nolint:gosec // emits well-formed HTML from typed input
}

func writeList(b *bytes.Buffer, items []NavItem) {
	if len(items) == 0 {
		return
	}
	b.WriteString(`<ul>`)
	for _, it := range items {
		b.WriteString(`<li>`)
		if it.Href != "" {
			fmt.Fprintf(b, `<a href="%s">%s</a>`,
				template.HTMLEscapeString(it.Href),
				template.HTMLEscapeString(it.Label),
			)
		} else {
			b.WriteString(template.HTMLEscapeString(it.Label))
		}
		if it.Subtitle != "" {
			fmt.Fprintf(b, ` <span class="muted">%s</span>`, template.HTMLEscapeString(it.Subtitle))
		}
		writeList(b, it.Children)
		b.WriteString(`</li>`)
	}
	b.WriteString(`</ul>`)
}

const layoutTpl = `<!doctype html>
<html lang="en">
<head>
  <meta charset="utf-8">
  <meta name="viewport" content="width=device-width, initial-scale=1">
  <title>{{.Title}}</title>
  <style>{{.CSS}}</style>
  {{.ExtraHead}}
</head>
<body>
  {{if .Crumb}}<p class="muted crumb">{{.Crumb}}</p>{{end}}
  <h1>{{.Title}}</h1>
  {{if .Subtitle}}<p class="muted subtitle">{{.Subtitle}}</p>{{end}}

  {{if .Nav}}
  <nav class="sidebar">
    <strong>Commands</strong>
    {{.Nav}}
  </nav>
  {{end}}

  <main>{{.Body}}</main>
</body>
</html>
`
