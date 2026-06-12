# EasyGioUI - Pure Library Implementation

EasyGioUI is now a pure library that integrates directly into Gio applications without requiring a custom runtime, engine, or app wrapper.

## API Reference

### `Load(path string) *UI`
Loads and parses a `.easy` file once at startup. The AST is cached in memory for efficient rendering across frames.

```go
ui := easygioui.Load("app/ui.easy")
if ui == nil {
    panic("failed to load UI")
}
```

### `Register(ops *op.Ops, gtx layout.Context, ui *UI)`
Renders the UI tree inside each `FrameEvent`. Must be called once per frame.

```go
case app.FrameEvent:
    gtx := layout.Context{
        Ops: &ops,
        Now: e.Now,
        Metric: e.Metric,
        Source: e.Source,
    }
    gtx.Constraints = layout.Exact(e.Size)
    
    easygioui.Register(&ops, gtx, ui)
    e.Frame(&ops)
```

### `SetText(nodeID string, value interface{})`
Updates the text value for a node by ID. Persists across frames and overrides the initial UI definition.

```go
easygioui.SetText("counterText", fmt.Sprintf("Count: %d", count))
```

### `Bind(app interface{})`
Registers an application instance for event binding. Only zero-argument methods are supported as event handlers.

```go
type App struct {
    count int
}

func (a *App) Increment() {
    a.count++
    easygioui.SetText("counter", fmt.Sprintf("Count: %d", a.count))
}

app := &App{}
easygioui.Bind(app)
```

## Supported UI Elements

### Layouts
- **Window**: Root container (logical only, not rendered)
- **VBox**: Vertical flex layout
- **HBox**: Horizontal flex layout

### Widgets
- **Text**: Text label
- **Button**: Clickable button with event handler

## .easy File Format

```
Window {
    title: "My App"
    
    VBox {
        Text {
            id: myLabel
            text: "Hello"
        }
        
        Button {
            id: myButton
            text: "Click Me"
            onClick: App.OnClick
        }
    }
}
```

## Complete Example

```go
package main

import (
    "fmt"
    "os"
    
    "easygioui"
    
    "gioui.org/app"
    "gioui.org/layout"
    "gioui.org/op"
    "gioui.org/unit"
)

type App struct {
    count int
}

func (a *App) Increment() {
    a.count++
    easygioui.SetText("counter", fmt.Sprintf("Count: %d", a.count))
}

var ui *easygioui.UI

func main() {
    ui = easygioui.Load("app/ui.easy")
    if ui == nil {
        panic("failed to load UI")
    }
    
    app := &App{}
    easygioui.Bind(app)
    
    go func() {
        w := new(app.Window)
        w.Option(app.Title("Counter"), app.Size(unit.Dp(400), unit.Dp(600)))
        
        var ops op.Ops
        
        for {
            evt := w.Event()
            
            switch e := evt.(type) {
            case app.FrameEvent:
                gtx := layout.Context{
                    Ops: &ops,
                    Now: e.Now,
                    Metric: e.Metric,
                    Source: e.Source,
                }
                gtx.Constraints = layout.Exact(e.Size)
                
                easygioui.Register(&ops, gtx, ui)
                e.Frame(&ops)
                
            case app.DestroyEvent:
                os.Exit(0)
            }
        }
    }()
    
    app.Main()
}
```

## Design Principles

1. **No Runtime**: Applications use standard Gio event loop
2. **No Engine**: All rendering happens within FrameEvent handlers
3. **No Wrapper**: Developer controls window and event loop
4. **Pure Library**: EasyGioUI is a thin declarative overlay
5. **Performance**: Single parse, cached AST, minimal allocations per frame

## State Management

- Text values can be updated via `SetText()` at any time
- Updates are reflected in the next frame
- State persists across frames without re-parsing

## Event Handling

- Button clicks are captured during `Register()` execution
- Click handlers are zero-argument methods on the bound app
- Handler path format: `App.MethodName`
- Methods are called via reflection using the `Bind()` configuration
