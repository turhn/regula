package evaluator

import (
	"github.com/turhn/regula/interpreter/ast"
	"github.com/turhn/regula/interpreter/local"
)

func EvaluateMetadataToMap(block *ast.KeyValueBlock, locals *local.SymbolTable) map[string]string {
	if block.Name != "Metadata" {
		panic("Not a metadata block")
	}

	return EvaluateKeyValuesToMap(block.Items, locals)
}

func EvaluateKeyValuesToMap(items []*ast.KeyValuePair, locals *local.SymbolTable) map[string]string {
	result := map[string]string{}

	for _, val := range items {
		key := val.Key.Value

		evaluator := NewEvaluator(locals)
		value := evaluator.Visit(val.Value)

		result[key] = value.NativeValue().(string)
	}

	return result
}