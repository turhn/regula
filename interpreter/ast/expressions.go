package ast

import (
	"fmt"

	"../token"
)

type notPrimitive struct{}

func (np *notPrimitive) IsPrimitive() bool {
	return false
}

type BinaryExpression struct {
	*BaseNode
	visitable
	notPrimitive
	Left     Expression
	Operator *Operator
	Right    Expression
}

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
	notPrimitive
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
