package reflection

import "testing"

func TestWalk(t *testing.T) {
	expected := "Chris"
	var actual []string

	x := struct {
		Name string
	}{expected}

	walk(x, func(input string) {
		actual = append(actual, input)
	})

	if len(actual) != 1 {
		t.Errorf("wrong number of function calls. Got %d expecting %d", len(actual), 1)
	}

	if actual[0] != expected {
		t.Errorf("expected %q but got %q", expected, actual[0])
	}
}
