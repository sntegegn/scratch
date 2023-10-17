package arrays_and_slices

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	sl := []int{2, 2, 2, 2}
	actual := SumSlice(sl)
	expected := 8

	if actual != expected {
		t.Errorf("expected '%d' but got '%d'", expected, actual)
	}

}

func TestSumAll(t *testing.T) {
	actual := SumAll([]int{2, 2, 2}, []int{1, 1, 1})
	expected := []int{6, 3}

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("expected '%v' but got '%v'", expected, actual)
	}
}

func TestSumAllTails(t *testing.T) {

	checkSums := func(t *testing.T, actual, expected []int) {
		t.Helper()
		if !reflect.DeepEqual(actual, expected) {
			t.Errorf("expected '%v' but got '%v'", expected, actual)
		}
	}
	t.Run("sum all tails", func(t *testing.T) {
		slice1 := []int{1, 2, 3}
		slice2 := []int{3, 4, 5, 6}
		actual := SumAllTails(slice1, slice2)
		expected := []int{5, 15}

		checkSums(t, actual, expected)
	})

	t.Run("safely sum empty slices", func(t *testing.T) {
		actual := SumAllTails([]int{1, 2, 3}, []int{})
		expected := []int{5, 0}

		checkSums(t, actual, expected)
	})

}
