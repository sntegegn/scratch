package integers

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	actual := Add(5, 6)
	expected := 11

	if actual != expected {
		t.Errorf("expected '%d' but got '%d'", expected, actual)
	}
}

func ExampleAdd() {
	sum := Add(2, 3)
	fmt.Println(sum)
	// Output: 5
}
