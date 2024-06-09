package main

import (
	"slices"
	"testing"
)

func TestSum(t *testing.T) {

	checkSums := func(t testing.TB, got, want []int) {
		t.Helper()
		if !slices.Equal(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}

	t.Run("using the slice type for variable input size", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		want := 15

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})

	t.Run("return slice with sums", func(t *testing.T) {
		firstSlice := []int{1, 2}
		secondSlice := []int{0, 9}

		got := SumSlices(firstSlice, secondSlice)
		want := []int{3, 9}

		checkSums(t, got, want)
	})

	t.Run("sum tails of slice", func(t *testing.T) {
		firstSlice := []int{1, 2}
		secondSlice := []int{0, 9}

		got := SumTails(firstSlice, secondSlice)
		want := []int{2, 9}

		checkSums(t, got, want)
	})

	t.Run("handle tail sum of empty slice", func(t *testing.T) {
		got := SumTails([]int{}, []int{3, 4, 5})
		want := []int{0, 9}

		checkSums(t, got, want)
	})

	t.Run("exclude head and tail in slice sum", func(t *testing.T) {
		got := SumSliceInner([]int{1, 2, 3, 4, 5}, []int{6, 7, 8, 9})
		want := []int{9, 15}

		checkSums(t, got, want)
	})
}
