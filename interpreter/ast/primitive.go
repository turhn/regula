package ast

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/turhn/regula/interpreter/token"
)

// StringLiteral is the AST node for ""
type StringLiteral struct {
	Expression
	*BaseNode
	Value string
}

func (sl *StringLiteral) NativeValue() interface{} {
	return sl.Value
}

func (sl *StringLiteral) Accept(visitor Visitor) Expression {
	return visitor.Visit(sl)
}

// NewStringLiteral creates StringLiteral instance from a token
func NewStringLiteral(token token.Token) *StringLiteral {
	return &StringLiteral{
		BaseNode: &BaseNode{
			Token:       token,
			isStatement: false,
		},
		Value: token.Literal,
	}
}

func (nl *NumberLiteral) NativeValue() interface{} {
	return nl.Value
}

// NumberLiteral is the AST node to hold numeric values
type NumberLiteral struct {
	*BaseNode
	Expression
	Value float64
}

func (nl *NumberLiteral) Accept(visitor Visitor) Expression {
	return visitor.Visit(nl)
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

// BooleanLiteral is to hold boolean values
type BooleanLiteral struct {
	Expression
	*BaseNode
	Value bool
}

func (bl *BooleanLiteral) NativeValue() interface{} {
	return bl.Value
}

// Identifier is identifier node
type Identifier struct {
	Expression
	*BaseNode
	Value string
}

func (i *Identifier) String() string {
	return i.Value
}

func (i *Identifier) Accept(visitor Visitor) Expression {
	return visitor.Visit(i)
}

// CompoundIdentifier is a special identifier with multiple words
type CompoundIdentifier struct {
	Expression
	*BaseNode
	Value []*Identifier
}

func (compoundIdentifier *CompoundIdentifier) String() string {
	var identifiers = ""

	for _, identifier := range compoundIdentifier.Value {
		identifiers += " " + identifier.String()
	}

	return strings.TrimSpace(identifiers)
}
