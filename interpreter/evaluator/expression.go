package evaluator

import (
	"fmt"

	"../ast"
	"../token"
)

// Evaluator evaluates the expressions
type Evaluator struct {
	ast.Visitor
}

// Visit sends the visitor to the nodes
func (e *Evaluator) Visit(expression ast.Expression) ast.Expression {
	switch expression.(type) {
	case *ast.ComparisonExpression:
		return e.comparisonExpression(expression.(*ast.ComparisonExpression))
	case *ast.NumberLiteral:
		return expression
	default:
		fmt.Printf("Unknown Expression %v\n", expression)
		return nil
	}
}

func (e *Evaluator) comparisonExpression(comparison *ast.ComparisonExpression) ast.Expression {
	left := comparison.Left.Accept(e)
	right := comparison.Right.Accept(e)
	operator := comparison.Operator.Token.Type

	switch operator {
	case token.LESS_EQUAL:
		switch left.(type) {
		case *ast.NumberLiteral:
			result := left.NativeValue().(float64) <= right.NativeValue().(float64)
			return &ast.BooleanLiteral{Value: result}
		default:
			fmt.Printf("I don't know how to compare a number with %v", right)
			panic("Panik atak")
		}
	default:
		fmt.Printf("Unknown Comparison Operator %v\n", comparison)
		return nil
	}
}
