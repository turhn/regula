package ast

type node interface {
	String() string
}

type visitable interface {
	Accept(Visitor) Expression
}

// Statement is the interface for all the statements
type Statement interface {
	node
}

// Expression is the interface for all the expressions
type Expression interface {
	node
	visitable
	IsPrimitive() bool
	NativeValue() interface{}
}

// Program is the primary AST node
type Program struct {
	Metadata *KeyValueBlock
	Rules    *RulesBlock
}
