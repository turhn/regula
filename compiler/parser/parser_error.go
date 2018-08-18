package parser

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"../token"
)

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
