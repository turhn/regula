package program

import (
	"encoding/json"
	"fmt"

	"github.com/turhn/regula/interpreter/ast"
	"github.com/turhn/regula/interpreter/evaluator"
	"github.com/turhn/regula/interpreter/lexer"
	"github.com/turhn/regula/interpreter/local"
	"github.com/turhn/regula/interpreter/parser"
)

// Runner is program runner
type Runner struct {
	Program *ast.Program
}

// New Runner
func New(source string) *Runner {
	lexer := lexer.New(source)
	parser := parser.New(lexer)
	program := parser.Parse()

	return &Runner{Program: program}
}

// Run the program against a fact
func (r *Runner) Run(fact string) string {
	locals := &local.SymbolTable{}

	// Register Fact into the symbol table
	evaluator.RegisterFact(fact, locals)

	// Definitions Block
	// Add definitions to the symbol table here

	// dumpObject(r.Program, "AST output....")

	// Rule Block
	rules := r.Program.Rules.Rules
	ruleBlockEvaluator := evaluator.NewRulesBlock(locals, rules)
	result := ruleBlockEvaluator.Evaluate()

	// Convert to JSON string
	jsonBytes, err := json.Marshal(result)

	if err != nil {
		fmt.Printf("An error occured while running the application: %v\n", err)
	}

	return string(jsonBytes)
}

func dumpObject(obj interface{}, message string) {
	fmt.Println(message)
	bytes, err := json.MarshalIndent(obj, "", "  ")

	if err != nil {
		fmt.Printf("An error occured while marshaling %v to json", obj)
	}

	fmt.Println(string(bytes))
}
