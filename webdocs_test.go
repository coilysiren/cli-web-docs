package webdocs_test

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	webdocs "github.com/coilysiren/cli-web-docs"
	"github.com/urfave/cli/v3"
)

func sampleCmd() *cli.Command {
	return &cli.Command{
		Name:  "demo",
		Usage: "a tiny demo cli",
		Flags: []cli.Flag{
			&cli.StringFlag{Name: "verbose", Aliases: []string{"v"}, Usage: "enable verbose output"},
		},
		Commands: []*cli.Command{
			{
				Name:      "hello",
				Usage:     "print a greeting",
				ArgsUsage: "<name>",
			},
			{
				Name:  "tree",
				Usage: "subcommand group",
				Commands: []*cli.Command{
					{Name: "list", Usage: "list things"},
					{Name: "rm", Usage: "remove a thing", ArgsUsage: "<id>"},
				},
			},
		},
	}
}

func TestRender_SinglePage(t *testing.T) {
	dir := t.TempDir()
	if err := webdocs.Render(sampleCmd(), webdocs.Options{OutputDir: dir, Title: "Demo Docs"}); err != nil {
		t.Fatalf("Render: %v", err)
	}
	b, err := os.ReadFile(filepath.Join(dir, "index.html"))
	if err != nil {
		t.Fatalf("read index: %v", err)
	}
	body := string(b)
	for _, want := range []string{"Demo Docs", "hello", "tree", "verbose"} {
		if !strings.Contains(body, want) {
			t.Errorf("index.html missing %q", want)
		}
	}
	// Nav lists each command with its dotted slug as the link target.
	for _, slug := range []string{"hello.html", "tree.html", "tree-list.html", "tree-rm.html"} {
		if !strings.Contains(body, slug) {
			t.Errorf("nav missing slug %q", slug)
		}
	}
}

func TestRender_PerPage(t *testing.T) {
	dir := t.TempDir()
	if err := webdocs.Render(sampleCmd(), webdocs.Options{OutputDir: dir, Title: "Demo Docs", PerPage: true}); err != nil {
		t.Fatalf("Render: %v", err)
	}
	wantFiles := []string{"index.html", "hello.html", "tree.html", "tree-list.html", "tree-rm.html"}
	for _, f := range wantFiles {
		if _, err := os.Stat(filepath.Join(dir, f)); err != nil {
			t.Errorf("missing %s: %v", f, err)
		}
	}
	// Per-page file should mention the command name.
	b, _ := os.ReadFile(filepath.Join(dir, "hello.html"))
	if !strings.Contains(string(b), "hello") {
		t.Errorf("hello.html does not mention 'hello'")
	}
	// Crumb back to index is present on per-page files.
	if !strings.Contains(string(b), `href="index.html"`) {
		t.Errorf("hello.html missing crumb back to index")
	}
}

func TestRender_NoOutputDir(t *testing.T) {
	if err := webdocs.Render(sampleCmd(), webdocs.Options{}); err == nil {
		t.Errorf("expected error on empty OutputDir")
	}
}

func TestRender_NilCommand(t *testing.T) {
	if err := webdocs.Render(nil, webdocs.Options{OutputDir: t.TempDir()}); err == nil {
		t.Errorf("expected error on nil command")
	}
}
