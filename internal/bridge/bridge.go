package bridge

import "easygioui/internal/runtime"

type Bridge interface {
	Context() *runtime.Context
}
