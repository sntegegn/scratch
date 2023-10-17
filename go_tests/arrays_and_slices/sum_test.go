package arrays_and_slices

import "testing"

func TestSum(t *testing.T) {
	t.Run("sum a collection of any size", func(t *testing.T) {
		sl := []int{2, 2, 2, 2}
		actual := SumArray(sl)
		expected := 8

		assertCorrectValue(t, actual, expected)
	})

}

func assertCorrectValue(t *testing.T, actual, expected int) {
	t.Helper()
	if actual != expected {
		t.Errorf("expected '%d' but got '%d'", expected, actual)
	}
}
