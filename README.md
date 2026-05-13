# cli-web-docs

[![Go Reference][goreference_badge]][goreference_link]
[![Go Report Card][goreportcard_badge]][goreportcard_link]
[![Tests status][test_badge]][test_link]

cli-web-docs is a **static HTML documentation generator** for [urfave/cli][urfave/cli] v3 command trees, featuring:

- canonical Markdown from [urfave/cli-docs][cli-docs], rendered to HTML by [goldmark][goldmark]
- single-page or per-command output via `Options.PerPage`
- shared layout subpackage (also consumed by [cli-web-ops]) so docs + ops pages read as one site
- dark-mode-aware default stylesheet, fully overridable via `Options.CSS`
- pure Go, no JavaScript in the generated output
- pairs with any static host (GitHub Pages, S3, Caddy via [`deploy/Caddyfile.example`](deploy/Caddyfile.example))

## Documentation

See [`docs/FEATURES.md`](docs/FEATURES.md) for a feature inventory and [`examples/`](examples/) for four runnable demos (`render`, `multi-page`, `custom-theme`, `deep-tree`). Local dev verbs live in [`.coily/coily.yaml`](.coily/coily.yaml); `coily lint` validates that against the [`Makefile`](Makefile).

## Support

If you found a bug or have a feature request, [create a new issue]. Participation in this community is governed by the [Code of Conduct](CODE_OF_CONDUCT.md). Security disclosures go through [SECURITY.md](SECURITY.md).

Sibling repos in the cli-* family: [cli-guard], [cli-mcp], [cli-web-ops].

### License

See [`LICENSE`](./LICENSE).

[test_badge]: https://github.com/coilysiren/cli-web-docs/actions/workflows/ci.yml/badge.svg
[test_link]: https://github.com/coilysiren/cli-web-docs/actions/workflows/ci.yml
[goreference_badge]: https://pkg.go.dev/badge/github.com/coilysiren/cli-web-docs.svg
[goreference_link]: https://pkg.go.dev/github.com/coilysiren/cli-web-docs
[goreportcard_badge]: https://goreportcard.com/badge/github.com/coilysiren/cli-web-docs
[goreportcard_link]: https://goreportcard.com/report/github.com/coilysiren/cli-web-docs
[urfave/cli]: https://github.com/urfave/cli
[cli-docs]: https://github.com/urfave/cli-docs
[goldmark]: https://github.com/yuin/goldmark
[create a new issue]: https://github.com/coilysiren/cli-web-docs/issues/new/choose
[cli-guard]: https://github.com/coilysiren/cli-guard
[cli-mcp]: https://github.com/coilysiren/cli-mcp
[cli-web-ops]: https://github.com/coilysiren/cli-web-ops
