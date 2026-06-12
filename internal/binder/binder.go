package binder

import (
	"context"
	"fmt"
	"reflect"
	"strings"

	"easygioui/internal/events"
	"easygioui/internal/runtime"
)

type EasyContext interface {
	Set(path string, value any) error
	Get(path string) (string, bool)
	Emit(eventRef string, payload any) error
}

type Engine struct {
	dispatcher *events.Dispatcher
	scripts    map[string]any
}

func New(dispatcher *events.Dispatcher) *Engine {
	return &Engine{dispatcher: dispatcher, scripts: map[string]any{}}
}

func (e *Engine) RegisterScript(name string, script any) {
	e.scripts[name] = script
}

func (e *Engine) BindTree(tree *runtime.Tree, uiCtx EasyContext) error {
	root := tree.RootSnapshot()
	if root == nil {
		return nil
	}
	var bindNode func(n *runtime.Node) error
	bindNode = func(n *runtime.Node) error {
		for _, ref := range n.BoundEvent {
			h, err := e.resolve(ref, uiCtx)
			if err != nil {
				return err
			}
			e.dispatcher.Register(ref, h)
		}
		for _, c := range n.Children {
			if err := bindNode(c); err != nil {
				return err
			}
		}
		return nil
	}
	return bindNode(root)
}

func (e *Engine) resolve(ref string, uiCtx EasyContext) (events.Handler, error) {
	parts := strings.Split(ref, ".")
	if len(parts) != 2 {
		return nil, fmt.Errorf("invalid handler ref %q", ref)
	}
	script := e.scripts[parts[0]]
	if script == nil {
		return nil, fmt.Errorf("script %q not found", parts[0])
	}
	mv := reflect.ValueOf(script).MethodByName(parts[1])
	if !mv.IsValid() {
		return nil, fmt.Errorf("method %q not found on %q", parts[1], parts[0])
	}
	ctxType := reflect.TypeOf((*EasyContext)(nil)).Elem()
	if mv.Type().NumIn() != 1 || !mv.Type().In(0).Implements(ctxType) {
		return nil, fmt.Errorf("handler %q must have signature func(EasyContext)", ref)
	}
	if mv.Type().NumOut() != 1 || !mv.Type().Out(0).Implements(reflect.TypeOf((*error)(nil)).Elem()) {
		return nil, fmt.Errorf("handler %q must return error", ref)
	}
	return func(_ context.Context, _ any) error {
		out := mv.Call([]reflect.Value{reflect.ValueOf(uiCtx)})
		if len(out) == 1 && !out[0].IsNil() {
			return out[0].Interface().(error)
		}
		return nil
	}, nil
}
