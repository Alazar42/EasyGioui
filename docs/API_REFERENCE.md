# API Reference

Complete Go API documentation for EasyGioUI.

## Package: `easygioui`

Main package for working with EasyGioUI.

```go
import easygioui "easygioui/sdk/easygio"
```

## Functions

### Load

Loads and parses a `.easy` UI file, returns a cached UI instance.

```go
func Load(path string) *UI
```

**Parameters:**
- `path` (string) - Path to the `.easy` file relative to the current working directory

**Returns:**
- `*UI` - Cached UI instance, or nil if parsing fails

**Example:**
```go
ui := easygioui.Load("app/ui.easy")
if ui == nil {
    log.Fatal("Failed to load UI")
}
```

**Notes:**
- File is parsed once at startup, not per-frame
- AST is cached in memory
- Path can be absolute or relative
- Returns nil on parse errors

---

### Register

Renders the UI for a single frame. Call this once per frame in your event handler.

```go
func Register(ops *op.Ops, gtx layout.Context, ui *UI)
```

**Parameters:**
- `ops` (*op.Ops) - Gio operations buffer
- `gtx` (layout.Context) - Current layout context
- `ui` (*UI) - UI instance from Load()

**Example:**
```go
if frame, ok := e.(app.FrameEvent); ok {
    gtx := layout.Context{
        Ops:         &ops,
        Now:         frame.Now,
        Metric:      frame.Metric,
        Source:      frame.Source,
        Constraints: layout.Exact(frame.Size),
    }
    
    easygioui.Register(&ops, gtx, ui)
    frame.Frame(&ops)
}
```

**Notes:**
- Must be called every frame
- Handles button clicks and text rendering
- Operations are added to the ops buffer

---

### Bind

Binds a Go struct for event handler execution. Handler methods are called on button clicks.

```go
func Bind(app interface{})
```

**Parameters:**
- `app` (interface{}) - Application struct instance

**Example:**
```go
type MyApp struct {
    count int
}

func (a *MyApp) Increment() {
    a.count++
    easygioui.SetText("counter", fmt.Sprintf("%d", a.count))
}

app := &MyApp{}
easygioui.Bind(app)
```

**Notes:**
- Should be called before each Register() call
- Handler methods must be public (capital letter)
- Handlers must take no arguments
- Struct type name must match handler references exactly

---

### SetText

Updates the text content of a Text component by ID.

```go
func SetText(id string, value interface{}) error
```

**Parameters:**
- `id` (string) - Element ID from the `.easy` file
- `value` (interface{}) - New text value (converted to string)

**Returns:**
- `error` - Returns error if element not found

**Example:**
```go
// In .easy file:
// Text { id: "counter" text: "0" }

// In Go code:
easygioui.SetText("counter", "42")
easygioui.SetText("counter", fmt.Sprintf("Count: %d", 100))
```

**Supported Types:**
- `string` - Used directly
- `int`, `int64`, etc. - Converted to string
- `float32`, `float64` - Converted to string
- `bool` - "true" or "false"
- Any type with `String()` method

**Notes:**
- Element must have an `id` in the `.easy` file
- Overrides the static text value
- Persists until explicitly changed
- Updates are rendered next frame

---

## Types

### UI

Represents a parsed and cached EasyGioUI file.

```go
type UI struct {
    // Private fields
}
```

**Usage:**
```go
ui := easygioui.Load("app/ui.easy")
// Use ui with Register()
easygioui.Register(&ops, gtx, ui)
```

### Handler Signature

Handlers bound to buttons must have this exact signature:

```go
func (a *AppType) HandlerName()
```

**Requirements:**
- Public method (capital letter)
- No parameters
- No return values
- Receiver is the bound app struct

**Valid:**
```go
func (a *Counter) Increment() { }
func (a *App) OnClick() { }
func (a *MyApp) Process() { }
```

**Invalid:**
```go
func (a *Counter) increment() { }              // Private
func (a *Counter) Increment(x int) { }         // Has parameters
func (a *Counter) Increment() error { }        // Has returns
```

---

## Complete Example

```go
package main

import (
    "fmt"
    "gioui.org/app"
    "gioui.org/layout"
    "gioui.org/op"
    easygioui "easygioui/sdk/easygio"
)

type Counter struct {
    value int
}

func (c *Counter) Increment() {
    c.value++
    easygioui.SetText("display", fmt.Sprintf("%d", c.value))
}

func (c *Counter) Decrement() {
    c.value--
    easygioui.SetText("display", fmt.Sprintf("%d", c.value))
}

func (c *Counter) Reset() {
    c.value = 0
    easygioui.SetText("display", "0")
}

func main() {
    // Load UI
    ui := easygioui.Load("app/ui.easy")
    if ui == nil {
        panic("Failed to load UI")
    }

    // Create app instance
    counter := &Counter{}

    // Run Gio event loop
    go func() {
        w := app.NewWindow()
        w.Option(app.Title("Counter"))

        var ops op.Ops
        for e := range w.Events() {
            if frame, ok := e.(app.FrameEvent); ok {
                gtx := layout.Context{
                    Ops:         &ops,
                    Now:         frame.Now,
                    Metric:      frame.Metric,
                    Source:      frame.Source,
                    Constraints: layout.Exact(frame.Size),
                }

                // Bind app and render UI
                easygioui.Bind(counter)
                easygioui.Register(&ops, gtx, ui)

                frame.Frame(&ops)
            }
        }
    }()

    app.Main()
}
```

**Corresponding `ui.easy`:**
```
Window {
    VBox {
        spacing: "8"
        
        Text {
            id: display
            text: "0"
            style: {
                size: "28"
                textColor: "white"
                bgColor: "black"
            }
        }
        
        HBox {
            weight: "1"
            spacing: "4"
            
            Button {
                text: "-"
                onClick: Counter.Decrement
                weight: "1"
            }
            
            Button {
                text: "Reset"
                onClick: Counter.Reset
                weight: "1"
            }
            
            Button {
                text: "+"
                onClick: Counter.Increment
                weight: "1"
            }
        }
    }
}
```

---

## Error Handling

### Load() Errors

Errors during parsing are logged but not returned. Check if ui is nil:

```go
ui := easygioui.Load("app/ui.easy")
if ui == nil {
    log.Fatal("Parser error - check console output")
}
```

### SetText() Errors

Element not found errors are returned:

```go
err := easygioui.SetText("nonexistent", "value")
if err != nil {
    log.Printf("Element not found: %v", err)
}
```

---

## Best Practices

**Do:**
- Load UI once at startup
- Bind app on each frame
- Use SetText() for dynamic updates
- Keep handlers lightweight
- Use meaningful element IDs

**Don't:**
- Reload UI every frame
- Create new app instances constantly
- Call SetText() with non-existent IDs
- Put heavy logic in handlers
- Use inconsistent struct/method names

---

## Performance Characteristics

| Operation | Time | Frequency |
|-----------|------|-----------|
| Load() | ~1-5ms | Once at startup |
| Bind() | <1ms | Per frame |
| Register() | ~0.5-2ms | Per frame |
| SetText() | <1ms | On demand |
| Button Click | Handle in handler | User action |

---

## See Also

- [Getting Started](GETTING_STARTED.md) - Quick start guide
- [Components Reference](COMPONENTS.md) - All UI elements
- [Styling Guide](STYLING.md) - Colors and styling
- [Layout System](LAYOUT.md) - Responsive layouts
- [examples/](../examples/) - Working code examples
