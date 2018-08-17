package ast

type block interface {
	Statement
	Name() string
}

type KeyValuePair struct {
	Statement
	Key string
	Value Expression
}

type KeyValueBlock struct {
	block
	Items []KeyValuePair
}
