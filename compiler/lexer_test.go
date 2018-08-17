package compiler

import "testing"

func TestLexingMetadataBlocks(t *testing.T) {
	source := `
Metadata:
version = "1.0"
name = "Stock alert"
`
	lexer := NewLexer(source)

	var result = make([]TokenType, 9)

	tokens := lexer.Scan()

	for _, token := range tokens {
		result = append(result, token.tokenType)
	}

	expected := []TokenType{
		METADATA,
		COLUMN,
		IDENTIFIER,
		EQUAL,
		STRING,
		IDENTIFIER,
		EQUAL,
		STRING,
		EOF,
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
	lexer := NewLexer(source)
	var result = make([]TokenType, 26)

	tokens := lexer.Scan()

	for _, token := range tokens {
		result = append(result, token.tokenType)
	}

	expected := []TokenType{
		DEFINITIONS,
		COLUMN,

		IDENTIFIER,
		IDENTIFIER,
		IDENTIFIER,
		EQUAL,
		LEFT_PARENTHESIS,
		STRING,
		COMMA,
		STRING,
		COMMA,
		STRING,
		RIGHT_PARENTHESIS,

		IDENTIFIER,
		IDENTIFIER,
		EQUAL,
		LEFT_PARENTHESIS,
		STRING,
		COMMA,
		STRING,
		RIGHT_PARENTHESIS,

		IDENTIFIER,
		IDENTIFIER,
		EQUAL,
		STRING,

		EOF,
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
	lexer := NewLexer(source)
	var result = make([]TokenType, 0)

	tokens := lexer.Scan()

	for _, token := range tokens {
		result = append(result, token.tokenType)
	}

	expected := []TokenType{
		RULES,
		COLUMN,

		WHEN,
		INDENT,

		IDENTIFIER,
		IDENTIFIER,
		IDENTIFIER,
		LESS_EQUAL,
		NUMBER,

		INDENT,
		INDENT,
		IDENTIFIER,
		EQUAL,
		STRING,

		INDENT,
		IDENTIFIER,
		IDENTIFIER,
		EQUAL,
		STRING,

		INDENT,
		INDENT,
		IDENTIFIER,
		EQUAL,
		STRING,

		EOF,
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
		expected []TokenType
	}{
		{source: "product is cheap", expected: []TokenType{IDENTIFIER, IS, IDENTIFIER, EOF}},
		{source: "product is not cheap", expected: []TokenType{IDENTIFIER, IS_NOT, IDENTIFIER, EOF}},
		{source: "product is in cheap products", expected: []TokenType{IDENTIFIER, IS_IN, IDENTIFIER, IDENTIFIER, EOF}},
		{source: "product is not in cheap products", expected: []TokenType{IDENTIFIER, IS_NOT_IN, IDENTIFIER, IDENTIFIER, EOF}},
	}

	for _, test := range table {
		result := make([]TokenType, 0)

		lexer := NewLexer(test.source)
		tokens := lexer.Scan()

		for _, token := range tokens {
			result = append(result, token.tokenType)
		}

		for i, val := range test.expected {
			if val != result[i] {
				t.Errorf("Expected %v got %v", test.expected, result)
				break
			}
		}
	}
}
