package runtime

import (
	"easygioui/internal/ast"
)

func BuildTreeFromAST(root *ast.File) *Tree {
	t := NewTree()
	if root != nil {
		t.BuildFromAST(root)
	}
	return t
}
