package reactivity

import "sync"

type Graph struct {
	mu          sync.RWMutex
	dependents  map[string]map[string]struct{}
	subscribers map[string][]func(changed string)
}

func NewGraph() *Graph {
	return &Graph{
		dependents:  map[string]map[string]struct{}{},
		subscribers: map[string][]func(changed string){},
	}
}

func (g *Graph) Bind(stateKey, nodeID string) {
	g.mu.Lock()
	defer g.mu.Unlock()
	m := g.dependents[stateKey]
	if m == nil {
		m = map[string]struct{}{}
		g.dependents[stateKey] = m
	}
	m[nodeID] = struct{}{}
}

func (g *Graph) Subscribe(stateKey string, fn func(changed string)) {
	g.mu.Lock()
	defer g.mu.Unlock()
	g.subscribers[stateKey] = append(g.subscribers[stateKey], fn)
}

func (g *Graph) Notify(stateKey string) {
	g.mu.RLock()
	subs := append([]func(changed string){}, g.subscribers[stateKey]...)
	g.mu.RUnlock()
	for _, fn := range subs {
		fn(stateKey)
	}
}
