# Components Reference

Complete reference for all components supported in `.easy` files.

## Component Overview

| Component | Type | Purpose | Container | Interactive |
|-----------|------|---------|-----------|-------------|
| `Window` | Container | Root application container | Yes | No |
| `VBox` | Container | Vertical layout | Yes | No |
| `HBox` | Container | Horizontal layout | Yes | No |
| `Text` | Display | Text labels | No | No |
| `Button` | Input | Clickable button | No | Yes |

## Window

Root container for your application. There should be exactly one `Window` per `.easy` file.

### Syntax

```
Window {
    property: "value"
    ... children ...
}
```

### Properties

| Property | Type | Default | Description |
|----------|------|---------|-------------|
| `title` | string | "" | Window title (optional) |

### Styling

```
Window {
    style: {
        bgColor: "blue"
    }
    VBox { ... }
}
```

### Example

```
Window {
    title: "My App"
    VBox {
        Text { text: "Hello" }
    }
}
```

## VBox (Vertical Box)

Arranges children vertically (top to bottom).

### Syntax

```
VBox {
    spacing: "4"
    weight: "1"
    ... children ...
}
```

### Properties

| Property | Type | Default | Description |
|----------|------|---------|-------------|
| `spacing` | number | 4 | Pixels between children |
| `weight` | number | - | Flex weight (use in parent containers) |

### Styling

```
VBox {
    style: {
        bgColor: "lightgray"
    }
}
```

### Example

```
VBox {
    spacing: "8"
    Text { text: "First" }
    Text { text: "Second" }
    Text { text: "Third" }
}
```

## HBox (Horizontal Box)

Arranges children horizontally (left to right).

### Syntax

```
HBox {
    spacing: "4"
    weight: "1"
    ... children ...
}
```

### Properties

| Property | Type | Default | Description |
|----------|------|---------|-------------|
| `spacing` | number | 4 | Pixels between children |
| `weight` | number | - | Flex weight (use in parent containers) |

### Styling

```
HBox {
    style: {
        bgColor: "gray"
    }
}
```

### Example

```
HBox {
    spacing: "4"
    Button { text: "Left" weight: "1" }
    Button { text: "Center" weight: "1" }
    Button { text: "Right" weight: "1" }
}
```

## Text

Displays text content.

### Syntax

```
Text {
    id: "myLabel"
    text: "Display this text"
    style: { ... }
}
```

### Properties

| Property | Type | Default | Description |
|----------|------|---------|-------------|
| `id` | string | - | Unique identifier for `SetText()` |
| `text` | string | "" | Text content to display |
| `value` | string | "" | Alternative to `text` |
| `weight` | number | - | Flex weight in flex containers |

### Styling

```
Text {
    text: "Styled"
    style: {
        size: "24"
        textColor: "white"
        bgColor: "blue"
    }
}
```

### Style Properties

| Property | Type | Default | Description |
|----------|------|---------|-------------|
| `size` | number | 16 | Font size in pixels |
| `textColor` | color | black | Text color |
| `bgColor` | color | - | Background color |

### Updating Text

```go
easygioui.SetText("myLabel", "New text")
```

### Example

```
Text {
    id: "counter"
    text: "Count: 0"
    style: {
        size: "20"
        textColor: "white"
        bgColor: "black"
    }
}
```

## Button

Clickable button that triggers Go methods.

### Syntax

```
Button {
    id: "myButton"
    text: "Click Me"
    onClick: App.MethodName
    weight: "1"
    style: { ... }
}
```

### Properties

| Property | Type | Required | Description |
|----------|------|----------|-------------|
| `id` | string | - | Optional identifier |
| `text` | string | Yes | Button label |
| `onClick` | handler | Yes | Handler to call on click |
| `weight` | number | - | Flex weight in containers |

### Handler Syntax

```
onClick: App.MethodName
```

Where:
- `App` is the struct type name (exactly as given)
- `MethodName` is the public method name

### Go Method

```go
func (a *App) MethodName() {
    // Handle click
}
```

### Styling

```
Button {
    text: "Styled"
    onClick: App.OnClick
    style: {
        bgColor: "green"
        textColor: "white"
        size: "18"
    }
}
```

### Style Properties

| Property | Type | Description |
|----------|------|-------------|
| `bgColor` | color | Button background color |
| `textColor` | color | Button text color |
| `size` | number | Button text size in pixels |

### Example

```
Button {
    text: "Increment"
    onClick: App.Increment
    weight: "1"
}

Button {
    text: "Decrement"
    onClick: App.Decrement
    weight: "1"
    style: {
        bgColor: "red"
    }
}
```

## Components (Custom Reusable Widgets)

Define named UI components and reuse them:

### Define Component

In your `.easy` file:

```
@component MyButton {
    Button {
        text: "MyButton"
        onClick: App.Handler
        style: {
            bgColor: "purple"
        }
    }
}
```

### Use Component

```
Window {
    VBox {
        MyButton {}
        MyButton {}
        MyButton {}
    }
}
```

### Import Components

In `components.easy`:

```
@component SuccessButton {
    Button {
        text: "Yes"
        onClick: App.OnYes
        style: {
            bgColor: "green"
        }
    }
}

@component DangerButton {
    Button {
        text: "No"
        onClick: App.OnNo
        style: {
            bgColor: "red"
        }
    }
}
```

In `main.easy`:

```
@import "components.easy"

Window {
    VBox {
        SuccessButton {}
        DangerButton {}
    }
}
```

## Element Properties Summary

### All Elements
- `style { ... }` - Optional styling

### Layout Elements (VBox, HBox)
- `spacing: "N"` - Pixels between children
- `weight: "N"` - Flex weight

### Display Elements (Text)
- `id: "name"` - Reference identifier
- `text: "value"` or `value: "value"` - Content

### Interactive Elements (Button)
- `id: "name"` - Optional identifier
- `text: "label"` - Button text
- `onClick: App.Method` - Click handler
- `weight: "N"` - Flex weight

## Styling Colors

Available for all `bgColor` and `textColor` properties:

- `red`, `green`, `blue`
- `white`, `black`
- `yellow`, `cyan`, `magenta`
- `orange`, `purple`
- `gray`, `lightgray`, `darkgray`

## Best Practices

**Do:**
- Use consistent IDs for referenced elements
- Give buttons clear, descriptive labels
- Use `weight` for responsive layouts
- Organize complex layouts with VBox/HBox nesting

**Don't:**
- Reference non-existent element IDs in `SetText()`
- Use invalid handler syntax
- Nest windows inside windows
- Forget to bind your app struct

## Examples

See [examples/](../examples/) directory for complete working examples:
- `calculator/` - Complex responsive calculator UI
- `counter/` - Simple counter app
