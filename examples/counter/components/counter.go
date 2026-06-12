package components

import (
	"easygioui"
	"fmt"
)

// Counter holds the counter state.
type Counter struct {
	count int
}

// Increment increments the counter.
func (a *Counter) Increment() {
	a.count++
	easygioui.SetText("counterText", fmt.Sprintf("Counter: %d", a.count))
}

// Decrement decrements the counter.
func (a *Counter) Decrement() {
	a.count--
	easygioui.SetText("counterText", fmt.Sprintf("Counter: %d", a.count))
}

// OnButtonClick is called when the nested button is clicked.
func (a *Counter) OnButtonClick() {
	easygioui.SetText("counterText", "Button clicked!")
}
