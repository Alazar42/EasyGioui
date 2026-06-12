# EasyGioUI

A pure, declarative UI library for Gio applications. Write beautiful UIs in `.easy` files without replacing Gio's runtime.

[![Go Version](https://img.shields.io/badge/Go-1.24-blue.svg)](https://golang.org)
[![License](https://img.shields.io/badge/License-MIT-green.svg)](LICENSE)

## Features

**Simple & Pure** - Just a library, no framework wrapper  
**Declarative UI** - Write UI in clean `.easy` syntax  
**Zero Overhead** - Single parse, cached AST, minimal allocations  
**Direct Gio Integration** - Inject into your event loop, no runtime replacement  
**Responsive Layout** - Flex weights and responsive spacing  
**Styling System** - Colors, sizes, backgrounds for all components  
**Easy Binding** - Connect button clicks to Go methods with reflection  

## Quick Start

### 1. Define UI in `.easy` Format

Create `app/ui.easy`:
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
                text: "+"
                onClick: App.Increment
            }
            
            Button {
                text: "-"
                onClick: App.Decrement
            }
        }
    }
}
```

### 2. Load and Bind in Go

```go
package main

import (
    "fmt"
    "gioui.org/app"
    "gioui.org/op"
    "gioui.org/layout"
    easygioui "easygioui"
)

type Counter struct {
    count int
}

func (c *Counter) Increment() {
    c.count++
    easygioui.SetText("display", fmt.Sprintf("%d", c.count))
}

func (c *Counter) Decrement() {
    c.count--
    easygioui.SetText("display", fmt.Sprintf("%d", c.count))
}

func main() {
    // Load UI once at startup
    ui := easygioui.Load("app/ui.easy")
    counter := &Counter{}
    
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
                
                // Bind counter app and render UI
                easygioui.Bind(counter)
                easygioui.Register(&ops, gtx, ui)
                
                frame.Frame(&ops)
            }
        }
    }()
    
    app.Main()
}
```

## Core API

Only 4 public functions:

```go
// Load parses and caches a .easy UI file
ui := easygioui.Load("app/ui.easy")

// Bind connects a Go struct for event handlers
easygioui.Bind(&myApp)

// Register renders the UI in a frame
easygioui.Register(&ops, gtx, ui)

// SetText updates the text of a component by ID
easygioui.SetText("myLabel", "New Text")
```

## Supported Components

| Component | Purpose | Example |
|-----------|---------|---------|
| `Window` | Root container | `Window { ... }` |
| `VBox` | Vertical layout | `VBox { spacing: "4" ... }` |
| `HBox` | Horizontal layout | `HBox { weight: "1" ... }` |
| `Text` | Text labels | `Text { id: "title" text: "Hello" }` |
| `Button` | Clickable buttons | `Button { text: "Click" onClick: App.Handler }` |

## Styling

All components support optional `style {}` blocks:

```
Text {
    text: "Styled Text"
    style: {
        size: "24"          // Font size
        textColor: "white"  // Text color
        bgColor: "blue"     // Background color
    }
}

Button {
    text: "Click Me"
    style: {
        bgColor: "green"    // Button background
    }
}
```

### Available Colors

`red`, `green`, `blue`, `white`, `black`, `yellow`, `cyan`, `magenta`, `orange`, `purple`, `gray`, `lightgray`, `darkgray`

## Responsive Layout

Use `weight` and `spacing` for responsive designs:

```
HBox {
    weight: "1"      // Takes 1 unit of parent space
    spacing: "4"     // 4 pixels between children
    
    Button {
        weight: "1"  // Expands to fill 1/3 of parent
    }
    
    Button {
        weight: "1"  // Expands to fill 1/3 of parent
    }
    
    Button {
        weight: "1"  // Expands to fill 1/3 of parent
    }
}
```

## Examples

- **[Calculator](examples/calculator)** - Full-featured calculator with responsive grid layout
- **[Counter](examples/counter)** - Simple counter app with color-coded sections

Run examples:
```bash
go run ./examples/calculator/main.go
go run ./examples/counter/main.go
```

## Documentation

- **[Getting Started](docs/GETTING_STARTED.md)** - Detailed setup guide
- **[Components Reference](docs/COMPONENTS.md)** - All component types and properties
- **[Styling Guide](docs/STYLING.md)** - Colors, sizes, and backgrounds
- **[Layout System](docs/LAYOUT.md)** - Flexbox-style responsive layouts
- **[API Reference](docs/API_REFERENCE.md)** - Complete function documentation

## Project Structure

```
.
├── cmd/easygio/              # CLI tool (dev)
├── examples/
│   ├── calculator/           # Calculator example
│   └── counter/              # Counter example
├── internal/
│   ├── ast/                  # AST node definitions
│   ├── parser/               # Lexer & parser
│   ├── renderer/             # Gio renderer
│   ├── binder/               # Event binding
│   └── ...
├── sdk/easygio/              # Public API
├── docs/                     # Documentation
└── README.md
```

## Performance

- **Single Parse**: `.easy` file parsed once at startup, not per-frame
- **Cached AST**: Reused across all frames
- **Minimal Allocations**: Hot path avoids allocations
- **Widget Caching**: Button/Text widgets cached between frames
- **Efficient Rendering**: Direct Gio operation streaming

## How It Works

1. **Load Phase**: Parse `.easy` file into AST, cache in memory
2. **Per-Frame Phase**: 
   - Traverse cached AST
   - Create/reuse Gio widgets
   - Handle button clicks
   - Apply text updates
   - Render to ops
3. **Event Phase**: Call bound methods on state changes

## Architecture

EasyGioUI is **not a framework**. It's a library that:
- Does NOT replace Gio's event loop
- Does NOT manage your app state
- Does NOT require special initialization
- Plugs directly into standard Gio code

You maintain full control over your application architecture.

## Contributing

Contributions are welcome! Areas for enhancement:
- Additional widget types (Input, Checkbox, Dropdown, etc.)
- More styling options (borders, shadows, etc.)
- Animation/transition system
- Hot reload with file watcher
- Build-time code generation

## License

MIT License - see [LICENSE](LICENSE) for details

## Acknowledgments

Built with [Gio](https://gioui.org) - Immediate mode GUI in Go
