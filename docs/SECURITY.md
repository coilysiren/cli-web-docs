# Security Policy

Hello and thank you for your interest! :tada: :lock:

## Supported versions

This package is at v0. Only the latest commit on `main` is supported for security fixes - there are no published releases yet to backport to.

| Version             | Supported          |
| ------------------- | ------------------ |
| `main` (latest)     | :white_check_mark: |
| any pinned commit   | :x: (upgrade)      |

## Reporting a vulnerability

Please disclose any vulnerabilities by emailing [coilysiren@gmail.com](mailto:coilysiren@gmail.com). Expect a first response within 48 hours; follow-up cadence by email after that. This project is run on volunteer time, so please have patience :bow:

## What counts as a vulnerability

cli-web-docs writes static HTML to disk; the attack surface is small. Specifically interested in:

- HTML injection through a wrapped `cli.Command`'s description, usage, or flag-help text reaching the page without escaping
- file-write paths that escape `Options.OutputDir` (path traversal via crafted command names)
- denial-of-service via pathological command trees (unbounded recursion, deeply-nested subcommands)

Out of scope:

- the static output's hosting layer - that is the host's job (see `deploy/Caddyfile.example` for one secure posture)
- bugs in [urfave/cli-docs](https://github.com/urfave/cli-docs) or [goldmark](https://github.com/yuin/goldmark) - report upstream
