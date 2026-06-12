package widgets

import (
	"fmt"

	"easygioui/internal/runtime"
)

type Widget interface {
	Type() string
	Props() map[string]string
}

type BasicWidget struct {
	typeName string
	props    map[string]string
}

func (b BasicWidget) Type() string             { return b.typeName }
func (b BasicWidget) Props() map[string]string { return b.props }

func FromNode(n *runtime.Node) (Widget, error) {
	if n == nil {
		return nil, fmt.Errorf("nil node")
	}
	return BasicWidget{typeName: n.Type, props: n.Props}, nil
}
