package evaluator

import (
	"strings"
	"testing"

	"github.com/turhn/regula/interpreter/ast"
	"github.com/turhn/regula/interpreter/local"
	"github.com/turhn/regula/interpreter/token"
)

func TestEvaluatingNumberLiteralExpression(t *testing.T) {
	expression := &ast.NumberLiteral{Value: 31}

	evaluator := &Evaluator{}
	result := evaluator.Visit(expression)

	if result.NativeValue() != 31.00 {
		t.Errorf("Expected 31, got %f", result.NativeValue())
	}
}

func TestEvaluatingNumbersComparisonExpression(t *testing.T) {
	evaluator := &Evaluator{}

	table := []struct {
		left     float64
		operator token.TokenType
		right    float64
		expected bool
	}{
		{left: 31, operator: token.LESS_EQUAL, right: 32, expected: true},
		{left: 31, operator: token.LESS_EQUAL, right: 31, expected: true},
		{left: 31, operator: token.LESS_EQUAL, right: 30, expected: false},
		{left: 31, operator: token.IS, right: 30, expected: false},
		{left: 31, operator: token.IS, right: 31, expected: true},
	}

	for _, test := range table {
		expression := comparisonExpression(test.left, test.operator, test.right)

		result := evaluator.Visit(expression)

		if result.NativeValue() != test.expected {
			t.Errorf("Expected %v, got %v", test.expected, result.NativeValue())
		}
	}
}

func TestEvaluatingIdentifierExpression(t *testing.T) {
	// Predefine some variables in the symbol table
	symbols := make(map[string]interface{})
	symbols["salary"] = &ast.NumberLiteral{Value: 40000.00}
	symbols["bonus"] = &ast.NumberLiteral{Value: 20000.00}
	symbols["name"] = &ast.StringLiteral{Value: "John Doe"}
	symbols["active"] = &ast.BooleanLiteral{Value: true}

	symbolTable := &local.SymbolTable{
		Symbols: symbols,
	}
	evaluator := NewEvaluator(symbolTable)

	table := []struct {
		name     string
		expected interface{}
	}{
		{name: "salary", expected: 40000.00},
		{name: "bonus", expected: 20000.00},
		{name: "name", expected: "John Doe"},
		{name: "active", expected: true},
	}

	for _, test := range table {
		expression := &ast.Identifier{Value: test.name}

		result := evaluator.Visit(expression)

		if result.NativeValue() != test.expected {
			t.Errorf("Expected %v, got %v", test.expected, result.NativeValue())
		}
	}
}

func TestEvaluatingCompoundIdentifierExpression(t *testing.T) {
	symbols := make(map[string]interface{})
	symbols["last year salary"] = &ast.NumberLiteral{Value: 40000.00}
	symbols["two keys"] = &ast.NumberLiteral{Value: 2000.00}

	symbolTable := &local.SymbolTable{Symbols: symbols}
	evaluator := NewEvaluator(symbolTable)

	table := []struct {
		name     string
		expected float64
	}{
		{name: "last year salary", expected: 40000.00},
		{name: "two keys", expected: 2000.00},
	}

	for _, test := range table {
		expression := compoundIdentifier(test.name)

		result := evaluator.Visit(expression)

		if result.NativeValue() != test.expected {
			t.Errorf("Expected %v, got %v", test.expected, result.NativeValue())
		}
	}
}

func compoundIdentifier(keys string) *ast.CompoundIdentifier {
	var identifiers []*ast.Identifier
	keyArray := strings.Split(keys, " ")

	for _, key := range keyArray {
		identifiers = append(identifiers, &ast.Identifier{Value: key})
	}

	return &ast.CompoundIdentifier{Value: identifiers}
}

func comparisonExpression(left float64, operator token.TokenType, right float64) *ast.ComparisonExpression {
	return &ast.ComparisonExpression{
		BinaryExpression: &ast.BinaryExpression{
			Left:     &ast.NumberLiteral{Value: left},
			Operator: &ast.Operator{BaseNode: &ast.BaseNode{Token: token.Token{Type: operator}}},
			Right:    &ast.NumberLiteral{Value: right},
		},
	}
}
