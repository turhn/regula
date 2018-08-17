package ast

type StringLiteral struct {
	*BaseNode
	Value string
}

func (stringLiteral *StringLiteral) expressionNode() {}

func (stringLiteral *StringLiteral) TokenLiteral() string {
	return stringLiteral.Token.Literal
}

func (stringLiteral *StringLiteral) String() string {
	return stringLiteral.TokenLiteral()
}

type NumberLiteral struct {
	*BaseNode
	Value float64
}

