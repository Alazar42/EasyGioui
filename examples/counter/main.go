package main

import (
	"fmt"
	"os"

	"easygioui"
	"easygioui/examples/counter/components"

	"gioui.org/app"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/unit"
)

var ui *easygioui.UI

func main() {
	// Load and cache the UI once
	ui = easygioui.Load("examples/counter/ui/counter.easy")
	if ui == nil {
		fmt.Println("Failed to load UI")
		os.Exit(1)
	}

	// Create app instance and bind it
	appInst := &components.Counter{}
	easygioui.Bind(appInst)

	go func() {
		w := new(app.Window)
		w.Option(app.Title("Egg Timer"), app.Size(unit.Dp(400), unit.Dp(600)))

		var ops op.Ops

		for {
			evt := w.Event()

			switch e := evt.(type) {
			case app.FrameEvent:
				// Reset ops for the new frame
				ops.Reset()

				// Create a layout context for rendering
				gtx := layout.Context{
					Ops:    &ops,
					Now:    e.Now,
					Metric: e.Metric,
					Source: e.Source,
					Values: make(map[string]interface{}),
				}
				gtx.Constraints = layout.Exact(e.Size)

				// Register and render the UI
				easygioui.Register(&ops, gtx, ui)

				e.Frame(&ops)

			case app.DestroyEvent:
				os.Exit(0)
			}
		}
	}()

	app.Main()
}
