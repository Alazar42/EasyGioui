package state

import "sync"

type Store struct {
	mu      sync.RWMutex
	global  map[string]any
	byScope map[string]map[string]any
}

func NewStore() *Store {
	return &Store{
		global:  map[string]any{},
		byScope: map[string]map[string]any{},
	}
}

func (s *Store) SetGlobal(key string, value any) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.global[key] = value
}

func (s *Store) GetGlobal(key string) (any, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	v, ok := s.global[key]
	return v, ok
}

func (s *Store) SetScoped(scope, key string, value any) {
	s.mu.Lock()
	defer s.mu.Unlock()
	m := s.byScope[scope]
	if m == nil {
		m = map[string]any{}
		s.byScope[scope] = m
	}
	m[key] = value
}

func (s *Store) GetScoped(scope, key string) (any, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	m := s.byScope[scope]
	if m == nil {
		return nil, false
	}
	v, ok := m[key]
	return v, ok
}
