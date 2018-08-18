package ast

import (
	"fmt"
	"os"
	"strconv"

	"../token"
)

// StringLiteral is the AST node for ""
type StringLiteral struct {
	*BaseNode
	Value string
}

func (stringLiteral *StringLiteral) expressionNode() {}

// TokenLiteral returns the value as string
func (stringLiteral *StringLiteral) TokenLiteral() string {
	return stringLiteral.Token.Literal
}

// String see TokenLiteral
func (stringLiteral *StringLiteral) String() string {
	return stringLiteral.TokenLiteral()
}

// NewStringLiteral creates StringLiteral instance from a token
func NewStringLiteral(token token.Token) *StringLiteral {
	return &StringLiteral{
		BaseNode: &BaseNode{
			Token:       token,
			isStatement: false,
		},
		Value: token.Lexeme,
	}
}

// NumberLiteral is the AST node to hold numeric values
type NumberLiteral struct {
	*BaseNode
	Value float64
}

func (numberLiteral *NumberLiteral) expressionNode() {}
func (numberLiteral *NumberLiteral) TokenLiteral() string {
	return numberLiteral.Token.Literal
}
func (numberLiteral *NumberLiteral) String() string {
	return numberLiteral.TokenLiteral()
}
func NewNumberLiteral(token token.Token) *NumberLiteral {
	value, err := strconv.ParseFloat(token.Lexeme, 64)

	if err != nil {
		fmt.Println("Unable to parse the float")
		os.Exit(1)
	}

	return &NumberLiteral{
		BaseNode: &BaseNode{
			Token:       token,
			isStatement: false,
		},
		Value: value,
	}
}

// Identifier is identifier node
type Identifier struct {
	*BaseNode
	Value string
}

func (identifier *Identifier) String() string {
	return identifier.Value
}

func (identifier *Identifier) TokenLiteral() string {
	return identifier.Value
}

func (identifier *Identifier) expressionNode() {}

type CompoundIdentifier struct {
	*BaseNode
	Value []*Identifier
}

func (compoundIdentifier *CompoundIdentifier) String() string {
	var identifiers = ""

	for _, identifier := range compoundIdentifier.Value {
		identifiers += " " + identifier.String()
	}

	return identifiers
}

func (compoundIdentifier *CompoundIdentifier) TokenLiteral() string {
	return compoundIdentifier.String()
}

func (compoundIdentifier *CompoundIdentifier) expressionNode() {}

type BinaryExpression struct {
	*BaseNode
	Left     Expression
	Operator *Operator
	Right    Expression
}

func (binaryExpression *BinaryExpression) String() string {
	return binaryExpression.Token.Lexeme
}

func (binaryExpression *BinaryExpression) TokenLiteral() string {
	return ""
}

func (binaryExpression *BinaryExpression) expressionNode() {}

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
