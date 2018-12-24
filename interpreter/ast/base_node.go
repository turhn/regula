package ast

import (
	"github.com/turhn/regula/interpreter/token"
)

// BaseNode for all AST
type BaseNode struct {
	Token       token.Token
	isStatement bool
}

// Line returns the line number of the token in the source code
func (b *BaseNode) Line() int {
	return b.Token.Line
}

// IsExpression shows if the node is an expression
func (b *BaseNode) IsExpression() bool {
	return !b.isStatement
}

// IsStatement shows if the node is a statement
func (b *BaseNode) IsStatement() bool {
	return b.isStatement
}
