package program

import "testing"

func TestRunnerCanRunBasicPrograms(t *testing.T) {
	source := `
Metadata:
version = "1.0"

Rules:
When
  foo is "bar"
    result = "success"
`

	fact := `{"foo": "bar"}`

	result := New(source).Run(fact)

	expected := "{}"

	if result != expected {
		t.Errorf("Runner expected to output %v, but got %v", expected, result)
	}
}
