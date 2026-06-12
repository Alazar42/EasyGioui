# Layout System

Master responsive layouts with EasyGioUI's flexbox-inspired layout system.

## Layout Basics

EasyGioUI uses a simple flexbox model for layouts:
- `VBox` arranges children vertically (top → bottom)
- `HBox` arranges children horizontally (left → right)
- Components automatically expand/shrink based on available space

## Spacing Between Items

Control gaps between children with the `spacing` property:

```
VBox {
    spacing: "4"    // 4 pixels between each child
    Text { text: "Item 1" }
    Text { text: "Item 2" }
    Text { text: "Item 3" }
}
```

### Spacing Values

| Value | Result |
|-------|--------|
| 0 | No gap |
| 4 | Small gap (default) |
| 8 | Medium gap |
| 12 | Large gap |
| 16+ | Extra large gap |

### Example with Different Spacing

```
VBox {
    Text { text: "No Spacing:" }
    VBox {
        spacing: "0"
        Button { text: "A" weight: "1" }
        Button { text: "B" weight: "1" }
        Button { text: "C" weight: "1" }
    }
    
    Text { text: "With Spacing:" }
    VBox {
        spacing: "8"
        Button { text: "A" weight: "1" }
        Button { text: "B" weight: "1" }
        Button { text: "C" weight: "1" }
    }
}
```

## Flex Weights

The `weight` property controls how components expand/shrink:

### Equal Distribution

Give equal weights to distribute available space equally:

```
HBox {
    weight: "1"      // Takes full width
    spacing: "4"
    
    Button {
        weight: "1"  // Gets 1/3 of width
        text: "Button 1"
    }
    
    Button {
        weight: "1"  // Gets 1/3 of width
        text: "Button 2"
    }
    
    Button {
        weight: "1"  // Gets 1/3 of width
        text: "Button 3"
    }
}
```

### Proportional Distribution

Use different weights for unequal distribution:

```
HBox {
    weight: "1"      // Takes full width
    
    VBox {
        weight: "3"  // Takes 3/5 of width
        Text { text: "Main Content" }
    }
    
    VBox {
        weight: "2"  // Takes 2/5 of width
        Text { text: "Sidebar" }
    }
}
```

### Rigid (No Weight)

Components without a weight take only their natural size:

```
HBox {
    spacing: "4"
    
    Button {
        text: "Fixed Size"  // Takes natural size
    }
    
    Button {
        text: "Expands"
        weight: "1"       // Takes remaining space
    }
}
```

## Building Responsive Grids

### 2x2 Grid

```
VBox {
    weight: "1"
    spacing: "4"
    
    HBox {
        weight: "1"
        spacing: "4"
        Button { weight: "1" text: "1" }
        Button { weight: "1" text: "2" }
    }
    
    HBox {
        weight: "1"
        spacing: "4"
        Button { weight: "1" text: "3" }
        Button { weight: "1" text: "4" }
    }
}
```

### 3x3 Grid

```
VBox {
    weight: "1"
    spacing: "4"
    
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
    
    HBox {
        weight: "1"
        spacing: "4"
        Button { weight: "1" text: "7" }
        Button { weight: "1" text: "8" }
        Button { weight: "1" text: "9" }
    }
}
```

### Calculator Layout

Typical calculator pattern: display + grid + operations:

```
VBox {
    style: { bgColor: "darkgray" }
    
    Text {
        id: display
        text: "0"
        weight: "1"
        style: {
            size: "28"
            textColor: "white"
            bgColor: "black"
        }
    }
    
    HBox {
        weight: "4"
        spacing: "4"
        
        VBox {
            weight: "3"
            spacing: "4"
            style: { bgColor: "gray" }
            // 4x4 button grid
        }
        
        VBox {
            weight: "1"
            spacing: "4"
            style: { bgColor: "orange" }
            // Operation buttons
        }
    }
}
```

## Nested Layouts

Combine VBox and HBox for complex layouts:

### Sidebar Layout

```
HBox {
    weight: "1"
    
    VBox {
        weight: "1"
        style: { bgColor: "lightgray" }
        spacing: "4"
        Text { text: "Menu" }
        Button { text: "Home" weight: "1" onClick: App.Home }
        Button { text: "About" weight: "1" onClick: App.About }
        Button { text: "Contact" weight: "1" onClick: App.Contact }
    }
    
    VBox {
        weight: "3"
        style: { bgColor: "white" }
        Text { text: "Main Content" }
    }
}
```

### Card Layout

```
VBox {
    spacing: "8"
    
    VBox {
        style: { bgColor: "white" }
        Text { text: "Card 1" }
        Button { text: "Action" weight: "1" onClick: App.Act1 }
    }
    
    VBox {
        style: { bgColor: "white" }
        Text { text: "Card 2" }
        Button { text: "Action" weight: "1" onClick: App.Act2 }
    }
    
    VBox {
        style: { bgColor: "white" }
        Text { text: "Card 3" }
        Button { text: "Action" weight: "1" onClick: App.Act3 }
    }
}
```

## Responsive Behavior

Layouts automatically adapt to different screen sizes:

```
HBox {
    weight: "1"
    spacing: "4"
    
    // On wide screens: each gets 1/3
    // On narrow screens: shrinks proportionally
    // Still maintains proper spacing
    
    Button { weight: "1" text: "A" }
    Button { weight: "1" text: "B" }
    Button { weight: "1" text: "C" }
}
```

The layout respects container constraints and adjusts automatically.

## Center Alignment

All containers use center alignment by default:

```
VBox {
    // Items center-aligned vertically
    // and distributed evenly
}

HBox {
    // Items center-aligned horizontally
    // and distributed evenly
}
```

## Common Patterns

### Button Bar

```
HBox {
    weight: "1"
    spacing: "4"
    Button { weight: "1" text: "Save" onClick: App.Save }
    Button { weight: "1" text: "Cancel" onClick: App.Cancel }
}
```

### Form Layout

```
VBox {
    spacing: "8"
    Text { text: "Name:" }
    // Text input would go here
    
    Text { text: "Email:" }
    // Email input would go here
    
    Button {
        text: "Submit"
        weight: "1"
        onClick: App.Submit
    }
}
```

### Dashboard

```
VBox {
    weight: "1"
    spacing: "8"
    
    HBox {
        weight: "1"
        spacing: "8"
        
        VBox {
            weight: "1"
            Text { text: "Widget 1" }
        }
        
        VBox {
            weight: "1"
            Text { text: "Widget 2" }
        }
    }
    
    HBox {
        weight: "1"
        spacing: "8"
        
        VBox {
            weight: "1"
            Text { text: "Widget 3" }
        }
        
        VBox {
            weight: "1"
            Text { text: "Widget 4" }
        }
    }
}
```

## Performance Tips

**Do:**
- Use weights for responsive sizing
- Nest containers logically
- Use consistent spacing
- Keep nesting depth reasonable

**Don't:**
- Use excessive nesting (> 5 levels)
- Mix rigid and flex layouts erratically
- Use inconsistent spacing
- Override layout with manual sizing

## Layout Properties Summary

### VBox / HBox Properties

| Property | Type | Default | Purpose |
|----------|------|---------|---------|
| `spacing` | number | 4 | Gap between children |
| `weight` | number | - | Flex weight in parent |

### Text / Button Properties

| Property | Type | Default | Purpose |
|----------|------|---------|---------|
| `weight` | number | - | Flex weight in parent |

## Next Steps

- Review [Styling Guide](STYLING.md) for visual design
- Check [Components Reference](COMPONENTS.md) for all properties
- Study [examples/calculator](../examples/calculator/) for real-world layout
