package compiler

type Token struct {
	tokenType TokenType
	lexeme    string
	literal   string
	line      int
	column    int
}
