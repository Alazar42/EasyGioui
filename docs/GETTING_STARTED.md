# Getting Started with EasyGioUI

This guide will help you set up your first EasyGioUI application.

## Installation

```bash
go get easygioui
```

## Prerequisites

- Go 1.20 or higher
- Gio library: `go get gioui.org`

## Basic Setup

### 1. Create Your First `.easy` UI File

Create a file `app/ui.easy`:

```
Window {
    VBox {
        spacing: "8"
        
        Text {
            text: "Hello, EasyGioUI!"
            style: {
                size: "24"
                textColor: "white"
                bgColor: "blue"
            }
        }
        
        Button {
            text: "Click Me"
            onClick: App.OnClick
        }
    }
}
```

### 2. Create Your Go Application

Create `main.go`:

```go
package main

import (
	"fmt"
	"gioui.org/app"
	"gioui.org/layout"
	"gioui.org/op"
	easygioui "easygioui"
)

// Define your application struct
type MyApp struct {
	clickCount int
}

// Define handler methods (must be public)
func (a *MyApp) OnClick() {
	a.clickCount++
	fmt.Printf("Button clicked %d times\n", a.clickCount)
}

func main() {
	// Create UI instance
	ui := easygioui.Load("app/ui.easy")
	if ui == nil {
		panic("Failed to load UI")
	}

	// Create app instance
	myApp := &MyApp{}

	// Run Gio event loop
	go func() {
		w := app.NewWindow()
		w.Option(app.Title("My App"))

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

				// Bind your app and render UI
				easygioui.Bind(myApp)
				easygioui.Register(&ops, gtx, ui)

				frame.Frame(&ops)
			}
		}
	}()

	app.Main()
}
```

### 3. Run Your Application

```bash
go run main.go
```

## Common Patterns

### Updating Text Dynamically

Use `SetText()` to update label content:

```go
func (a *MyApp) UpdateCounter() {
	a.count++
	easygioui.SetText("counterLabel", fmt.Sprintf("Count: %d", a.count))
}
```

In `.easy`:
```
Text {
	id: counterLabel
	text: "Count: 0"
}

Button {
	text: "Increment"
	onClick: App.UpdateCounter
}
```

### Multiple Sections

Organize your UI with nested VBox/HBox:

```
Window {
	VBox {
		spacing: "12"
		
		VBox {
			style: { bgColor: "green" }
			Text { text: "Section 1" }
			Button { text: "Action 1" onClick: App.Action1 }
		}
		
		VBox {
			style: { bgColor: "red" }
			Text { text: "Section 2" }
			Button { text: "Action 2" onClick: App.Action2 }
		}
	}
}
```

### Responsive Grid

Create grid-like layouts with flex weights:

```
Window {
	VBox {
		HBox {
			weight: "1"
			spacing: "4"
			Button { weight: "1" text: "1" }
			Button { weight: "1" text: "2" }
			Button { weight: "1" text: "3" }
		}
		HBox {
			weight: "1"
			spacing: "4"
			Button { weight: "1" text: "4" }
			Button { weight: "1" text: "5" }
			Button { weight: "1" text: "6" }
		}
	}
}
```

## Handler Methods

Handler methods must:
- Be **public** (capital letter)
- Take **no arguments**
- Return nothing (or `error` for error handling)

Valid:
```go
func (a *App) Increment() { }
func (a *App) OnButtonClick() error { return nil }
```

Invalid:
```go
func (a *App) private() { }              // private
func (a *App) Handler(x int) { }         // has arguments
func (a *App) Process() (int, error) { } // returns values
```

## Debugging

### Check UI File Loads Correctly

```go
ui := easygioui.Load("app/ui.easy")
if ui == nil {
	log.Fatalf("Failed to load UI")
}
```

### Verify Handler References

Make sure handler syntax matches: `App.MethodName`

The first part should match your bound struct type exactly.

### Common Issues

| Issue | Solution |
|-------|----------|
| "Handler not found" | Check method is public and name matches exactly |
| Blank screen | Verify `Window` block exists and has children |
| Layout issues | Check `weight` and `spacing` values |
| Text not updating | Confirm `id` matches in `SetText()` call |

## Next Steps

- Read [Components Reference](COMPONENTS.md) for all widget types
- Check [Styling Guide](STYLING.md) for colors and styling options
- Explore [Layout System](LAYOUT.md) for responsive design
- Review provided [examples](../examples) for working code

## Example Projects

- **Calculator** - `examples/calculator/` - Full-featured calculator app
- **Counter** - `examples/counter/` - Simple counter with multiple sections
