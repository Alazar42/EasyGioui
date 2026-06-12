package runtime

import (
	"fmt"
	"sync"

	"easygioui/internal/ast"
)

type Node struct {
	Type       string
	ID         string
	Props      map[string]string
	Children   []*Node
	Version    uint64
	BoundEvent map[string]string
}

type Tree struct {
	mu      sync.RWMutex
	Root    *Node
	byID    map[string]*Node
	version uint64
}

func NewTree() *Tree {
	return &Tree{byID: map[string]*Node{}}
}

func (t *Tree) BuildFromAST(f *ast.File) {
	t.mu.Lock()
	defer t.mu.Unlock()
	t.byID = map[string]*Node{}
	if len(f.Nodes) == 0 {
		t.Root = nil
		return
	}
	t.Root = t.fromAST(f.Nodes[0])
	t.version++
}

func (t *Tree) fromAST(n *ast.Node) *Node {
	rn := &Node{
		Type:       n.Type,
		ID:         n.ID,
		Props:      map[string]string{},
		BoundEvent: map[string]string{},
	}
	for k, v := range n.Properties {
		rn.Props[k] = v.Raw
		if len(k) > 2 && k[:2] == "on" {
			rn.BoundEvent[k] = v.Raw
		}
	}
	if rn.ID != "" {
		t.byID[rn.ID] = rn
	}
	for _, c := range n.Children {
		rn.Children = append(rn.Children, t.fromAST(c))
	}
	return rn
}

func (t *Tree) Set(path string, value any) error {
	t.mu.Lock()
	defer t.mu.Unlock()
	parts := splitPath(path)
	if len(parts) == 1 {
		if t.Root == nil {
			return fmt.Errorf("root is nil")
		}
		t.Root.Props[parts[0]] = fmt.Sprint(value)
		t.Root.Version++
		t.version++
		return nil
	}
	n := t.byID[parts[0]]
	if n == nil {
		return fmt.Errorf("node id %q not found", parts[0])
	}
	n.Props[parts[1]] = fmt.Sprint(value)
	n.Version++
	t.version++
	return nil
}

func (t *Tree) Get(path string) (string, bool) {
	t.mu.RLock()
	defer t.mu.RUnlock()
	parts := splitPath(path)
	if len(parts) == 1 {
		if t.Root == nil {
			return "", false
		}
		v, ok := t.Root.Props[parts[0]]
		return v, ok
	}
	n := t.byID[parts[0]]
	if n == nil {
		return "", false
	}
	v, ok := n.Props[parts[1]]
	return v, ok
}

func (t *Tree) Version() uint64 {
	t.mu.RLock()
	defer t.mu.RUnlock()
	return t.version
}

func (t *Tree) RootSnapshot() *Node {
	t.mu.RLock()
	defer t.mu.RUnlock()
	return cloneNode(t.Root)
}

func cloneNode(in *Node) *Node {
	if in == nil {
		return nil
	}
	n := &Node{
		Type:       in.Type,
		ID:         in.ID,
		Props:      map[string]string{},
		BoundEvent: map[string]string{},
		Version:    in.Version,
	}
	for k, v := range in.Props {
		n.Props[k] = v
	}
	for k, v := range in.BoundEvent {
		n.BoundEvent[k] = v
	}
	for _, c := range in.Children {
		n.Children = append(n.Children, cloneNode(c))
	}
	return n
}

func splitPath(path string) []string {
	for i := 0; i < len(path); i++ {
		if path[i] == '.' {
			return []string{path[:i], path[i+1:]}
		}
	}
	return []string{path}
}
