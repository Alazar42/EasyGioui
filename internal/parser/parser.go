package parser

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"

	"easygioui/internal/ast"
)

type Parser struct {
	lx      *lexer
	cur     token
	baseDir string
}

func Parse(src string) (*ast.File, error) {
	p := &Parser{lx: newLexer(src)}
	if err := p.bump(); err != nil {
		return nil, err
	}
	file := &ast.File{}
	for p.cur.kind != tokEOF {
		n, err := p.parseNode()
		if err != nil {
			return nil, err
		}
		file.Nodes = append(file.Nodes, n)
	}
	return file, nil
}

// ParseWithDir parses a source string and resolves relative @include directives from baseDir.
func ParseWithDir(src string, baseDir string) (*ast.File, error) {
	p := &Parser{lx: newLexer(src), baseDir: baseDir}
	if err := p.bump(); err != nil {
		return nil, err
	}
	file := &ast.File{}
	for p.cur.kind != tokEOF {
		// Check for include directive (include "filename")
		if p.cur.lit == "include" && p.cur.kind == tokIdent {
			if err := p.bump(); err != nil { // consume "include"
				return nil, err
			}
			if p.cur.kind != tokString {
				return nil, fmt.Errorf("expected string after include at %d", p.cur.pos)
			}
			includePath := p.cur.lit
			if err := p.bump(); err != nil {
				return nil, err
			}

			// Load the included file
			fullPath := filepath.Join(baseDir, includePath)
			data, err := os.ReadFile(fullPath)
			if err != nil {
				return nil, fmt.Errorf("failed to read include %q: %w", includePath, err)
			}

			// Parse included file recursively
			incFile, err := ParseWithDir(string(data), filepath.Dir(fullPath))
			if err != nil {
				return nil, fmt.Errorf("failed to parse include %q: %w", includePath, err)
			}
			file.Nodes = append(file.Nodes, incFile.Nodes...)
			continue
		}

		n, err := p.parseNode()
		if err != nil {
			return nil, err
		}
		file.Nodes = append(file.Nodes, n)
	}
	return file, nil
}

func (p *Parser) parseNode() (*ast.Node, error) {
	if p.cur.kind != tokIdent {
		return nil, fmt.Errorf("expected node type at %d", p.cur.pos)
	}
	n := &ast.Node{Type: p.cur.lit, Properties: map[string]*ast.Value{}}
	if err := p.bump(); err != nil {
		return nil, err
	}
	if err := p.expect(tokLBrace); err != nil {
		return nil, err
	}
	for p.cur.kind != tokRBrace && p.cur.kind != tokEOF {
		if p.cur.kind != tokIdent {
			return nil, fmt.Errorf("expected property or child at %d", p.cur.pos)
		}
		name := p.cur.lit
		if err := p.bump(); err != nil {
			return nil, err
		}

		if p.cur.kind == tokLBrace {
			child := &ast.Node{Type: name, Properties: map[string]*ast.Value{}}
			if err := p.parseNodeBody(child); err != nil {
				return nil, err
			}
			n.Children = append(n.Children, child)
			continue
		}

		if err := p.expect(tokColon); err != nil {
			return nil, err
		}
		val, err := p.parseValue()
		if err != nil {
			return nil, err
		}
		n.Properties[name] = val
		if name == "id" {
			n.ID = val.Raw
		}
	}
	if err := p.expect(tokRBrace); err != nil {
		return nil, err
	}
	return n, nil
}

func (p *Parser) parseNodeBody(n *ast.Node) error {
	if err := p.expect(tokLBrace); err != nil {
		return err
	}
	for p.cur.kind != tokRBrace && p.cur.kind != tokEOF {
		if p.cur.kind != tokIdent {
			return fmt.Errorf("expected property or child at %d", p.cur.pos)
		}
		name := p.cur.lit
		if err := p.bump(); err != nil {
			return err
		}
		if p.cur.kind == tokLBrace {
			child := &ast.Node{Type: name, Properties: map[string]*ast.Value{}}
			if err := p.parseNodeBody(child); err != nil {
				return err
			}
			n.Children = append(n.Children, child)
			continue
		}
		if err := p.expect(tokColon); err != nil {
			return err
		}
		val, err := p.parseValue()
		if err != nil {
			return err
		}
		n.Properties[name] = val
		if name == "id" {
			n.ID = val.Raw
		}
	}
	return p.expect(tokRBrace)
}

func (p *Parser) parseValue() (*ast.Value, error) {
	switch p.cur.kind {
	case tokString:
		v := &ast.Value{Kind: ast.ValueString, Raw: p.cur.lit}
		return v, p.bump()
	case tokIdent:
		raw := p.cur.lit
		if err := p.bump(); err != nil {
			return nil, err
		}
		for p.cur.kind == tokDot {
			if err := p.bump(); err != nil {
				return nil, err
			}
			if p.cur.kind != tokIdent {
				return nil, fmt.Errorf("expected ident after dot at %d", p.cur.pos)
			}
			raw += "." + p.cur.lit
			if err := p.bump(); err != nil {
				return nil, err
			}
		}
		if _, err := strconv.ParseFloat(raw, 64); err == nil {
			return &ast.Value{Kind: ast.ValueNumber, Raw: raw}, nil
		}
		return &ast.Value{Kind: ast.ValueIdent, Raw: raw}, nil
	default:
		return nil, fmt.Errorf("unexpected value token at %d", p.cur.pos)
	}
}

func (p *Parser) expect(kind tokenKind) error {
	if p.cur.kind != kind {
		return fmt.Errorf("expected %v at %d", kind, p.cur.pos)
	}
	return p.bump()
}

func (p *Parser) bump() error {
	t, err := p.lx.nextToken()
	if err != nil {
		return err
	}
	p.cur = t
	return nil
}
