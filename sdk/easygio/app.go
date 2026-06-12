package easygio

import (
	"easygioui/internal/ast"
	runtimepkg "easygioui/internal/runtime"
)

type App struct {
	nodes []*ast.Node
}

func New() *App {
	return &App{}
}

func (a *App) VBox(children ...*ast.Node) *ast.Node {
	return &ast.Node{Type: "VBox", Children: children, Properties: map[string]*ast.Value{}}
}

func (a *App) HBox(children ...*ast.Node) *ast.Node {
	return &ast.Node{Type: "HBox", Children: children, Properties: map[string]*ast.Value{}}
}

func (a *App) Text(id, value string) *ast.Node {
	return &ast.Node{
		Type: "Text",
		ID:   id,
		Properties: map[string]*ast.Value{
			"id":    {Kind: ast.ValueIdent, Raw: id},
			"value": {Kind: ast.ValueString, Raw: value},
		},
	}
}

func (a *App) Button(id, text, onClick string) *ast.Node {
	return &ast.Node{
		Type: "Button",
		ID:   id,
		Properties: map[string]*ast.Value{
			"id":      {Kind: ast.ValueIdent, Raw: id},
			"text":    {Kind: ast.ValueString, Raw: text},
			"onClick": {Kind: ast.ValueIdent, Raw: onClick},
		},
	}
}

func (a *App) Window(title string, child *ast.Node) *App {
	n := &ast.Node{
		Type: "Window",
		Properties: map[string]*ast.Value{
			"title": {Kind: ast.ValueString, Raw: title},
		},
		Children: []*ast.Node{child},
	}
	a.nodes = []*ast.Node{n}
	return a
}

func (a *App) File() *ast.File {
	return &ast.File{Nodes: a.nodes}
}

func (a *App) BuildTree() *runtimepkg.Tree {
	return runtimepkg.BuildTreeFromAST(a.File())
}
