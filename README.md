# jzero-plugins

Example plugins for the `jzero` CLI.

This repository collects standalone `jzero-*` plugin examples and keeps the usage model aligned with the official guide:

- [Custom CLI plugins](https://docs.jzero.io/guide/cli-plugin.html)

## What A CLI Plugin Is

When `jzero` receives an unknown command, it can delegate that command to an external executable instead of modifying the main `jzero` binary.

A plugin should:

- be named `jzero-<command>`
- be executable
- be available in `PATH`

Examples:

- `jzero hello` -> `jzero-hello`
- `jzero foo bar` -> first tries `jzero-foo-bar`, then falls back to `jzero-foo`

After a plugin is matched, the remaining arguments are passed to the plugin unchanged.

Plugins are discovered dynamically, so they are not part of the static built-in command list shown by `jzero --help`.

## Repository Layout

- [`jzero-hello`](./jzero-hello): a minimal Go plugin example built with Cobra

Each plugin directory is an independent Go module that produces its own `jzero-*` executable.

## Example Plugin

[`jzero-hello`](./jzero-hello) demonstrates two common patterns:

- exposing a custom `jzero hello` command
- reading parsed project metadata from `desc/api`, `desc/proto`, and `desc/sql` through `github.com/jzero-io/jzero/cmd/jzero/pkg/plugin`

Build the example:

```bash
cd jzero-hello
go build -o "$GOBIN/jzero-hello" .
```

Make sure `$GOBIN` is set and included in `PATH`.

Use the plugin:

```bash
jzero hello --help

cd /path/to/your-jzero-project
jzero hello desc
```

The `desc` command reads the current working directory and prints any API, Proto, and SQL metadata found under `desc/`.

## Notes For Plugin Authors

- Keep the executable name aligned with the subcommand you want to expose.
- Prefer simple command names. Inside each command segment, `jzero` normalizes `-` to `_` before lookup, so `jzero my-cmd` maps to `jzero-my_cmd`.
- `plugin.New()` reads from the plugin process working directory, so it is typically used from a `jzero` project root.
- Multi-level commands are supported. For example, `jzero foo bar baz` will first try `jzero-foo-bar`, then fall back to `jzero-foo`.

## Reference

- Official guide: [https://docs.jzero.io/guide/cli-plugin.html](https://docs.jzero.io/guide/cli-plugin.html)
- Example plugin README: [./jzero-hello/README.md](./jzero-hello/README.md)
