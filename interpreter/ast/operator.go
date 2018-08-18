package ast

import "../token"

// Operator node ...
type Operator struct {
	*BaseNode
}

// NewOperator make an operator from the given token
func NewOperator(operatorToken token.Token) *Operator {
	return &Operator{BaseNode: &BaseNode{Token: operatorToken}}
}
