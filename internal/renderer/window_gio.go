package renderer

import (
	"gioui.org/app"
	"gioui.org/io/event"
)

type appWindow struct {
	w *app.Window
}

func newWindow(title string) *appWindow {
	_ = title
	return &appWindow{w: new(app.Window)}
}

func (w *appWindow) NextEvent() event.Event {
	return w.w.Event()
}
