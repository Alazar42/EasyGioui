# EasyGioUI

EasyGioUI is a minimal declarative UI library that plugs directly into Gio applications without replacing the runtime.

## What It Is

- A pure library (not a framework)
- Parses `.easy` declarative UI files once at startup
- Renders UI into standard Gio `FrameEvent` handlers
- No custom runtime, no engine, no app wrapper
- Direct injection into your Gio event loop

## What It's Not

- Not a Gio replacement
- Not a custom event loop
- Not an opinionated runtime architecture
- Not auto-generated code

## Core API

Only 4 public functions:

1. **`Load(path string) *UI`** - Parse and cache `.easy` file
2. **`Register(ops *op.Ops, gtx layout.Context, ui *UI)`** - Render in FrameEvent
3. **`SetText(id string, value any)`** - Update text values
4. **`Bind(app any)`** - Bind zero-arg methods to click handlers

## Usage Pattern

You write normal Gio code:

```go
var ui = easygioui.Load("app/ui.easy")

func main() {
    go func() {
        w := new(app.Window)
        w.Option(app.Title("My App"))
        
        var ops op.Ops
        for {
            evt := w.Event()
            
            if e, ok := evt.(app.FrameEvent); ok {
                gtx := layout.Context{
                    Ops: &ops,
                    Now: e.Now,
                    Metric: e.Metric,
                    Source: e.Source,
                }
                gtx.Constraints = layout.Exact(e.Size)
                
                easygioui.Register(&ops, gtx, ui)
                e.Frame(&ops)
            }
        }
    }()
    
    app.Main()
}
```

## Supported UI Elements

**.easy files support:**
- `Window` (root container, logical only)
- `VBox`, `HBox` (flex layouts)
- `Text` (labels)
- `Button` (click handlers)

Example:

```
Window {
    title: "Counter"
    
    VBox {
        Text {
            id: counter
            text: "Count: 0"
        }
        
        Button {
            text: "Increment"
            onClick: App.Increment
        }
    }
}
```

## Binding Handlers

```go
type App struct {
    count int
}

func (a *App) Increment() {
    a.count++
    easygioui.SetText("counter", fmt.Sprintf("Count: %d", a.count))
}

easygioui.Bind(&App{})
```

## How It Works

1. **Load Phase**: Parse `.easy` file once, build AST, cache in memory
2. **Per-Frame Phase**: 
   - Traverse cached AST
   - Convert nodes to Gio widgets
   - Handle button clicks
   - Apply `SetText` overrides
   - Render to `gtx.Ops`
3. **Update Phase**: Call methods on bound app, update text state

## Performance

- Single parse (no per-frame overhead)
- Cached AST reuse
- Minimal allocations in hot path
- Reflection only during setup
- Widget state cached between frames

## See Also

- [LIBRARY_API.md](LIBRARY_API.md) - Full API documentation
- [examples/counter](examples/counter) - Working example

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
