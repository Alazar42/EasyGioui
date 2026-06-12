// Package easygioui provides a minimal declarative UI overlay for Gio applications.
// It injects directly into Gio's event loop without replacing the runtime.
package easygioui

import (
	"fmt"
	"os"

	"easygioui/internal/ast"
	"easygioui/internal/parser"
	"easygioui/internal/renderer"

	"gioui.org/layout"
	"gioui.org/op"
)

// UI holds the parsed, cached AST and rendering state.
type UI struct {
	parsed    *ast.File
	renderer  *renderer.Renderer
	bindings  map[string]interface{}
	stateText map[string]string
}

// Load parses a .easy file once and caches the AST.
// Returns nil if parsing fails (error is logged).
func Load(path string) *UI {
	data, err := os.ReadFile(path)
	if err != nil {
		fmt.Printf("easygioui: failed to read %q: %v\n", path, err)
		return nil
	}

	file, err := parser.Parse(string(data))
	if err != nil {
		fmt.Printf("easygioui: failed to parse %q: %v\n", path, err)
		return nil
	}

	ui := &UI{
		parsed:    file,
		renderer:  renderer.New(),
		bindings:  make(map[string]interface{}),
		stateText: make(map[string]string),
	}

	// Store in global for SetText/Bind convenience functions
	globalUI = ui
	return ui
}

// Register renders the UI tree inside a FrameEvent.
// It must be called once per frame, inside app.FrameEvent handling.
func Register(ops *op.Ops, gtx layout.Context, ui *UI) {
	if ui == nil || ui.parsed == nil {
		return
	}

	// Render the cached AST into Gio operations
	ui.renderer.Render(ops, gtx, ui.parsed, ui.stateText, ui.bindings)
}

// SetText updates the text value for a node by ID.
// This persists across frames and overrides the initial value.
func SetText(nodeID string, value interface{}) {
	if globalUI == nil {
		return
	}
	globalUI.stateText[nodeID] = fmt.Sprint(value)
}

// Bind registers application methods for event handlers.
// Only zero-argument methods are supported.
func Bind(app interface{}) {
	if globalUI == nil {
		return
	}
	globalUI.bindings["App"] = app
}

// Global UI instance for SetText and Bind convenience functions.
var globalUI *UI
