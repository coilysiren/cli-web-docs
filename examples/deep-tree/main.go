// Command deep-tree renders documentation for a CLI with three levels
// of nested subcommands and a realistic flag surface. Demonstrates that
// cli-web-docs handles deeply-nested trees without truncation and shows
// how flags / args / descriptions render in practice.
//
// Mirrors the shape of a real homelab ops CLI (think: a small fraction
// of coily). Use this as a comparison point when sizing a docs site for
// your own command tree.
package main

import (
	"fmt"
	"os"

	webdocs "github.com/coilysiren/cli-web-docs"
	"github.com/urfave/cli/v3"
)

func main() {
	app := &cli.Command{
		Name:        "homelab",
		Usage:       "operator CLI for a single-node homelab",
		Description: "Three-level command tree exercising the rendering surface end to end. The actual implementation is irrelevant - what matters is the shape of the rendered docs.",
		Flags: []cli.Flag{
			&cli.StringFlag{Name: "config", Aliases: []string{"c"}, Usage: "path to config file", Value: "~/.homelab.yaml"},
			&cli.BoolFlag{Name: "verbose", Aliases: []string{"v"}, Usage: "enable debug logging"},
			&cli.StringFlag{Name: "context", Usage: "named cluster context", Value: "default"},
		},
		Commands: []*cli.Command{
			{
				Name:  "cluster",
				Usage: "manage the homelab cluster",
				Commands: []*cli.Command{
					{
						Name:  "node",
						Usage: "operate on a single node",
						Commands: []*cli.Command{
							{
								Name:      "drain",
								Usage:     "cordon and drain a node before maintenance",
								ArgsUsage: "<node-name>",
								Flags: []cli.Flag{
									&cli.DurationFlag{Name: "timeout", Usage: "max time to wait for pods to relocate", Value: 0},
									&cli.BoolFlag{Name: "force", Usage: "evict pods with PDB violations"},
								},
							},
							{
								Name:      "uncordon",
								Usage:     "mark a node schedulable again",
								ArgsUsage: "<node-name>",
							},
							{
								Name:      "labels",
								Usage:     "show or set node labels",
								ArgsUsage: "<node-name> [key=value]...",
							},
						},
					},
					{
						Name:  "pod",
						Usage: "operate on individual pods",
						Commands: []*cli.Command{
							{Name: "ls", Usage: "list pods, optionally filtered by namespace", Flags: []cli.Flag{&cli.StringFlag{Name: "namespace", Aliases: []string{"n"}, Usage: "namespace to scope to"}}},
							{Name: "logs", Usage: "tail logs for a pod", ArgsUsage: "<pod>", Flags: []cli.Flag{&cli.BoolFlag{Name: "follow", Aliases: []string{"f"}, Usage: "stream live logs"}}},
							{Name: "exec", Usage: "run a one-shot command inside a pod", ArgsUsage: "<pod> -- <argv...>"},
						},
					},
				},
			},
			{
				Name:  "deploy",
				Usage: "deploy and roll out services",
				Commands: []*cli.Command{
					{
						Name:      "service",
						Usage:     "deploy a named service from the manifest repo",
						ArgsUsage: "<service>",
						Flags: []cli.Flag{
							&cli.StringFlag{Name: "version", Usage: "image tag or git ref to deploy"},
							&cli.BoolFlag{Name: "dry-run", Usage: "render the manifest but skip apply"},
						},
					},
					{
						Name:      "rollback",
						Usage:     "rollback a service to a previous revision",
						ArgsUsage: "<service> <revision>",
					},
				},
			},
			{
				Name:  "obs",
				Usage: "observability stack helpers",
				Commands: []*cli.Command{
					{Name: "tail", Usage: "tail aggregated logs"},
					{Name: "query", Usage: "PromQL helper"},
					{Name: "dashboards", Usage: "list deployed Grafana dashboards"},
				},
			},
			{
				Name:  "backup",
				Usage: "kick off ad-hoc backups (cron handles the scheduled path)",
				Flags: []cli.Flag{
					&cli.StringFlag{Name: "target", Usage: "named backup target", Value: "all"},
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
		Title:     "homelab CLI docs",
		PerPage:   true,
	}); err != nil {
		fmt.Fprintln(os.Stderr, "render:", err)
		os.Exit(1)
	}
	fmt.Println("wrote", out)
}
