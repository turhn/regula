package compiler

import (
	"fmt"
	"os"
	"unicode"
)

// Lexer struct
type Lexer struct {
	source   string
	start    int
	current  int
	column   int
	line     int
	tokens   []Token
	keywords map[string]TokenType
}

// NewLexer asd
func NewLexer(source string) *Lexer {
	return &Lexer{
		source:  source,
		start:   0,
		current: 0,
		column:  0,
		line:    1,
		keywords: map[string]TokenType{
			"Metadata":    METADATA,
			"Alias":       ALIAS,
			"Rules":       RULES,
			"Definitions": DEFINITIONS,
			"When":        WHEN,
			"then":        THEN,
			"is":          IS,
			"and":         AND,
			"or":          OR,
		},
	}
}

// Scan tokens
func (l *Lexer) Scan() []Token {
	for !l.isAtEnd() {
		l.start = l.current
		l.scanToken()
	}

	endToken := Token{EOF, "", "", l.line, l.column}

	l.tokens = append(l.tokens, endToken)

	return l.tokens
}

func (l *Lexer) scanToken() {
	c := l.advance()

	switch c {
	case '(':
		l.addEmptyToken(LEFT_PARENTHESIS)
	case ')':
		l.addEmptyToken(RIGHT_PARENTHESIS)
	case ':':
		l.addEmptyToken(COLUMN)
	case ',':
		l.addEmptyToken(COMMA)
	case '=':
		l.addEmptyToken(EQUAL)
	case '<':
		if l.match('=') {
			l.addEmptyToken(LESS_EQUAL)
		} else {
			l.addEmptyToken(LESS)
		}
	case '>':
		if l.match('=') {
			l.addEmptyToken(GREATER_EQUAL)
		} else {
			l.addEmptyToken(GREATER)
		}
	case '"':
		l.scanString()
	case '#':
		for l.peek() != '\n' && !l.isAtEnd() {
			l.advance()
		}
	case ' ':
		if l.match(' ') {
			l.addEmptyToken(INDENT)
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

	if l.lastTokenType() == IS {
		if text == "not" {
			l.replaceLastToken(IS_NOT)
			return
		} else if text == "in" {
			l.replaceLastToken(IS_IN)
			return
		}
	} else if l.lastTokenType() == IS_NOT && text == "in" {
		l.replaceLastToken(IS_NOT_IN)
		return
	}

	l.addEmptyToken(IDENTIFIER)
}

func (l *Lexer) replaceLastToken(tokenType TokenType) {
	l.tokens = l.tokens[:len(l.tokens)-1]
	l.addEmptyToken(tokenType)
}

func (l *Lexer) lastTokenType() TokenType {
	tokensLen := len(l.tokens)
	if tokensLen > 0 {
		return l.tokens[tokensLen-1].tokenType
	}

	return TokenType(-1)
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

	l.addToken(NUMBER, l.source[l.start:l.current])
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
	l.addToken(STRING, literal)
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

func (l *Lexer) addToken(tokenType TokenType, literal string) {
	text := l.getCurrent()
	token := Token{tokenType, text, literal, l.line, l.column}
	l.tokens = append(l.tokens, token)
}

func (l *Lexer) addEmptyToken(tokenType TokenType) {
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
