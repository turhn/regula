package program

import (
	"testing"
)

func TestRunnerCanRunBasicPrograms(t *testing.T) {
	source := `
Metadata:
version = "1.0"

Rules:
When
  foo is "bar"
    result = "success"
    foo = foo
`
	fact := `{"foo": "bar"}`

	compiled := New(source).Run(fact)

	// Test Metadata block
	if compiled.Metadata["version"] != "1.0" {
		t.Errorf("Expected Metadata key \"%s\", but got \"%s\"", "1.0", compiled.Metadata["version"])
	}

	// Test it has 1 rule
	if len(compiled.RuleTable) != 1 {
		t.Errorf("Runner expected to have %d rules but got %d", 1, len(compiled.RuleTable))
	}

	// Test it generates the response as test
	expectedData := map[string]string {}
	expectedData["result"] = "success"
	expectedData["foo"] = "bar"

	for key := range expectedData {
		if compiled.Data[key] != expectedData[key] {
			t.Errorf("Expected '%s:%s' got '%s:%s'", key, expectedData[key], key, compiled.Data[key])
		}
	}
}

