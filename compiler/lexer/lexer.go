package lexer

import (
	"fmt"
	"os"
	"unicode"
)

import "../token"

// Lexer struct
type Lexer struct {
	source   string
	start    int
	current  int
	column   int
	line     int
	tokens   []token.Token
	keywords map[string]token.TokenType
}

// New Lexer
func New(source string) *Lexer {
	return &Lexer{
		source:  source,
		start:   0,
		current: 0,
		column:  0,
		line:    1,
		keywords: map[string]token.TokenType{
			"Metadata":    token.METADATA,
			"Alias":       token.ALIAS,
			"Rules":       token.RULES,
			"Definitions": token.DEFINITIONS,
			"When":        token.WHEN,
			"then":        token.THEN,
			"is":          token.IS,
			"and":         token.AND,
			"or":          token.OR,
		},
	}
}

// Scan tokens
func (l *Lexer) Scan() []token.Token {
	for !l.isAtEnd() {
		l.start = l.current
		l.scanToken()
	}

	endToken := token.Token{token.EOF, "", "", l.line, l.column}

	l.tokens = append(l.tokens, endToken)

	return l.tokens
}

func (l *Lexer) scanToken() {
	c := l.advance()

	switch c {
	case '(':
		l.addEmptyToken(token.LEFT_PARENTHESIS)
	case ')':
		l.addEmptyToken(token.RIGHT_PARENTHESIS)
	case ':':
		l.addEmptyToken(token.COLUMN)
	case ',':
		l.addEmptyToken(token.COMMA)
	case '=':
		l.addEmptyToken(token.EQUAL)
	case '<':
		if l.match('=') {
			l.addEmptyToken(token.LESS_EQUAL)
		} else {
			l.addEmptyToken(token.LESS)
		}
	case '>':
		if l.match('=') {
			l.addEmptyToken(token.GREATER_EQUAL)
		} else {
			l.addEmptyToken(token.GREATER)
		}
	case '"':
		l.scanString()
	case '#':
		for l.peek() != '\n' && !l.isAtEnd() {
			l.advance()
		}
	case ' ':
		if l.match(' ') {
			l.addEmptyToken(token.INDENT)
		}
	case '\r':
	case '\t':
	case '\n':
	default:
		if l.isDigit(rune(c)) {
			l.scanNumber()
		} else if l.isAlpha(c) {
			l.scanIdentifier()
		} else {
			fmt.Printf("Unknown token '%c' at %d:%d", c, l.line, l.current)
			os.Exit(1)
		}
	}
}

func (l *Lexer) scanIdentifier() {
	for l.isAlpha(l.peek()) {
		l.advance()
	}

	text := l.getCurrent()

	if val, ok := l.keywords[text]; ok {
		l.addEmptyToken(val)
		return
	}

	// Compound 'IS' operators

	if l.lastType() == token.IS {
		if text == "not" {
			l.replaceLastToken(token.IS_NOT)
			return
		} else if text == "in" {
			l.replaceLastToken(token.IS_IN)
			return
		}
	} else if l.lastType() == token.IS_NOT && text == "in" {
		l.replaceLastToken(token.IS_NOT_IN)
		return
	}

	l.addEmptyToken(token.IDENTIFIER)
}

func (l *Lexer) replaceLastToken(tokenType token.TokenType) {
	l.tokens = l.tokens[:len(l.tokens)-1]
	l.addEmptyToken(tokenType)
}

func (l *Lexer) lastType() token.TokenType {
	tokensLen := len(l.tokens)
	if tokensLen > 0 {
		return l.tokens[tokensLen-1].Type
	}

	return token.TokenType(-1)
}

func (l *Lexer) isAlpha(expected rune) bool {
	return unicode.IsLetter(expected)
}

func (l *Lexer) scanNumber() {
	for l.isDigit(l.peek()) {
		l.advance()
	}

	if l.peek() == '.' && l.isDigit(l.peekNext()) {
		l.advance()

		for l.isDigit(l.peek()) {
			l.advance()
		}
	}

	l.addToken(token.NUMBER, l.source[l.start:l.current])
}

func (l *Lexer) isDigit(char rune) bool {
	return unicode.IsDigit(char)
}

func (l *Lexer) peekNext() rune {
	currentNext := l.current + 1
	if currentNext >= len(l.source) {
		return ' '
	}
	return rune(l.source[currentNext])
}

func (l *Lexer) peek() rune {
	if l.current >= len(l.source) {
		return rune(' ')
	}
	return rune(l.source[l.current])
}

func (l *Lexer) scanString() {
	for l.peek() != '"' && !l.isAtEnd() {
		var first = l.peek()

		if first == '\n' {
			l.line++
		}

		l.advance()
	}

	if l.isAtEnd() {
		fmt.Printf("Unterminated string literal at line %d", l.line)
		os.Exit(1)
	}

	l.advance()

	literal := l.source[l.start+1 : l.current-1]
	l.addToken(token.STRING, literal)
}

func (l *Lexer) match(expected rune) bool {
	if l.isAtEnd() {
		return false
	}

	if rune(l.source[l.current]) != expected {
		return false
	}

	l.current++
	l.column++

	return true
}

func (l *Lexer) getCurrent() string {
	return l.source[l.start:l.current]
}

func (l *Lexer) addToken(tokenType token.TokenType, literal string) {
	text := l.getCurrent()
	token := token.Token{tokenType, text, literal, l.line, l.column}
	l.tokens = append(l.tokens, token)
}

func (l *Lexer) addEmptyToken(tokenType token.TokenType) {
	l.addToken(tokenType, "")
}

func (l *Lexer) advance() rune {
	l.current++
	l.column++
	return rune(l.source[l.current-1])
}

func (l *Lexer) isAtEnd() bool {
	return l.current >= len(l.source)
}
