package evaluator

import (
	"testing"

	"../ast"
	"../token"
)

func TestEvaluatingNumberLiteralExpression(t *testing.T) {
	expression := &ast.NumberLiteral{Value: 31}

	evaluator := &Evaluator{}
	result := evaluator.Visit(expression)

	if result.NativeValue() != 31.00 {
		t.Errorf("Expected 31, got %f", result.NativeValue())
	}
}

func TestEvaluatingComparisonExpression(t *testing.T) {
	expression := comparisonExpression(31, token.LESS_EQUAL, 32)

	evaluator := &Evaluator{}
	result := evaluator.Visit(expression)

	if result.NativeValue() != true {
		t.Errorf("Expected true, got %v", result.NativeValue())
	}
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
