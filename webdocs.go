// Package webdocs renders a urfave/cli v3 command tree as a static HTML
// documentation site.
//
// Pipeline:
//
//  1. urfave/cli-docs/v3 generates canonical Markdown from the command tree.
//  2. yuin/goldmark renders the Markdown as HTML (GFM tables, footnotes,
//     fenced code blocks).
//  3. webdocs wraps the rendered HTML in the shared layout shell from
//     cli-web-docs/layout (also consumed by cli-web-ops, so docs and
//     ops pages share visual identity).
//
// Two layout modes:
//
//   - Single-page (default): one index.html with cli-docs's full Markdown
//     for the root command (subcommands inlined) rendered into one page.
//   - Multi-page (Options.PerPage = true): one HTML file per visible
//     subcommand, plus an index.html that lists the tree.
//
// Independent of every other cli-* extension at the type level. Operates
// on any *cli.Command via the public urfave/cli API.
package webdocs

import (
	"bytes"
	"fmt"
	"html/template"
	"os"
	"path/filepath"
	"strings"

	"github.com/coilysiren/cli-web-docs/layout"
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
	// shared layout.DefaultCSS.
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

	if opts.PerPage {
		return renderPerPage(cmd, title, opts.CSS, opts.OutputDir)
	}
	return renderSinglePage(cmd, title, opts.CSS, opts.OutputDir)
}

func md2html(md string) (template.HTML, error) {
	g := goldmark.New(goldmark.WithExtensions(extension.GFM))
	var buf bytes.Buffer
	if err := g.Convert([]byte(md), &buf); err != nil {
		return "", fmt.Errorf("webdocs: goldmark convert: %w", err)
	}
	return template.HTML(buf.String()), nil //nolint:gosec // goldmark output is HTML by intent
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
	nav := buildNavList(cmd)
	return writePage(filepath.Join(outDir, "index.html"), layout.Page{
		Title: title,
		Nav:   nav,
		Body:  body,
		CSS:   css,
	})
}

func renderPerPage(root *cli.Command, title, css, outDir string) error {
	rootMD, err := docs.ToMarkdown(root)
	if err != nil {
		return fmt.Errorf("webdocs: cli-docs ToMarkdown root: %w", err)
	}
	rootBody, err := md2html(rootMD)
	if err != nil {
		return err
	}
	nav := buildNavList(root)

	if err := writePage(filepath.Join(outDir, "index.html"), layout.Page{
		Title: title,
		Nav:   nav,
		Body:  rootBody,
		CSS:   css,
	}); err != nil {
		return err
	}

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
		if err := writePage(filepath.Join(outDir, slug), layout.Page{
			Title: strings.Join(path, " ") + " - " + title,
			Crumb: crumb,
			Nav:   nav,
			Body:  body,
			CSS:   css,
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

// buildNavList walks the root command's children, returning a layout
// NavItem list with per-page slug hrefs.
func buildNavList(root *cli.Command) template.HTML {
	var walk func(prefix []string, cmd *cli.Command) []layout.NavItem
	walk = func(prefix []string, cmd *cli.Command) []layout.NavItem {
		if cmd.Hidden {
			return nil
		}
		path := append(append([]string(nil), prefix...), cmd.Name)
		slug := strings.Join(path, "-") + ".html"
		item := layout.NavItem{
			Label:    strings.Join(path, " "),
			Href:     slug,
			Subtitle: cmd.Usage,
		}
		for _, sub := range cmd.Commands {
			item.Children = append(item.Children, walk(path, sub)...)
		}
		return []layout.NavItem{item}
	}
	var items []layout.NavItem
	for _, sub := range root.Commands {
		items = append(items, walk(nil, sub)...)
	}
	return layout.BuildNavList(items)
}

func writePage(path string, p layout.Page) error {
	f, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("webdocs: create %s: %w", path, err)
	}
	defer f.Close()
	return layout.Render(f, p)
}
