package evaluator

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/turhn/regula/interpreter/ast"
	"github.com/turhn/regula/interpreter/local"
)

// RegisterFact parses the given fact and registers into the symbol table
func RegisterFact(fact string, symbolTable *local.SymbolTable) {
	var unmarshalled map[string]interface{}

	err := json.Unmarshal([]byte(fact), &unmarshalled)

	if err != nil {
		fmt.Println(err)
	}

	for key := range unmarshalled {
		value := unmarshalled[key]

		switch value.(type) {
		case string:
			symbolTable.Define(key, &ast.StringLiteral{Value: value.(string)})
		default:
			fmt.Println("Unknown type")
			os.Exit(1)
		}
	}
}
