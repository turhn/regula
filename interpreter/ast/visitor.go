package ast

// Visitor interface
type Visitor interface {
	Visit(Expression) Expression
}
