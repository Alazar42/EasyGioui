package scripts

import (
	"fmt"
	"runtime"
)

type Mode string

const (
	ModeCompiled Mode = "compiled"
	ModePlugin   Mode = "plugin"
)

type Registry struct {
	Scripts map[string]any
}

func NewRegistry() *Registry {
	return &Registry{Scripts: map[string]any{}}
}

func (r *Registry) Register(name string, script any) {
	r.Scripts[name] = script
}

func Load(mode Mode, appDir string) (*Registry, error) {
	r := NewRegistry()
	switch mode {
	case ModeCompiled:
		return r, nil
	case ModePlugin:
		if runtime.GOOS == "windows" {
			return nil, fmt.Errorf("plugin mode is not supported on windows")
		}
		return r, nil
	default:
		return nil, fmt.Errorf("unknown scripts mode: %s", mode)
	}
}
