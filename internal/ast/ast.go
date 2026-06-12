package ast

type File struct {
	Components map[string]*Node // Named components/templates
	Nodes      []*Node
}

type Node struct {
	Type       string
	ID         string
	Properties map[string]*Value
	Styles     map[string]*Value // Optional styling properties
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
