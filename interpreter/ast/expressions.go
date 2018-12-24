package ast

import (
	"fmt"

	"github.com/turhn/regula/interpreter/token"
)

// BinaryExpression ...
type BinaryExpression struct {
	*BaseNode
	visitable
	Left     Expression
	Operator *Operator
	Right    Expression
}

// NativeValue ...
func (be *BinaryExpression) NativeValue() interface{} {
	return nil
}

func (be *BinaryExpression) String() string {
	return fmt.Sprintf("%v %v %v", be.Left, be.Operator, be.Right)
}

// LogicalExpression ...
type LogicalExpression struct {
	*BinaryExpression
}

func NewLogicalExpression(token token.Token, leftHand Expression, operator *Operator, rightHand Expression) *LogicalExpression {
	return &LogicalExpression{
		BinaryExpression: &BinaryExpression{
			Left: leftHand, Operator: operator, Right: rightHand,
		},
	}
}

// ComparisonExpression ...
type ComparisonExpression struct {
	*BinaryExpression
}

func NewComparisonExpression(token token.Token, leftHand Expression, operator *Operator, rightHand Expression) *ComparisonExpression {
	return &ComparisonExpression{
		BinaryExpression: &BinaryExpression{
			BaseNode: &BaseNode{
				Token:       token,
				isStatement: false,
			},
			Left:     leftHand,
			Operator: operator,
			Right:    rightHand,
		},
	}
}
