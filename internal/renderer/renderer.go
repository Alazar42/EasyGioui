package renderer

import "easygioui/internal/runtime"

// Renderer is the Gio translation boundary.
// This baseline keeps retained-tree tracking stable while Gio API glue evolves.
// Renderer is the Gio translation boundary.
// This baseline keeps retained-tree tracking stable while Gio API glue evolves.
type Renderer struct {
	lastVersion uint64
}

func New() *Renderer {
	return &Renderer{}
}

func (r *Renderer) RunWindow(window *appWindow, tree *runtime.Tree, onEvent func(eventRef string)) error {
	_ = window
	_ = onEvent
	if tree != nil {
		r.lastVersion = tree.Version()
	}
	return nil
}
