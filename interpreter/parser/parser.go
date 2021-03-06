package parser

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/turhn/regula/interpreter/ast"
	"github.com/turhn/regula/interpreter/lexer"
	"github.com/turhn/regula/interpreter/token"
)

type Parser struct {
	lexer          *lexer.Lexer
	tokens         []token.Token
	reservedBlocks []token.TokenType
	current        int
	level          int
}

func New(lexer *lexer.Lexer) *Parser {
	return &Parser{
		lexer:   lexer,
		tokens:  lexer.Scan(),
		current: 0,
		level:   0,
		reservedBlocks: []token.TokenType{
			token.METADATA,
			token.ALIAS,
			token.DEFINITIONS,
			token.RULES,
			token.EXPOSE,
		},
	}
}

// Parse parses to Token stream and return an instance of Program
func (parser *Parser) Parse() *ast.Program {
	var metadataBlock *ast.KeyValueBlock
	var rulesBlock *ast.RulesBlock

	for !parser.isAtEnd() {
		peekedTokenType := parser.peek().Type

		parser.consume(peekedTokenType, "A block identifier expected.")
		parser.consume(token.COLUMN, "A column expected after the block identifier")

		switch peekedTokenType {
		case token.METADATA:
			metadataBlock = parser.parseKeyValueBlock("Metadata")
		case token.RULES:
			rulesBlock = parser.parseRulesBlock()
		default:
			fmt.Printf("Unknown block detected. Block: '%s'", peekedTokenType)
			os.Exit(1)
		}
	}

	return &ast.Program{
		Metadata: metadataBlock,
		Rules:    rulesBlock,
	}
}

func (parser *Parser) parseRulesBlock() *ast.RulesBlock {
	var rules []*ast.Rule
	parser.consume(token.WHEN, "Rules block must have When in the beginning")

	for !parser.isAtEnd() {
		rules = append(rules, parser.parseRule())
	}

	return &ast.RulesBlock{Block: ast.Block{Name: "Rules"}, Rules: rules}
}

func (parser *Parser) parseRule() *ast.Rule {
	parser.level++
	parser.consumeIndents()

	expression := parser.parseComparisonExpression()
	var name *ast.Identifier

	if parser.match(token.THEN) {
		name = parser.parseIdentifier()
	}

	result := parser.parseRuleResult()
	parser.level--

	return &ast.Rule{
		RuleExpression: expression,
		Identifier:     name,
		Result:         result,
	}
}

func (parser *Parser) parseRuleResult() *ast.RuleResult {
	var result []*ast.KeyValuePair

	for !parser.isAtEnd() && parser.check(token.INDENT) {
		parser.level++
		parser.consumeIndents()

		result = append(result, parser.parseKeyValuePair())

		parser.level--
	}

	return &ast.RuleResult{Items: result}
}

func (parser *Parser) consumeIndents() {
	for i := 1; i <= parser.level; i++ {
		parser.consume(token.INDENT, fmt.Sprintf("Expected indent level %d", parser.level))
	}
}

func (parser *Parser) parseKeyValueBlock(name string) *ast.KeyValueBlock {
	var pairs []*ast.KeyValuePair

	for !parser.isReservedBlock(parser.peek().Type) && !parser.isAtEnd() {
		pair := parser.parseKeyValuePair()
		pairs = append(pairs, pair)
	}

	return &ast.KeyValueBlock{Block: ast.Block{Name: name}, Items: pairs}
}

func (parser *Parser) parseKeyValuePair() *ast.KeyValuePair {
	key := parser.parseIdentifier()
	parser.consume(token.EQUAL, "'=' expected")
	value := parser.parseExpression()

	return &ast.KeyValuePair{Key: key, Value: value}
}

func (parser *Parser) parseIdentifier() *ast.Identifier {
	parser.consume(token.IDENTIFIER, "Identifier expected")
	return &ast.Identifier{Value: parser.previous().Lexeme}
}

func (parser *Parser) parseCompoundIdentifier() *ast.CompoundIdentifier {
	var identifiers []*ast.Identifier
	for parser.check(token.IDENTIFIER) {
		identifiers = append(identifiers, parser.parseIdentifier())
	}
	return &ast.CompoundIdentifier{Value: identifiers}
}

func (parser *Parser) parseExpression() ast.Expression {
	return parser.parseOrExpression()
}

func (parser *Parser) parseOrExpression() ast.Expression {
	expression := parser.parseAndExpression()

	for parser.match(token.OR) {
		operator := ast.NewOperator(parser.previous())
		rightHand := parser.parseAndExpression()
		expression = ast.NewLogicalExpression(parser.peek(), expression, operator, rightHand)
	}

	return expression
}

func (parser *Parser) parseAndExpression() ast.Expression {
	expression := parser.parseComparisonExpression()

	for parser.match(token.AND) {
		operator := ast.NewOperator(parser.previous())
		rightHand := parser.parseComparisonExpression()
		expression = ast.NewLogicalExpression(parser.peek(), expression, operator, rightHand)
	}

	return expression
}

func (parser *Parser) parseComparisonExpression() ast.Expression {
	expression := parser.parsePrimaryExpression()

	for parser.match(token.EQUAL, token.LESS_EQUAL, token.LESS, token.GREATER, token.GREATER_EQUAL, token.IS, token.IS_NOT, token.IS_NOT_IN, token.IS_IN) {
		operator := ast.NewOperator(parser.previous())
		rightHand := parser.parsePrimaryExpression()
		expression = ast.NewComparisonExpression(parser.peek(), expression, operator, rightHand)
	}

	return expression
}

func (parser *Parser) parsePrimaryExpression() ast.Expression {
	current := parser.peek()

	switch current.Type {
	case token.NUMBER:
		return ast.NewNumberLiteral(parser.consume(current.Type, ""))
	case token.STRING:
		return ast.NewStringLiteral(parser.consume(current.Type, ""))
	case token.IDENTIFIER:
		if parser.peekNext().Type == token.IDENTIFIER {
			return parser.parseCompoundIdentifier()
		}
		return parser.parseIdentifier()
	}

	parser.parserError(current, "Unknown Expression")

	return nil
}

func (parser *Parser) match(types ...token.TokenType) bool {
	for _, t := range types {
		if parser.check(t) {
			parser.advance()
			return true
		}
	}

	return false
}

func (parser *Parser) consume(tokenType token.TokenType, message string) token.Token {
	if parser.check(tokenType) {
		return parser.advance()
	}
	current := parser.peek()
	parser.parserError(current, message)
	return token.Token{}
}

func (parser *Parser) check(tokenType token.TokenType) bool {
	if parser.isAtEnd() {
		return false
	}

	return parser.peek().Type == tokenType
}

func (parser *Parser) advance() token.Token {
	if !parser.isAtEnd() {
		parser.current++
	}

	return parser.previous()
}

func (parser *Parser) previous() token.Token {
	return parser.tokens[parser.current-1]
}

func (parser *Parser) peek() token.Token {
	return parser.tokens[parser.current]
}

func (parser *Parser) peekNext() token.Token {
	return parser.tokens[parser.current+1]
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

func (parser *Parser) parserError(currentToken token.Token, message string) {
	source := parser.lexer.Source
	line := currentToken.Line
	column := currentToken.Column

	pointer := "^-----"

	scanner := bufio.NewScanner(strings.NewReader(source))

	// Iterate through the error line
	for i := 1; i <= line; i++ {
		scanner.Scan()
	}

	code := scanner.Text()

	// Iterate trough the error column
	for i := 0; i < column; i++ {
		pointer = " " + pointer
	}

	fullMessage := fmt.Sprintf("Parser Error [%d:%d]: %s\n", line, column, message)

	fmt.Println(code)
	fmt.Println(pointer)
	fmt.Println(fullMessage)
	fmt.Printf("Current token is %v\n", currentToken)

	os.Exit(1)
}