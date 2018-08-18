package parser

import (
	"fmt"
	"testing"

	"../lexer"
)

func TestParsingMetadataBlocks(t *testing.T) {
	source := `
Metadata:
version = "1.0"

Rules:
When
  last year income <= 40000
    classification = "last"
`

	lexer := lexer.New(source)
	parser := New(lexer.Scan())

	result := parser.Parse().Metadata

	if result.Name != "Metadata" {
		t.Error("Metadata block does not a have correct name")
	}

	firstPair := result.Items[0]
	key := firstPair.Key
	value := firstPair.Value

	if key.String() != "version" {
		t.Errorf("Expected 'version' got %v", key)
	}

	if fmt.Sprintf("%v", value) != "1.0" {
		t.Errorf("Expected '1.0' got %v", value)
	}
}
