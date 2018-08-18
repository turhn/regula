package ast

import "../token"

// Operator node ...
type Operator struct {
	*BaseNode
}

func NewOperator(operatorToken token.Token) *Operator {
	return &Operator{BaseNode: &BaseNode{Token: operatorToken}}
}
