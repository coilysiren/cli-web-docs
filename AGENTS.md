# Agent instructions

Workspace-level conventions (git workflow, test/lint autonomy, readonly ops, writing voice, deploy knowledge) are loaded globally via `~/.claude/CLAUDE.md` → `agentic-os-kai/AGENTS.md`. This file covers only what's specific to this repo.

## What cli-web-docs is

Static HTML documentation generator for urfave/cli v3 command trees. Pipeline: `*cli.Command` → urfave/cli-docs Markdown → goldmark HTML → shared layout shell. Inventory: [`docs/FEATURES.md`](docs/FEATURES.md). Demos: [`examples/`](examples/).

## Dev verbs

Route through agent-guard, not bare go. The `.agent-guard/agent-guard.yaml` ↔ `Makefile` contract is checked on every CI run via `agent-guard lint`:

- `agent-guard exec build` - compile every package.
- `agent-guard exec test` - run the unit test suite.
- `agent-guard exec lint` - golangci-lint v2.12.2 with the urfave-mirrored `.golangci.yaml`.
- `agent-guard exec vet`, `tidy`, `cover` - the usual.

## Layout subpackage is a contract

`github.com/coilysiren/cli-web-docs/layout` is imported by [cli-web-ops](https://github.com/coilysiren/cli-web-ops) for visual consistency between docs and ops pages. Treat its API (Page fields, NavItem shape, DefaultCSS class names) as a quiet contract. Breaking changes here require a paired cli-web-ops bump.

## Pipeline stability

The rendering chain (urfave/cli-docs → goldmark) is intentionally thin. Do not reimplement Markdown generation; let urfave/cli-docs own it upstream. Do not reach for a different Markdown library without surveying what cli-docs ships - features added upstream land here for free.

## No JavaScript

The generated site is pure HTML + CSS. If a feature would require JS, push back: it is the wrong place. Live-execution UIs belong in [cli-web-ops](https://github.com/coilysiren/cli-web-ops), which shares the layout shell.

## Filing issues

One issue per discrete additive change. Every commit closes a same-repo issue with `closes #N`.

## See also

- [README.md](README.md) - human-facing intro.
- [docs/FEATURES.md](docs/FEATURES.md) - inventory of what ships today.
- [.agent-guard/agent-guard.yaml](.agent-guard/agent-guard.yaml) - allowlisted commands.

Cross-reference convention from [coilysiren/agentic-os#59](https://github.com/coilysiren/agentic-os/issues/59).
