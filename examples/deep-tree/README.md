# deep-tree example

Three levels of nested subcommands with realistic flags, aliases, and arg shapes. Mirrors the shape of a small homelab ops CLI to show how cli-web-docs scales past the "tiny demo" case.

```
$ go run ./examples/deep-tree ./site
$ open ./site/index.html
```

Renders in multi-page mode by default because the surface is large enough that single-page would be hard to navigate. Switch the `PerPage` field to `false` in `main.go` to see the same tree as one big page.
