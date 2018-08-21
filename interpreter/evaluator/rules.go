package evaluator

import "../ast"
import "../local"

// RulesBlock the Evaluator object for the rules
type RulesBlock struct {
	SymbolTable *local.SymbolTable
	Rules       []*ast.Rule
}

// EvaluationResult is the resulting object
type EvaluationResult struct {
	// Keep this while developing
	// TODO: Make a more informative output
	Rule *ast.Rule
}

// NewRulesBlock initializes a Rules evalutator
func NewRulesBlock(st *local.SymbolTable, rules []*ast.Rule) *RulesBlock {
	return &RulesBlock{SymbolTable: st, Rules: rules}
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

	// Store evaluation result in the local table
	// TODO: Make this a compound identifier
	if rule.Identifier != nil {
		rb.storeRuleResult(rule.Identifier.String(), result.(*ast.BooleanLiteral))
	}

	return &EvaluationResult{Rule: rule}
}

func (rb *RulesBlock) storeRuleResult(name string, value *ast.BooleanLiteral) {
	rb.SymbolTable.Define(name, value)
}
