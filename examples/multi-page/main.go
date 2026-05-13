// Command multi-page renders a small CLI as a multi-page documentation
// site: one HTML file per visible subcommand plus an index.html that
// lists the tree. Use when the command surface is large enough that a
// single page becomes unwieldy to scroll.
package main

import (
	"context"
	"fmt"
	"os"

	webdocs "github.com/coilysiren/cli-web-docs"
	"github.com/urfave/cli/v3"
)

func main() {
	app := &cli.Command{
		Name:        "ops",
		Usage:       "operator CLI - multi-page docs demo",
		Description: "A medium-sized command surface where multi-page output keeps each page focused on one command.",
		Commands: []*cli.Command{
			{Name: "start", Usage: "start the service", ArgsUsage: "<service>"},
			{Name: "stop", Usage: "stop the service", ArgsUsage: "<service>"},
			{Name: "restart", Usage: "restart the service", ArgsUsage: "<service>"},
			{Name: "status", Usage: "show service status"},
			{
				Name:  "logs",
				Usage: "tail or query service logs",
				Commands: []*cli.Command{
					{Name: "tail", Usage: "tail logs in real time"},
					{Name: "since", Usage: "show logs since a given time", ArgsUsage: "<timestamp>"},
				},
			},
			{
				Name:  "config",
				Usage: "inspect or change runtime config",
				Commands: []*cli.Command{
					{Name: "get", Usage: "read a config value", ArgsUsage: "<key>"},
					{Name: "set", Usage: "write a config value", ArgsUsage: "<key> <value>"},
				},
			},
		},
	}

	out := "site"
	if len(os.Args) > 1 {
		out = os.Args[1]
	}

	if err := webdocs.Render(app, webdocs.Options{
		OutputDir: out,
		Title:     "ops docs",
		PerPage:   true,
	}); err != nil {
		fmt.Fprintln(os.Stderr, "render:", err)
		os.Exit(1)
	}
	fmt.Println("wrote", out)
	_ = context.Background()
}
