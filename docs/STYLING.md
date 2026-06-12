# Styling Guide

Learn how to style your EasyGioUI components with colors, sizes, and backgrounds.

## Inline Styling

All components support optional `style {}` blocks inline:

```
Text {
    text: "Styled Text"
    style: {
        size: "24"
        textColor: "blue"
        bgColor: "lightgray"
    }
}
```

## Style Properties

### Size

Font size for `Text` and `Button` components.

```
Text {
    text: "Large Text"
    style: {
        size: "32"    // Pixels
    }
}

Button {
    text: "Big Button"
    onClick: App.Handler
    style: {
        size: "20"
    }
}
```

| Size | Usage |
|------|-------|
| 12-16 | Small labels |
| 16-20 | Body text |
| 20-28 | Section headers |
| 28+ | Main titles |

### Text Color (`textColor`)

Text color for `Text` and `Button` components.

```
Text {
    text: "Colored Text"
    style: {
        textColor: "white"
    }
}

Button {
    text: "Click Me"
    onClick: App.Handler
    style: {
        textColor: "white"
    }
}
```

### Background Color (`bgColor`)

Background color for any component.

```
VBox {
    style: {
        bgColor: "blue"
    }
    Text { text: "On blue background" }
}
```

## Available Colors

### Primary Colors

```
red      - Bright red (#FF0000)
green    - Bright green (#00FF00)
blue     - Bright blue (#0000FF)
white    - White (#FFFFFF)
black    - Black (#000000)
```

### Secondary Colors

```
yellow   - Bright yellow (#FFFF00)
cyan     - Cyan (#00FFFF)
magenta  - Magenta (#FF00FF)
orange   - Orange (#FFA500)
purple   - Purple (#800080)
```

### Neutral Colors

```
gray      - Medium gray (#808080)
lightgray - Light gray (#C8C8C8)
darkgray  - Dark gray (#404040)
```

## Style Properties Reference

### Supported Properties

| Property | Components | Type | Example | Notes |
|----------|-----------|------|---------|-------|
| `size` | Text, Button | number | `"16"` | Font size in pixels |
| `textColor` | Text, Button | color | `"white"` | Text color name |
| `bgColor` | All | color | `"blue"` | Background color name |

### Property Support by Component

**Text:**
- `size` - Font size in pixels
- `textColor` - Text color
- `bgColor` - Background color

**Button:**
- `size` - Button text size in pixels
- `textColor` - Button text color
- `bgColor` - Button background color

**VBox / HBox:**
- `bgColor` - Background color for the container

**Window:**
- `bgColor` - Window background color

## Styling Patterns

### Dark Theme

```
Window {
    VBox {
        style: {
            bgColor: "black"
        }
        
        Text {
            text: "Dark Theme"
            style: {
                size: "24"
                textColor: "white"
            }
        }
    }
}
```

### Light Theme

```
Window {
    VBox {
        style: {
            bgColor: "white"
        }
        
        Text {
            text: "Light Theme"
            style: {
                textColor: "black"
            }
        }
    }
}
```

### Colored Sections

```
HBox {
    weight: "1"
    spacing: "4"
    
    VBox {
        weight: "1"
        style: { bgColor: "red" }
        Text {
            text: "Danger"
            style: { textColor: "white" }
        }
    }
    
    VBox {
        weight: "1"
        style: { bgColor: "green" }
        Text {
            text: "Success"
            style: { textColor: "white" }
        }
    }
    
    VBox {
        weight: "1"
        style: { bgColor: "blue" }
        Text {
            text: "Info"
            style: { textColor: "white" }
        }
    }
}
```

### Highlight Buttons

```
VBox {
    spacing: "4"
    
    Button {
        text: "Primary"
        onClick: App.Action
        style: { 
            bgColor: "blue"
            textColor: "white"
            size: "16"
        }
        weight: "1"
    }
    
    Button {
        text: "Success"
        onClick: App.Action
        style: { 
            bgColor: "green"
            textColor: "white"
            size: "16"
        }
        weight: "1"
    }
    
    Button {
        text: "Danger"
        onClick: App.Action
        style: { 
            bgColor: "red"
            textColor: "white"
            size: "16"
        }
        weight: "1"
    }
}
```

### Display with Contrast Background

```
Text {
    id: "display"
    text: "0"
    style: {
        size: "28"
        textColor: "white"
        bgColor: "black"
    }
}
```

## Real-World Examples

### Calculator Display

```
VBox {
    style: { bgColor: "darkgray" }
    
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
        weight: "3"
        style: { bgColor: "gray" }
        // ... buttons ...
    }
    
    HBox {
        weight: "1"
        style: { bgColor: "orange" }
        // ... operation buttons ...
    }
}
```

### Status Panel

```
VBox {
    spacing: "8"
    style: { bgColor: "lightgray" }
    
    Text {
        text: "Status: Online"
        style: {
            textColor: "green"
            size: "16"
        }
    }
    
    Button {
        text: "Disconnect"
        onClick: App.Disconnect
        style: { bgColor: "red" }
        weight: "1"
    }
}
```

### Form Layout

```
VBox {
    spacing: "12"
    style: { bgColor: "white" }
    
    Text {
        text: "Login Form"
        style: {
            size: "24"
            textColor: "darkgray"
        }
    }
    
    Text {
        text: "Username"
        style: { textColor: "gray" }
    }
    // ... input field ...
    
    Text {
        text: "Password"
        style: { textColor: "gray" }
    }
    // ... password field ...
    
    Button {
        text: "Login"
        onClick: App.Login
        style: { bgColor: "blue" }
        weight: "1"
    }
}
```

## Style Combinations

### Header with Color

```
Text {
    text: "Section Title"
    style: {
        size: "20"
        textColor: "white"
        bgColor: "blue"
    }
}
```

### Large Clickable Button

```
Button {
    text: "START"
    onClick: App.Start
    weight: "1"
    style: {
        bgColor: "green"
    }
}
```

### Disabled-Looking Button

```
Button {
    text: "Pending..."
    onClick: App.DoNothing
    style: {
        bgColor: "lightgray"
    }
}
```

## Styling Tips

**Do:**
- Use high contrast (dark text on light, light text on dark)
- Limit color palette to 3-5 colors
- Use consistent sizing for related elements
- Apply related styling to containers, not individual items
- Use semantic colors (red = danger, green = success)

**Don't:**
- Mix too many colors
- Use very small sizes (< 12px) for text
- Use low-contrast color combinations
- Style every component differently
- Use light text on light backgrounds

## Responsive Styling

Styling works across all screen sizes. The layout automatically adapts:

```
HBox {
    weight: "1"      // Takes full width
    spacing: "4"
    
    Button {
        weight: "1"  // Takes 1/3 of width
        text: "A"
        style: { bgColor: "red" }
    }
    
    Button {
        weight: "1"  // Takes 1/3 of width
        text: "B"
        style: { bgColor: "green" }
    }
    
    Button {
        weight: "1"  // Takes 1/3 of width
        text: "C"
        style: { bgColor: "blue" }
    }
}
```

This works on all screen sizes, automatically resizing buttons.

## Performance

Styling has minimal performance impact:
- Colors are evaluated once during rendering
- No style recalculation per frame
- Cached internally

## Next Steps

- Read [Layout System](LAYOUT.md) for responsive design
- Check [Components Reference](COMPONENTS.md) for all properties
- See [examples/](../examples/) for real-world usage
