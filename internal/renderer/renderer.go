package renderer

import (
	"fmt"
	"image"
	"image/color"
	"reflect"
	"strconv"
	"strings"

	"easygioui/internal/ast"

	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/op/clip"
	"gioui.org/op/paint"
	"gioui.org/unit"
	"gioui.org/widget"
	"gioui.org/widget/material"
)

// Renderer converts an AST into Gio operations.
type Renderer struct {
	buttons     map[string]*widget.Clickable
	theme       *material.Theme
	stateBuffer map[string]string
	bindings    map[string]interface{}
	components  map[string]*ast.Node
	eventQueue  []Event
}

// Event represents a UI event (button click, etc).
type Event struct {
	handler string
	app     interface{}
}

// New creates a new renderer.
func New() *Renderer {
	th := material.NewTheme()
	return &Renderer{
		buttons:     make(map[string]*widget.Clickable),
		theme:       th,
		stateBuffer: make(map[string]string),
		eventQueue:  make([]Event, 0),
	}
}

// Render processes the AST and renders it into Gio operations.
func (r *Renderer) Render(ops *op.Ops, gtx layout.Context, file *ast.File,
	stateText map[string]string, bindings map[string]interface{}) {

	// Store state for this frame
	r.stateBuffer = stateText
	r.bindings = bindings
	r.components = file.Components

	// Find and render the Window node
	for _, node := range file.Nodes {
		if node.Type == "Window" {
			r.renderNode(gtx, node)
			break
		}
	}

	// Process any queued events
	r.processEvents()
}

// renderNode recursively renders a node and its children.
func (r *Renderer) renderNode(gtx layout.Context, node *ast.Node) layout.Dimensions {
	// Check if this is a component reference
	if comp, ok := r.components[node.Type]; ok {
		// Render the component definition with the component's properties/children merged
		return r.renderNode(gtx, comp)
	}

	switch node.Type {
	case "Window":
		return r.renderWindow(gtx, node)
	case "VBox":
		return r.renderVBox(gtx, node)
	case "HBox":
		return r.renderHBox(gtx, node)
	case "Text":
		return r.renderText(gtx, node)
	case "Button":
		return r.renderButton(gtx, node)
	default:
		return layout.Dimensions{}
	}
}

// renderWindow renders the root window content.
func (r *Renderer) renderWindow(gtx layout.Context, node *ast.Node) layout.Dimensions {
	// Window is a logical container, render its children
	if len(node.Children) == 0 {
		return layout.Dimensions{}
	}

	// Render the first child (typically a VBox or HBox)
	return r.renderNode(gtx, node.Children[0])
}

// renderVBox renders a vertical box layout.
func (r *Renderer) renderVBox(gtx layout.Context, node *ast.Node) layout.Dimensions {
	// Create FlexChild items for each child
	children := make([]layout.FlexChild, len(node.Children))
	for i, child := range node.Children {
		child := child // capture for closure
		children[i] = layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return r.renderNode(gtx, child)
		})
	}

	flex := layout.Flex{Axis: layout.Vertical}
	dims := flex.Layout(gtx, children...)

	// Apply background color if specified
	if bgColor, ok := node.Styles["bgColor"]; ok {
		c := GetColor(bgColor.Raw)
		shp := clip.Rect{Max: dims.Size}
		paint.FillShape(gtx.Ops, c, shp.Op())
	}

	return dims
}

// renderHBox renders a horizontal box layout.
func (r *Renderer) renderHBox(gtx layout.Context, node *ast.Node) layout.Dimensions {
	// Create FlexChild items for each child
	children := make([]layout.FlexChild, len(node.Children))
	for i, child := range node.Children {
		child := child // capture for closure
		children[i] = layout.Rigid(func(gtx layout.Context) layout.Dimensions {
			return r.renderNode(gtx, child)
		})
	}

	flex := layout.Flex{Axis: layout.Horizontal}
	dims := flex.Layout(gtx, children...)

	// Apply background color if specified
	if bgColor, ok := node.Styles["bgColor"]; ok {
		c := GetColor(bgColor.Raw)
		shp := clip.Rect{Max: dims.Size}
		paint.FillShape(gtx.Ops, c, shp.Op())
	}

	return dims
}

// renderText renders a text label.
func (r *Renderer) renderText(gtx layout.Context, node *ast.Node) layout.Dimensions {
	// Get text value from state or properties
	text := r.getText(node)

	// Get text size (default 16), check properties first then styles
	size := 16
	if sizeStr, ok := node.Properties["size"]; ok {
		if num := parseNumber(sizeStr.Raw); num > 0 {
			size = num
		}
	}
	if sizeStr, ok := node.Styles["size"]; ok {
		if num := parseNumber(sizeStr.Raw); num > 0 {
			size = num
		}
	}

	label := material.Label(r.theme, unit.Sp(float32(size)), text)

	// Apply custom text color if specified, check properties first then styles
	if colorProp, ok := node.Properties["textColor"]; ok {
		label.Color = GetColor(colorProp.Raw)
	}
	if colorProp, ok := node.Styles["textColor"]; ok {
		label.Color = GetColor(colorProp.Raw)
	}

	// Apply background color if specified
	if bgColor, ok := node.Styles["bgColor"]; ok {
		c := GetColor(bgColor.Raw)
		shp := clip.Rect{Max: image.Pt(gtx.Constraints.Max.X, gtx.Constraints.Max.Y)}
		paint.FillShape(gtx.Ops, c, shp.Op())
	}

	return label.Layout(gtx)
}

// getText retrieves the text for a node, preferring state override.
func (r *Renderer) getText(node *ast.Node) string {
	// Check state override first
	if state, ok := r.stateBuffer[node.ID]; ok && state != "" {
		return state
	}

	// Check properties
	if val, ok := node.Properties["text"]; ok {
		return val.Raw
	}
	if val, ok := node.Properties["value"]; ok {
		return val.Raw
	}

	return ""
}

// renderButton renders a button widget.
func (r *Renderer) renderButton(gtx layout.Context, node *ast.Node) layout.Dimensions {
	// Ensure button exists in cache
	if _, ok := r.buttons[node.ID]; !ok {
		r.buttons[node.ID] = &widget.Clickable{}
	}

	btn := r.buttons[node.ID]
	text := r.getText(node)

	// Check for clicks and queue event
	if btn.Clicked(gtx) {
		if handler, ok := node.Properties["onClick"]; ok {
			r.eventQueue = append(r.eventQueue, Event{
				handler: handler.Raw,
				app:     r.bindings["App"],
			})
		}
	}

	// Render button with optional bgColor styling
	btStyle := material.Button(r.theme, btn, text)

	// Apply background color from styles if present
	if bgColorProp, ok := node.Styles["bgColor"]; ok {
		btStyle.Background = GetColor(bgColorProp.Raw)
	}

	return btStyle.Layout(gtx)
}

// processEvents executes any queued events.
func (r *Renderer) processEvents() {
	for _, evt := range r.eventQueue {
		r.executeHandler(evt.handler, evt.app)
	}
	r.eventQueue = r.eventQueue[:0] // Clear queue
}

// executeHandler calls a bound method.
func (r *Renderer) executeHandler(handlerPath string, app interface{}) {
	if app == nil {
		return
	}

	// Parse "App.MethodName" -> "MethodName"
	parts := strings.SplitN(handlerPath, ".", 2)
	if len(parts) != 2 {
		fmt.Printf("renderer: invalid handler path %q\n", handlerPath)
		return
	}

	methodName := parts[1]
	rv := reflect.ValueOf(app)

	// Get method
	method := rv.MethodByName(methodName)
	if !method.IsValid() {
		fmt.Printf("renderer: method %q not found\n", methodName)
		return
	}

	// Call with no arguments
	results := method.Call(nil)
	if len(results) > 0 && method.Type().NumOut() > 0 {
		if errIface := results[len(results)-1].Interface(); errIface != nil {
			if err, ok := errIface.(error); ok {
				fmt.Printf("renderer: handler error: %v\n", err)
			}
		}
	}
}

// GetColor is a utility function that wasn't used yet
func GetColor(s string) color.NRGBA {
	switch s {
	case "red":
		return color.NRGBA{R: 255, A: 255}
	case "green":
		return color.NRGBA{G: 255, A: 255}
	case "blue":
		return color.NRGBA{B: 255, A: 255}
	case "white":
		return color.NRGBA{R: 255, G: 255, B: 255, A: 255}
	case "black":
		return color.NRGBA{R: 0, G: 0, B: 0, A: 255}
	case "yellow":
		return color.NRGBA{R: 255, G: 255, B: 0, A: 255}
	case "cyan":
		return color.NRGBA{R: 0, G: 255, B: 255, A: 255}
	case "magenta":
		return color.NRGBA{R: 255, G: 0, B: 255, A: 255}
	case "orange":
		return color.NRGBA{R: 255, G: 165, B: 0, A: 255}
	case "purple":
		return color.NRGBA{R: 128, G: 0, B: 128, A: 255}
	case "gray", "grey":
		return color.NRGBA{R: 128, G: 128, B: 128, A: 255}
	case "lightgray", "lightgrey":
		return color.NRGBA{R: 200, G: 200, B: 200, A: 255}
	case "darkgray", "darkgrey":
		return color.NRGBA{R: 64, G: 64, B: 64, A: 255}
	default:
		return color.NRGBA{R: 200, G: 200, B: 200, A: 255}
	}
}

// parseNumber tries to parse a string as an integer.
func parseNumber(s string) int {
	n, _ := strconv.Atoi(strings.TrimSpace(s))
	return n
}
