package loader

import (
	"errors"
	"os"
	"path/filepath"
	"strings"

	"easygioui/internal/ast"
	"easygioui/internal/parser"
	"easygioui/internal/scripts"
)

type Result struct {
	UI      *ast.File
	Scripts *scripts.Registry
}

func LoadApp(appDir string, mode scripts.Mode) (*Result, error) {
	ui, err := loadEasyFiles(appDir)
	if err != nil {
		return nil, err
	}
	s, err := scripts.Load(mode, appDir)
	if err != nil {
		return nil, err
	}
	return &Result{UI: ui, Scripts: s}, nil
}

func loadEasyFiles(appDir string) (*ast.File, error) {
	entries, err := os.ReadDir(appDir)
	if err != nil {
		return nil, err
	}
	var combined ast.File
	for _, e := range entries {
		if e.IsDir() || !strings.HasSuffix(e.Name(), ".easy") {
			continue
		}
		b, err := os.ReadFile(filepath.Join(appDir, e.Name()))
		if err != nil {
			return nil, err
		}
		f, err := parser.Parse(string(b))
		if err != nil {
			return nil, err
		}
		combined.Nodes = append(combined.Nodes, f.Nodes...)
	}
	if len(combined.Nodes) == 0 {
		return &ast.File{}, nil
	}
	if len(combined.Nodes) > 1 {
		return nil, errors.New("only one root node is supported for now")
	}
	return &combined, nil
}
