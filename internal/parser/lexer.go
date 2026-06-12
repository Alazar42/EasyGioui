package parser

import (
	"fmt"
	"strings"
	"unicode"
)

type tokenKind int

const (
	tokEOF tokenKind = iota
	tokIdent
	tokString
	tokLBrace
	tokRBrace
	tokColon
	tokDot
)

type token struct {
	kind tokenKind
	lit  string
	pos  int
}

type lexer struct {
	src []rune
	pos int
}

func newLexer(in string) *lexer {
	return &lexer{src: []rune(in)}
}

func (l *lexer) nextToken() (token, error) {
	l.skipSpace()
	if l.pos >= len(l.src) {
		return token{kind: tokEOF, pos: l.pos}, nil
	}

	ch := l.src[l.pos]
	switch ch {
	case '{':
		l.pos++
		return token{kind: tokLBrace, lit: "{", pos: l.pos - 1}, nil
	case '}':
		l.pos++
		return token{kind: tokRBrace, lit: "}", pos: l.pos - 1}, nil
	case ':':
		l.pos++
		return token{kind: tokColon, lit: ":", pos: l.pos - 1}, nil
	case '.':
		l.pos++
		return token{kind: tokDot, lit: ".", pos: l.pos - 1}, nil
	case '"':
		return l.readString()
	default:
		if isIdentStart(ch) {
			return l.readIdent(), nil
		}
		return token{}, fmt.Errorf("unexpected char %q at %d", ch, l.pos)
	}
}

func (l *lexer) skipSpace() {
	for l.pos < len(l.src) {
		ch := l.src[l.pos]
		if ch == '#' {
			for l.pos < len(l.src) && l.src[l.pos] != '\n' {
				l.pos++
			}
			continue
		}
		if ch == '/' && l.pos+1 < len(l.src) && l.src[l.pos+1] == '/' {
			for l.pos < len(l.src) && l.src[l.pos] != '\n' {
				l.pos++
			}
			continue
		}
		if !unicode.IsSpace(ch) {
			return
		}
		l.pos++
	}
}

func (l *lexer) readIdent() token {
	start := l.pos
	for l.pos < len(l.src) {
		ch := l.src[l.pos]
		if !(unicode.IsLetter(ch) || unicode.IsDigit(ch) || ch == '_' || ch == '-' || ch == '@') {
			break
		}
		l.pos++
	}
	return token{kind: tokIdent, lit: string(l.src[start:l.pos]), pos: start}
}

func (l *lexer) readString() (token, error) {
	start := l.pos
	l.pos++
	var b strings.Builder
	for l.pos < len(l.src) {
		ch := l.src[l.pos]
		if ch == '"' {
			l.pos++
			return token{kind: tokString, lit: b.String(), pos: start}, nil
		}
		if ch == '\\' && l.pos+1 < len(l.src) {
			n := l.src[l.pos+1]
			switch n {
			case 'n':
				b.WriteRune('\n')
			case 't':
				b.WriteRune('\t')
			default:
				b.WriteRune(n)
			}
			l.pos += 2
			continue
		}
		b.WriteRune(ch)
		l.pos++
	}
	return token{}, fmt.Errorf("unterminated string at %d", start)
}

func isIdentStart(ch rune) bool {
	return unicode.IsLetter(ch) || ch == '_' || ch == '@'
}
