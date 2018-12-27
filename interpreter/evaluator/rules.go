package evaluator

import (
	"github.com/turhn/regula/interpreter/ast"
	"github.com/turhn/regula/interpreter/local"
)

// RulesBlock the Evaluator object for the rules
type RulesBlock struct {
	SymbolTable *local.SymbolTable
	Rules       []*ast.Rule
	Data        map[string]string
}

// EvaluationResult is the resulting object
type EvaluationResult struct {
	Rule     *ast.Rule
	Matching bool
}

// NewRulesBlock initializes a Rules evaluator
func NewRulesBlock(st *local.SymbolTable, rules []*ast.Rule) *RulesBlock {
	return &RulesBlock{
		SymbolTable: st,
		Rules: rules,
		Data: make(map[string]string),
	}
}

// Evaluate starts the evaluation
func (rb *RulesBlock) Evaluate() []*EvaluationResult {
	rules := rb.Rules

	var ruleResults []*EvaluationResult

	for _, rule := range rules {
		ruleResults = append(ruleResults, rb.evaluateRule(rule))
	}

	return ruleResults
}

func (rb *RulesBlock) evaluateRule(rule *ast.Rule) *EvaluationResult {
	evaluator := NewEvaluator(rb.SymbolTable)

	result := evaluator.Visit(rule.RuleExpression)
	value := result.(*ast.BooleanLiteral)

	// Store evaluation result in the local table
	// TODO: Make this a compound identifier
	if rule.Identifier != nil {
		rb.storeRuleResult(rule.Identifier.String(), value)
	}

	ruleResultsMap := EvaluateKeyValuesToMap(rule.Result.Items, rb.SymbolTable)
	rb.appendToResultData(ruleResultsMap)

	return &EvaluationResult{Rule: rule, Matching: value.Value}
}

func (rb *RulesBlock) storeRuleResult(name string, value *ast.BooleanLiteral) {
	rb.SymbolTable.Define(name, value)
}

func (rb *RulesBlock) appendToResultData(ruleResultMap map[string]string) {
	for key, val := range ruleResultMap {
		rb.Data[key] = val
	}
}