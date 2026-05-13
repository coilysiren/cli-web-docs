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
	for _, want := range []string{"Demo Docs", "demo hello", "demo tree list", "demo tree rm", "--verbose"} {
		if !strings.Contains(body, want) {
			t.Errorf("index.html missing %q", want)
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

func TestRender_Metadata(t *testing.T) {
	cmd := &cli.Command{
		Name:     "demo",
		Metadata: map[string]any{"since": "v1.2.0", "stability": "stable"},
	}
	dir := t.TempDir()
	if err := webdocs.Render(cmd, webdocs.Options{OutputDir: dir, MetadataKeys: []string{"since", "stability"}}); err != nil {
		t.Fatalf("Render: %v", err)
	}
	b, _ := os.ReadFile(filepath.Join(dir, "index.html"))
	body := string(b)
	if !strings.Contains(body, "v1.2.0") || !strings.Contains(body, "stable") {
		t.Errorf("metadata not surfaced in output. Got:\n%s", body)
	}
}
