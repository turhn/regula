package parser

import (
	"../token"
	"../ast"
)

type Parser struct {
	tokens []token.Token
	reservedBlocks []token.TokenType
	current int
	level int
}

func New(tokens []token.Token) *Parser {
	return &Parser{
		tokens: tokens,
		current: 0,
		level: 0,
		reservedBlocks: []token.TokenType{
			token.METADATA,
			token.ALIAS,
			token.DEFINITIONS,
			token.RULES,
			token.EXPOSE,
		},
	}
}

func (parser *Parser) Parse() *ast.Program {
	var metadataBlock ast.KeyValueBlock
	
	for !parser.isAtEnd() {
		peekedTokenType := parser.peek().Type

		
	}
}

func (parser *Parser) peek() token.Token {
	return parser.tokens[parser.current]
}

func (parser *Parser) isAtEnd() bool {
	return parser.peek().Type == token.EOF
}

func (parser *Parser) isReservedBlock(tokenType token.TokenType) bool {
	for _, t := range parser.reservedBlocks {
		if t == tokenType {
			return true
		}
	}

	return false
}