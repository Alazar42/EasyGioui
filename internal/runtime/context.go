package runtime

import (
	"context"
	"fmt"

	"easygioui/internal/events"
	"easygioui/internal/reactivity"
	"easygioui/internal/state"
)

type Context struct {
	Tree       *Tree
	State      *state.Store
	Reactivity *reactivity.Graph
	Events     *events.Dispatcher
}

func NewContext(tree *Tree, st *state.Store, rx *reactivity.Graph, ev *events.Dispatcher) *Context {
	return &Context{Tree: tree, State: st, Reactivity: rx, Events: ev}
}

func (c *Context) Set(path string, value any) error {
	if err := c.Tree.Set(path, value); err != nil {
		c.State.SetGlobal(path, value)
		c.Reactivity.Notify(path)
		return nil
	}
	c.Reactivity.Notify(path)
	return nil
}

func (c *Context) Get(path string) (string, bool) {
	if v, ok := c.Tree.Get(path); ok {
		return v, true
	}
	if v, ok := c.State.GetGlobal(path); ok {
		return fmt.Sprint(v), true
	}
	return "", false
}

func (c *Context) Emit(eventRef string, payload any) error {
	return c.Events.Emit(context.Background(), eventRef, payload)
}
