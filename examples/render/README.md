# render example - the basic case

Single-page rendering. A tiny `demo` CLI with three subcommands gets serialized into one `index.html` with every command inlined under section anchors.

```
$ go run ./examples/render ./site
wrote ./site
$ open ./site/index.html
```

This is the right starting point for most CLIs. Switch to [multi-page](../multi-page/) when the tree is large enough that one page becomes unwieldy.
