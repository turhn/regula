package evaluator

import (
	"testing"

	"github.com/turhn/regula/interpreter/ast"
	"github.com/turhn/regula/interpreter/local"
)

func TestFactEvaluatorRegistersFacts(t *testing.T) {
	table := []struct {
		fact     string
		key      string
		expected interface{}
	}{
		{fact: `{"foo": "bar"}`, key: "foo", expected: "bar"},
	}

	for _, test := range table {
		locals := &local.SymbolTable{}
		RegisterFact(test.fact, locals)

		if _, ok := locals.Symbols[test.key]; !ok {
			t.Errorf("The key '%s' is not registered", test.key)
		}

		value := locals.Resolve(test.key).(ast.Expression).NativeValue()

		if value != test.expected {
			t.Errorf("Expected %v, but got %v", test.expected, value)
		}
	}
}
