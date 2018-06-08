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

	for _, token := range lexer.Scan() {
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
