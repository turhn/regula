package evaluator

import (
	"fmt"
	"reflect"

	"github.com/turhn/regula/interpreter/ast"
	"github.com/turhn/regula/interpreter/local"
	"github.com/turhn/regula/interpreter/token"
)

// Evaluator evaluates the expressions
type Evaluator struct {
	ast.Visitor
	symbolTable *local.SymbolTable
}

// NewEvaluator initializes a new expression evaluator
func NewEvaluator(symbolTable *local.SymbolTable) *Evaluator {
	return &Evaluator{
		symbolTable: symbolTable,
	}
}

// Visit sends the visitor to the nodes
func (e *Evaluator) Visit(expression ast.Expression) ast.Expression {
	switch expression.(type) {
	case *ast.ComparisonExpression:
		return e.comparisonExpression(expression.(*ast.ComparisonExpression))
	case *ast.Identifier:
		return e.identifierExpression(expression.(*ast.Identifier))
	case *ast.CompoundIdentifier:
		return e.identifierExpression(expression.(*ast.CompoundIdentifier))
	default:
		return expression
	}
}

// Can handle both identifier expression and compound identifier expressions
func (e *Evaluator) identifierExpression(identifier ast.Expression) ast.Expression {
	symbol := e.symbolTable.Resolve(identifier.String())

	switch v := symbol.(type) {
	case *ast.NumberLiteral:
		return symbol.(*ast.NumberLiteral)
	case *ast.StringLiteral:
		return symbol.(*ast.StringLiteral)
	case *ast.BooleanLiteral:
		return symbol.(*ast.BooleanLiteral)
	default:
		panic(fmt.Sprintf("Unknown type: %v", v))
	}
}

func (e *Evaluator) comparisonExpression(comparison *ast.ComparisonExpression) ast.Expression {
	left := comparison.Left.Accept(e)
	right := comparison.Right.Accept(e)
	operator := comparison.Operator.Token.Type

	// TODO: Refactor these ugly code repetitions
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
	case token.IS:
		switch left.(type) {
		case *ast.NumberLiteral:
			result := left.NativeValue().(float64) == right.NativeValue().(float64)
			return &ast.BooleanLiteral{Value: result}
		case *ast.StringLiteral:
			result := left.NativeValue().(string) == right.NativeValue().(string)
			return &ast.BooleanLiteral{Value: result}
		default:
			fmt.Printf("I don't know how to compare a %s with %s", reflect.TypeOf(left), reflect.TypeOf(right))
			panic("Panik atak")
		}
	default:
		fmt.Printf("Unknown Comparison Operator %v\n", comparison)
		return nil
	}
}
