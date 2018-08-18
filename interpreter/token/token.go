package token

type Token struct {
	Type TokenType
	Lexeme    string
	Literal   string
	Line      int
	Column    int
}
