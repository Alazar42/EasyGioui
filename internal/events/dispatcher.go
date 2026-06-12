package events

import "context"

type Handler func(context.Context, any) error

type Dispatcher struct {
	routes map[string]Handler
}

func NewDispatcher() *Dispatcher {
	return &Dispatcher{routes: map[string]Handler{}}
}

func (d *Dispatcher) Register(eventRef string, handler Handler) {
	d.routes[eventRef] = handler
}

func (d *Dispatcher) Emit(ctx context.Context, eventRef string, payload any) error {
	h := d.routes[eventRef]
	if h == nil {
		return nil
	}
	return h(ctx, payload)
}
