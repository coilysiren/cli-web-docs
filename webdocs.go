// Package webdocs renders a urfave/cli v3 command tree as static HTML
// documentation. The output is plain HTML with embedded CSS and no
// JavaScript - serve it from any static host, ship it in a release
// artifact, point a browser at it offline.
//
// Two layout modes:
//
//   - Single-page (Options.PerPage = false, default): one index.html with
//     every command inlined under a section anchor.
//   - Multi-page (Options.PerPage = true): one HTML file per command at
//     "<path>.html", index.html lists the tree.
//
// Independent of every other cli-* extension. Operates on any
// *cli.Command via the public urfave/cli API.
package webdocs

import (
	"fmt"
	"html/template"
	"os"
	"path/filepath"
	"strings"

	"github.com/urfave/cli/v3"
)

// Options control the rendered output. Zero value is usable: single-page
// mode, no metadata surfaced, default CSS, title taken from cmd.Name.
type Options struct {
	// OutputDir is the directory to write into. Created if it does not
	// exist. Required.
	OutputDir string

	// Title overrides the page <title>. Defaults to the root command name.
	Title string

	// PerPage = true emits one HTML file per command at "<path>.html"
	// instead of a single index.html.
	PerPage bool

	// MetadataKeys are cli.Command.Metadata keys to surface in the
	// rendered docs. Order is preserved. Unknown keys are skipped silently.
	MetadataKeys []string

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

	root := buildNode(cmd, nil)

	if opts.PerPage {
		return renderPerPage(root, title, css, opts.MetadataKeys, opts.OutputDir)
	}
	return renderSinglePage(root, title, css, opts.MetadataKeys, opts.OutputDir)
}

// Node is the in-memory shape of a command in the docs tree.
type Node struct {
	Name        string
	Path        []string // ["myapp", "sub", "leaf"]
	Usage       string
	Description string
	ArgsUsage   string
	Flags       []FlagDoc
	Subcommands []*Node
	Metadata    map[string]any
}

// FlagDoc is the rendered shape of a flag.
type FlagDoc struct {
	Name    string
	Aliases []string
	Usage   string
	Default string
}

func buildNode(cmd *cli.Command, parentPath []string) *Node {
	path := append(append([]string(nil), parentPath...), cmd.Name)
	n := &Node{
		Name:        cmd.Name,
		Path:        path,
		Usage:       cmd.Usage,
		Description: cmd.Description,
		ArgsUsage:   cmd.ArgsUsage,
		Metadata:    cmd.Metadata,
	}
	for _, f := range cmd.Flags {
		n.Flags = append(n.Flags, flagDoc(f))
	}
	for _, sub := range cmd.Commands {
		if sub.Hidden {
			continue
		}
		n.Subcommands = append(n.Subcommands, buildNode(sub, path))
	}
	return n
}

func flagDoc(f cli.Flag) FlagDoc {
	d := FlagDoc{}
	names := f.Names()
	if len(names) > 0 {
		d.Name = names[0]
		if len(names) > 1 {
			d.Aliases = names[1:]
		}
	}
	if df, ok := f.(cli.DocGenerationFlag); ok {
		d.Usage = df.GetUsage()
		d.Default = df.GetDefaultText()
	}
	return d
}

// Anchor returns the HTML id used for a node within a single-page render.
func (n *Node) Anchor() string {
	return strings.Join(n.Path, "-")
}

// Slug returns the per-page filename for a node ("path-to-cmd.html").
func (n *Node) Slug() string {
	if len(n.Path) <= 1 {
		return "index.html"
	}
	return strings.Join(n.Path[1:], "-") + ".html"
}

type pageData struct {
	Title        string
	CSS          template.CSS
	Root         *Node
	MetadataKeys []string
	PerPage      bool
	// Current is the node being rendered on a per-page emission.
	Current *Node
}

func renderSinglePage(root *Node, title, css string, mdKeys []string, outDir string) error {
	t, err := template.New("page").Funcs(tplFuncs).Parse(singlePageTpl)
	if err != nil {
		return fmt.Errorf("webdocs: parse template: %w", err)
	}
	f, err := os.Create(filepath.Join(outDir, "index.html"))
	if err != nil {
		return fmt.Errorf("webdocs: create index.html: %w", err)
	}
	defer f.Close()
	return t.Execute(f, pageData{
		Title:        title,
		CSS:          template.CSS(css),
		Root:         root,
		MetadataKeys: mdKeys,
	})
}

func renderPerPage(root *Node, title, css string, mdKeys []string, outDir string) error {
	t, err := template.New("page").Funcs(tplFuncs).Parse(perPageTpl)
	if err != nil {
		return fmt.Errorf("webdocs: parse template: %w", err)
	}
	// Index page lists the tree.
	idx, err := template.New("index").Funcs(tplFuncs).Parse(indexTpl)
	if err != nil {
		return fmt.Errorf("webdocs: parse index template: %w", err)
	}
	f, err := os.Create(filepath.Join(outDir, "index.html"))
	if err != nil {
		return fmt.Errorf("webdocs: create index.html: %w", err)
	}
	if err := idx.Execute(f, pageData{
		Title:        title,
		CSS:          template.CSS(css),
		Root:         root,
		MetadataKeys: mdKeys,
		PerPage:      true,
	}); err != nil {
		f.Close()
		return err
	}
	f.Close()

	// Walk and render each non-root node.
	return walk(root, func(n *Node) error {
		if len(n.Path) == 1 {
			return nil
		}
		out, err := os.Create(filepath.Join(outDir, n.Slug()))
		if err != nil {
			return fmt.Errorf("webdocs: create %s: %w", n.Slug(), err)
		}
		defer out.Close()
		return t.Execute(out, pageData{
			Title:        title,
			CSS:          template.CSS(css),
			Root:         root,
			MetadataKeys: mdKeys,
			PerPage:      true,
			Current:      n,
		})
	})
}

func walk(n *Node, fn func(*Node) error) error {
	if err := fn(n); err != nil {
		return err
	}
	for _, c := range n.Subcommands {
		if err := walk(c, fn); err != nil {
			return err
		}
	}
	return nil
}

var tplFuncs = template.FuncMap{
	"join": func(parts []string, sep string) string {
		return strings.Join(parts, sep)
	},
	"hasMetadata": func(md map[string]any, keys []string) bool {
		for _, k := range keys {
			if _, ok := md[k]; ok {
				return true
			}
		}
		return false
	},
	"metaValue": func(md map[string]any, k string) any {
		return md[k]
	},
}
