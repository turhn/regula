package lexer

import "testing"
import "../token"

func TestLexingMetadataBlocks(t *testing.T) {
	source := `
Metadata:
version = "1.0"
name = "Stock alert"
`
	lexer := New(source)

	var result = make([]token.TokenType, 9)

	tokens := lexer.Scan()

	for _, t := range tokens {
		result = append(result, t.Type)
	}

	expected := []token.TokenType{
		token.METADATA,
		token.COLUMN,
		token.IDENTIFIER,
		token.EQUAL,
		token.STRING,
		token.IDENTIFIER,
		token.EQUAL,
		token.STRING,
		token.EOF,
	}

	for i, val := range expected {
		if val != expected[i] {
			t.Errorf("Result does not match")
		}
	}
}

func TestLexingDefinitionBlocks(t *testing.T) {
	source := `
Definitions:
multiple word definitions = ("xxx", "yyy", "zzz")
good options = ("aaa", "ccc")
other attribute = "hello"
`
	lexer := New(source)
	var result = make([]token.TokenType, 26)

	tokens := lexer.Scan()

	for _, t := range tokens {
		result = append(result, t.Type)
	}

	expected := []token.TokenType{
		token.DEFINITIONS,
		token.COLUMN,

		token.IDENTIFIER,
		token.IDENTIFIER,
		token.IDENTIFIER,
		token.EQUAL,
		token.LEFT_PARENTHESIS,
		token.STRING,
		token.COMMA,
		token.STRING,
		token.COMMA,
		token.STRING,
		token.RIGHT_PARENTHESIS,

		token.IDENTIFIER,
		token.IDENTIFIER,
		token.EQUAL,
		token.LEFT_PARENTHESIS,
		token.STRING,
		token.COMMA,
		token.STRING,
		token.RIGHT_PARENTHESIS,

		token.IDENTIFIER,
		token.IDENTIFIER,
		token.EQUAL,
		token.STRING,

		token.EOF,
	}

	for i, val := range expected {
		if val != expected[i] {
			t.Errorf("Result does not match")
		}
	}
}

func TestLexingRulesBlock(t *testing.T) {
	source := `
Rules:
When
  last year income <= 40000
    classification = "first"
  some attribute = "asdfasdf"
    classification = "second"
`
	lexer := New(source)
	var result = make([]token.TokenType, 0)

	tokens := lexer.Scan()

	for _, t := range tokens {
		result = append(result, t.Type)
	}

	expected := []token.TokenType{
		token.RULES,
		token.COLUMN,

		token.WHEN,
		token.INDENT,

		token.IDENTIFIER,
		token.IDENTIFIER,
		token.IDENTIFIER,
		token.LESS_EQUAL,
		token.NUMBER,

		token.INDENT,
		token.INDENT,
		token.IDENTIFIER,
		token.EQUAL,
		token.STRING,

		token.INDENT,
		token.IDENTIFIER,
		token.IDENTIFIER,
		token.EQUAL,
		token.STRING,

		token.INDENT,
		token.INDENT,
		token.IDENTIFIER,
		token.EQUAL,
		token.STRING,

		token.EOF,
	}

	for i, val := range expected {
		if val != result[i] {
			t.Errorf("Expected %v\n                got %v", expected, result)
			break
		}
	}
}

func TestLexingIsIsInIsNotIsNotInOperators(t *testing.T) {
	table := []struct {
		source   string
		expected []token.TokenType
	}{
		{source: "product is cheap", expected: []token.TokenType{token.IDENTIFIER, token.IS, token.IDENTIFIER, token.EOF}},
		{source: "product is not cheap", expected: []token.TokenType{token.IDENTIFIER, token.IS_NOT, token.IDENTIFIER, token.EOF}},
		{source: "product is in cheap products", expected: []token.TokenType{token.IDENTIFIER, token.IS_IN, token.IDENTIFIER, token.IDENTIFIER, token.EOF}},
		{source: "product is not in cheap products", expected: []token.TokenType{token.IDENTIFIER, token.IS_NOT_IN, token.IDENTIFIER, token.IDENTIFIER, token.EOF}},
	}

	for _, test := range table {
		result := make([]token.TokenType, 0)

		lexer := New(test.source)
		tokens := lexer.Scan()

		for _, t := range tokens {
			result = append(result, t.Type)
		}

		for i, val := range test.expected {
			if val != result[i] {
				t.Errorf("Expected %v got %v", test.expected, result)
				break
			}
		}
	}
}
