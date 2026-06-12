package ast

type File struct {
	Nodes []*Node
}

type Node struct {
	Type       string
	ID         string
	Properties map[string]*Value
	Children   []*Node
}

type ValueKind int

const (
	ValueString ValueKind = iota
	ValueNumber
	ValueIdent
)

type Value struct {
	Kind ValueKind
	Raw  string
}
