package program

import (
	"github.com/turhn/regula/interpreter/ast"
	"github.com/turhn/regula/interpreter/evaluator"
)

// ProgramResult result the ultimate object to be be served
type Result struct {
	// Program is the full AST for debugging and development purposes
	Program *ast.Program `json:"program"`

	// RuleTable is the table of the Evaluation Results of the rules
	RuleTable []*evaluator.EvaluationResult `json:"rule_table"`

	// Metadata is the static metadata information of the program
	Metadata map[string]string `json:"metadata"`

	// Data is the calculated result of the rules
	Data map[string]string `json:"data"`
}