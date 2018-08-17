package ast

type node interface {
	TokenLiteral() string
	String() string
	Line() int
	IsExpression() bool
	IsStatement() bool
}

type Statement interface {
	node
	statementNode()
}

type Expression interface {
	node
	expressionNode()
}

type Program struct {
	Metadata KeyValueBlock
}