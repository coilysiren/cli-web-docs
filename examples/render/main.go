// Command render is a tiny example that emits cli-web-docs HTML for a
// sample CLI into ./site.
package main

import (
	"fmt"
	"os"

	webdocs "github.com/coilysiren/cli-web-docs"
	"github.com/urfave/cli/v3"
)

func main() {
	app := &cli.Command{
		Name:        "demo",
		Usage:       "a tiny demo cli",
		Description: "Showcases what cli-web-docs renders for a real urfave/cli v3 application.",
		Flags: []cli.Flag{
			&cli.StringFlag{Name: "verbose", Aliases: []string{"v"}, Usage: "enable verbose output"},
		},
		Commands: []*cli.Command{
			{
				Name:      "hello",
				Usage:     "print a greeting",
				ArgsUsage: "<name>",
				Metadata:  map[string]any{"since": "v0.1.0"},
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

	out := "site"
	if len(os.Args) > 1 {
		out = os.Args[1]
	}

	if err := webdocs.Render(app, webdocs.Options{
		OutputDir:    out,
		Title:        "demo docs",
		MetadataKeys: []string{"since"},
	}); err != nil {
		fmt.Fprintln(os.Stderr, "render:", err)
		os.Exit(1)
	}
	fmt.Println("wrote", out)
}
