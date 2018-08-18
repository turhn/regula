package ast

// Block is one of the program blocks like Metadata or Rules
type Block struct {
	Statement
	Name string
}

// KeyValuePair is a key value structure
type KeyValuePair struct {
	Statement
	Key   *Identifier
	Value Expression
}

// KeyValueBlock is a program block which consist of all KeyValuePair
type KeyValueBlock struct {
	Block
	Items []*KeyValuePair
}

// RulesBlock is the program block where the rules logic is applied
type RulesBlock struct {
	Block
	Rules []*Rule
}

type Rule struct {
	Statement
	RuleExpression Expression
	Identifier     *Identifier
	Result         *RuleResult
}

type RuleResult struct {
	Items []*KeyValuePair
}
