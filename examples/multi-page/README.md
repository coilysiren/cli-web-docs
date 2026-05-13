# multi-page example

Same shape as `render/` but with `Options.PerPage = true`. Each visible subcommand gets its own `<name>.html`; the index lists the tree.

```
$ go run ./examples/multi-page ./site
wrote ./site
$ ls site/
config-get.html  config.html  logs-since.html  logs.html  restart.html  status.html  stop.html
config-set.html  index.html   logs-tail.html   ops.html   start.html
```

Right for CLIs with more than ~10 commands or deep nesting where a single page becomes hard to navigate.
