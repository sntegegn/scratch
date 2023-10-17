package iterations

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	actual := Repeat("a", 7)
	expected := "aaaaaaa"

	if actual != expected {
		t.Errorf("expected %q got %q", expected, actual)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 7)
	}
}

func ExampleRepeat() {
	repeatedString := Repeat("ab", 3)
	fmt.Println(repeatedString)
	// Output: ababab
}
