# EasyGioUI

EasyGioUI is a retained-style UI runtime built on top of Gio (gioui.org).
It supports:

- Declarative `.easy` UI files
- Go script handlers with runtime binding
- Go-only UI construction through SDK APIs
- Event dispatch + state mutation + reactive rerender trigger model

## Architecture

Core layers:

1. UI layer (`.easy`): declarative structure tree.
2. Script layer (`.go`): behavior, events, and state mutations.
3. Runtime layer: retained tree + reactivity + event dispatch.
4. Renderer: maps retained nodes to Gio widgets.

Immediate-mode Gio is wrapped with a retained model by:

- stable node IDs
- runtime property map
- event binding registry
- tree snapshots and version bumps for diff/reconcile hooks

## Folder Structure

- `cmd/easygio`: CLI entry
- `internal/ast`: AST types
- `internal/parser`: lexer + parser for `.easy`
- `internal/loader`: unified app loader
- `internal/runtime`: retained tree + runtime context
- `internal/binder`: `.easy` to Go handler binder
- `internal/events`: event dispatcher
- `internal/state`: global + scoped state store
- `internal/reactivity`: dependency graph and subscriptions
- `internal/scripts`: script loading modes
- `internal/renderer`: Gio mapping layer
- `internal/layouts`: layout wrappers
- `internal/widgets`: widget abstraction points
- `internal/hotreload`: polling hot reload loop
- `internal/devtools`: doctor command checks
- `sdk/easygio`: Go-first builder API
- `examples/counter`: mixed and Go-only usage seed

## CLI

- `easygio run -app .`
- `easygio dev -app .`
- `easygio build -app . -o easygio-app`
- `easygio create -name myapp`
- `easygio doctor`

## Declarative `.easy` Example

```easy
Window {
    title: "App"
    VBox {
        Text {
            id: titleText
            value: "Hello"
        }
        Button {
            id: btn
            text: "Click"
            onClick: App.OnClick
        }
    }
}
```

## Go Script Handler Contract

Handlers are bound by reflection using the reference syntax `Script.Method`.
Current required signature:

```go
func (a *App) OnClick(ctx binder.EasyContext) error
```

## Go-only UI Mode

```go
app := easygio.New()
ui := app.Window("App", app.VBox(
    app.Text("titleText", "Hello"),
    app.Button("btn", "Click", "App.OnClick"),
))
tree := ui.BuildTree()
```

## Current Status

This repository provides a production-oriented foundation:

- parser/AST and retained tree are implemented
- runtime context supports `Set/Get/Emit`
- event dispatch and reflective binding are wired
- Gio renderer supports `Window`, `VBox`, `HBox`, `Text`, `Button`
- hot reload loop scaffolded for dev mode

Next extensions:

- real diff/reconciliation patches
- file watcher based hot reload and script rebind
- build-mode code generation from `.easy`
- richer widget catalog and style system
