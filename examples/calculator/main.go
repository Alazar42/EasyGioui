package main

import (
	"fmt"
	"os"

	"easygioui"
	"easygioui/examples/calculator/components"

	"gioui.org/app"
	"gioui.org/layout"
	"gioui.org/op"
	"gioui.org/unit"
)

var ui *easygioui.UI

func main() {
	// Load the calculator UI once
	ui = easygioui.Load("examples/calculator/ui/calc.easy")
	if ui == nil {
		fmt.Println("Failed to load calculator UI")
		os.Exit(1)
	}

	// Create and bind calculator instance
	calc := components.NewCalculator()
	easygioui.Bind(calc)

	go func() {
		w := new(app.Window)
		w.Option(app.Title("Calculator"), app.Size(unit.Dp(350), unit.Dp(500)))

		var ops op.Ops

		for {
			evt := w.Event()

			switch e := evt.(type) {
			case app.FrameEvent:
				ops.Reset()

				gtx := layout.Context{
					Ops:    &ops,
					Now:    e.Now,
					Metric: e.Metric,
					Source: e.Source,
					Values: make(map[string]interface{}),
				}
				gtx.Constraints = layout.Exact(e.Size)

				easygioui.Register(&ops, gtx, ui)
				e.Frame(&ops)

			case app.DestroyEvent:
				os.Exit(0)
			}
		}
	}()
	app.Main()
}
