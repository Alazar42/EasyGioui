package core

import (
	"context"
	"errors"
	"fmt"
	"os/exec"
	"path/filepath"
	"runtime"
	"time"

	"easygioui/internal/binder"
	"easygioui/internal/events"
	"easygioui/internal/hotreload"
	"easygioui/internal/loader"
	"easygioui/internal/reactivity"
	"easygioui/internal/renderer"
	runtimepkg "easygioui/internal/runtime"
	"easygioui/internal/scripts"
	"easygioui/internal/state"
)

func Run(appDir string, devMode bool) error {
	res, err := loader.LoadApp(filepath.Join(appDir, "app"), scripts.ModeCompiled)
	if err != nil {
		return err
	}
	tree := runtimepkg.NewTree()
	st := state.NewStore()
	rx := reactivity.NewGraph()
	ev := events.NewDispatcher()
	ctx := runtimepkg.NewContext(tree, st, rx, ev)

	if len(res.UI.Nodes) > 0 {
		tree.BuildFromAST(res.UI)
	}

	bind := binder.New(ev)
	for name, script := range res.Scripts.Scripts {
		bind.RegisterScript(name, script)
	}
	_ = bind.BindTree(tree, ctx)

	if devMode {
		go func() {
			w := hotreload.New()
			_ = w.Run(context.Background(), func() error {
				_, err := loader.LoadApp(filepath.Join(appDir, "app"), scripts.ModeCompiled)
				return err
			})
		}()
	}

	r := renderer.New()
	window := renderer.NewAppWindow(titleFromTree(tree))
	if err := r.RunWindow(window, tree, func(eventRef string) {
		_ = ctx.Emit(eventRef, nil)
	}); err != nil {
		return err
	}
	return nil
}

func Build(appDir, output string) error {
	cmd := exec.Command("go", "build", "-o", output, filepath.Join(appDir, "cmd", "easygio"))
	out, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("build failed: %w: %s", err, string(out))
	}
	return nil
}

func titleFromTree(tree *runtimepkg.Tree) string {
	root := tree.RootSnapshot()
	if root == nil {
		return "EasyGioUI"
	}
	if root.Type == "Window" {
		if t := root.Props["title"]; t != "" {
			return t
		}
	}
	return "EasyGioUI"
}

func WaitForInterrupt(ctx context.Context) error {
	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-time.After(100 * 365 * 24 * time.Hour):
		return errors.New("interrupted")
	}
}

func RuntimeInfo() string {
	return runtime.GOOS + "/" + runtime.GOARCH
}
