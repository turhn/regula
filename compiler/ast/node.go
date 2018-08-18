package ast

type node interface {
	TokenLiteral() string
	String() string
	Line() int
	IsExpression() bool
	IsStatement() bool
}

// Statement is the interface for all the statements
type Statement interface {
	node
	statementNode()
}

// Expression is the interface for all the expressions
type Expression interface {
	node
	expressionNode()
}

// Program is the primary AST node
type Program struct {
	Metadata   *KeyValueBlock
	RulesBlock *RulesBlock
}
