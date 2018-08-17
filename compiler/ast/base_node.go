package ast

import (
	"../token"
)

type BaseNode struct {
	Token token.Token
	isStatement bool
}

func (b *BaseNode) Line() int {
	return b.Token.Line
}

func (b *BaseNode) IsExpression() bool {
	return !b.isStatement
}
