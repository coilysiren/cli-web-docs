// Package webdocs renders a urfave/cli v3 command tree as a static HTML
// documentation site.
//
// Pipeline:
//
//  1. urfave/cli-docs/v3 generates canonical Markdown from the command tree.
//  2. yuin/goldmark renders the Markdown as HTML (GFM tables, footnotes,
//     fenced code blocks).
//  3. webdocs wraps the rendered HTML in a layout shell: title, dark-mode-
//     aware CSS, command-tree nav, optional per-command pages.
//
// The pipeline keeps this package thin. Any improvement to cli-docs'
// Markdown shape lands here automatically; goldmark gives CommonMark + GFM
// for free.
//
// Two layout modes:
//
//   - Single-page (default): one index.html with cli-docs's full Markdown
//     for the root command (subcommands inlined) rendered into one page.
//   - Multi-page (Options.PerPage = true): one HTML file per visible
//     subcommand, plus an index.html that lists the tree.
//
// Independent of every other cli-* extension. Operates on any
// *cli.Command via the public urfave/cli API.
package webdocs

import (
	"bytes"
	"fmt"
	"html/template"
	"os"
	"path/filepath"
	"strings"

	docs "github.com/urfave/cli-docs/v3"
	"github.com/urfave/cli/v3"
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
)

// Options control the rendered output. Zero value is usable: single-page
// mode, default CSS, title taken from cmd.Name.
type Options struct {
	// OutputDir is the directory to write into. Created if it does not
	// exist. Required.
	OutputDir string

	// Title overrides the page <title>. Defaults to the root command name.
	Title string

	// PerPage = true emits one HTML file per visible subcommand at
	// "<path>.html" plus an index.html that lists the tree.
	PerPage bool

	// CSS overrides the embedded default stylesheet. Empty = use the
	// default.
	CSS string
}

// Render writes documentation for cmd into Options.OutputDir.
func Render(cmd *cli.Command, opts Options) error {
	if cmd == nil {
		return fmt.Errorf("webdocs: nil command")
	}
	if opts.OutputDir == "" {
		return fmt.Errorf("webdocs: Options.OutputDir is required")
	}
	if err := os.MkdirAll(opts.OutputDir, 0o755); err != nil {
		return fmt.Errorf("webdocs: mkdir %s: %w", opts.OutputDir, err)
	}
	title := opts.Title
	if title == "" {
		title = cmd.Name
	}
	css := opts.CSS
	if css == "" {
		css = defaultCSS
	}

	if opts.PerPage {
		return renderPerPage(cmd, title, css, opts.OutputDir)
	}
	return renderSinglePage(cmd, title, css, opts.OutputDir)
}

func md2html(md string) (template.HTML, error) {
	g := goldmark.New(goldmark.WithExtensions(extension.GFM))
	var buf bytes.Buffer
	if err := g.Convert([]byte(md), &buf); err != nil {
		return "", fmt.Errorf("webdocs: goldmark convert: %w", err)
	}
	return template.HTML(buf.String()), nil //nolint:gosec // goldmark output is HTML by intent
}

// page is the layout-template input.
type page struct {
	Title    string
	CSS      template.CSS
	Heading  string
	Crumb    template.HTML
	Body     template.HTML
	Nav      template.HTML
	IsRoot   bool
	HomeHref string
}

func renderSinglePage(cmd *cli.Command, title, css, outDir string) error {
	md, err := docs.ToMarkdown(cmd)
	if err != nil {
		return fmt.Errorf("webdocs: cli-docs ToMarkdown: %w", err)
	}
	body, err := md2html(md)
	if err != nil {
		return err
	}
	nav, err := buildNavHTML(cmd, "")
	if err != nil {
		return err
	}
	return writeLayout(filepath.Join(outDir, "index.html"), page{
		Title:   title,
		CSS:     template.CSS(css),
		Heading: title,
		Body:    body,
		Nav:     nav,
		IsRoot:  true,
	})
}

func renderPerPage(root *cli.Command, title, css, outDir string) error {
	// Index lists the tree only; the root's own Markdown is rendered
	// inline under it.
	rootMD, err := docs.ToMarkdown(root)
	if err != nil {
		return fmt.Errorf("webdocs: cli-docs ToMarkdown root: %w", err)
	}
	rootBody, err := md2html(rootMD)
	if err != nil {
		return err
	}
	nav, err := buildNavHTML(root, "")
	if err != nil {
		return err
	}
	if err := writeLayout(filepath.Join(outDir, "index.html"), page{
		Title:   title,
		CSS:     template.CSS(css),
		Heading: title,
		Body:    rootBody,
		Nav:     nav,
		IsRoot:  true,
	}); err != nil {
		return err
	}

	// One page per visible subcommand.
	var walk func(prefix []string, cmd *cli.Command) error
	walk = func(prefix []string, cmd *cli.Command) error {
		if cmd.Hidden {
			return nil
		}
		path := append(append([]string(nil), prefix...), cmd.Name)
		md, err := docs.ToMarkdown(cmd)
		if err != nil {
			return fmt.Errorf("webdocs: cli-docs ToMarkdown %s: %w", strings.Join(path, " "), err)
		}
		body, err := md2html(md)
		if err != nil {
			return err
		}
		crumb := template.HTML( //nolint:gosec
			`<a href="index.html">` + template.HTMLEscapeString(root.Name) + `</a> / ` +
				`<span class="path-segment">` + template.HTMLEscapeString(strings.Join(path, " ")) + `</span>`,
		)
		slug := strings.Join(path, "-") + ".html"
		if err := writeLayout(filepath.Join(outDir, slug), page{
			Title:    strings.Join(path, " ") + " - " + title,
			CSS:      template.CSS(css),
			Heading:  strings.Join(path, " "),
			Crumb:    crumb,
			Body:     body,
			Nav:      nav,
			HomeHref: "index.html",
		}); err != nil {
			return err
		}
		for _, sub := range cmd.Commands {
			if err := walk(path, sub); err != nil {
				return err
			}
		}
		return nil
	}
	for _, sub := range root.Commands {
		if err := walk(nil, sub); err != nil {
			return err
		}
	}
	return nil
}

// buildNavHTML renders the command tree as a <nav> block. linkBase is "" for
// single-page (anchors) or "" for multi-page (per-page slugs); we always
// produce slug links and the layout decides whether to resolve them.
//
// Single-page mode emits anchor links keyed off the heading text that
// cli-docs / goldmark produce; per-page mode emits per-file slugs.
func buildNavHTML(root *cli.Command, _ string) (template.HTML, error) {
	var b bytes.Buffer
	b.WriteString(`<ul class="tree">`)
	var walk func(prefix []string, cmd *cli.Command)
	walk = func(prefix []string, cmd *cli.Command) {
		if cmd.Hidden {
			return
		}
		path := append(append([]string(nil), prefix...), cmd.Name)
		slug := strings.Join(path, "-") + ".html"
		fmt.Fprintf(&b, `<li><a href="%s"><code>%s</code></a>`,
			template.HTMLEscapeString(slug),
			template.HTMLEscapeString(strings.Join(path, " ")),
		)
		if cmd.Usage != "" {
			fmt.Fprintf(&b, ` <span class="muted">%s</span>`, template.HTMLEscapeString(cmd.Usage))
		}
		if len(cmd.Commands) > 0 {
			b.WriteString(`<ul>`)
			for _, sub := range cmd.Commands {
				walk(path, sub)
			}
			b.WriteString(`</ul>`)
		}
		b.WriteString(`</li>`)
	}
	for _, sub := range root.Commands {
		walk(nil, sub)
	}
	b.WriteString(`</ul>`)
	return template.HTML(b.String()), nil //nolint:gosec
}

func writeLayout(path string, p page) error {
	t, err := template.New("layout").Parse(layoutTpl)
	if err != nil {
		return fmt.Errorf("webdocs: parse layout: %w", err)
	}
	f, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("webdocs: create %s: %w", path, err)
	}
	defer f.Close()
	return t.Execute(f, p)
}
