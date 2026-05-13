package webdocs_test

import (
	"fmt"
	"os"
	"path/filepath"

	webdocs "github.com/coilysiren/cli-web-docs"
	"github.com/urfave/cli/v3"
)

// Render the documentation for a *cli.Command into a directory.
// Single-page output is the default; pass Options.PerPage = true for one
// HTML file per visible subcommand.
func ExampleRender() {
	app := &cli.Command{
		Name:  "demo",
		Usage: "tiny demo cli",
		Commands: []*cli.Command{
			{Name: "hello", Usage: "say hello"},
		},
	}

	dir, _ := os.MkdirTemp("", "webdocs-example-*")
	defer os.RemoveAll(dir)

	if err := webdocs.Render(app, webdocs.Options{
		OutputDir: dir,
		Title:     "Demo Docs",
	}); err != nil {
		fmt.Println("render:", err)
		return
	}

	_, err := os.Stat(filepath.Join(dir, "index.html"))
	fmt.Println("index.html exists:", err == nil)
	// Output: index.html exists: true
}
