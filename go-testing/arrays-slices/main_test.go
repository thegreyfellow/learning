package main

import (
	"fmt"
	"slices"
	"testing"
)

func ExampleSumAll() {
	sumList := SumAll([]int{1, 2, 3}, []int{1, 2, 3, 4, 5})
	fmt.Println(sumList)
	// Output: [6 15]
}

func BenchmarkSumAll(b *testing.B) {
	for b.Loop() {
		SumAll([]int{1, 2, 3}, []int{1, 2, 3, 4, 5})
	}
}

func TestSum(t *testing.T) {
	t.Run("sum of collection of any size", func(t *testing.T) {
		slice := []int{1, 2, 3}
		got := Sum(slice)
		want := 6

		if got != want {
			t.Errorf("want %d but got %d", want, got)
		}
	})

	t.Run("sum of empty collection", func(t *testing.T) {
		slice := []int{}
		got := Sum(slice)
		want := 0

		if got != want {
			t.Errorf("want %d but got %d", want, got)
		}
	})
}

func TestSumAll(t *testing.T) {
	checkSums := func(t testing.TB, got, want []int) {
		t.Helper()
		if !slices.Equal(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}

	t.Run("sum of 2 collections of any size", func(t *testing.T) {
		got := SumAll([]int{1, 2, 3}, []int{1, 2, 3, 4, 5})
		want := []int{6, 15}

		checkSums(t, got, want)
	})

	t.Run("sum of 10 collections of any size", func(t *testing.T) {
		num1 := []int{1, 2, 3}
		num2 := []int{1, 2, 3}
		num3 := []int{1, 2, 3}
		num4 := []int{1, 2, 3}
		num5 := []int{1, 2, 3}
		num6 := []int{1, 2, 3}
		num7 := []int{1, 2, 3}
		num8 := []int{1, 2, 3}
		num9 := []int{1, 2, 3}
		num10 := []int{1, 2, 3}
		got := SumAll(num1, num2, num3, num4, num5, num6, num7, num8, num9, num10)
		want := []int{6, 6, 6, 6, 6, 6, 6, 6, 6, 6}

		if !slices.Equal(got, want) {
			t.Errorf("want %d but got %d", want, got)
		}
	})
}

func TestSumAllTails(t *testing.T) {
	checkSums := func(t testing.TB, got, want []int) {
		t.Helper()
		if !slices.Equal(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}

	t.Run("sum of 2 collections tails", func(t *testing.T) {
		got := SumAllTails([]int{1, 2, 3}, []int{1, 2, 3, 4, 5})
		want := []int{5, 14}

		checkSums(t, got, want)
	})

	t.Run("sum of 2 empty collections tails", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{})
		want := []int{0, 0}

		checkSums(t, got, want)
	})
}

func TestSumAllHeads(t *testing.T) {
	checkSums := func(t testing.TB, got, want []int) {
		t.Helper()
		if !slices.Equal(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}

	t.Run("sum of 2 collections heads", func(t *testing.T) {
		got := SumAllHeads([]int{1, 2, 3}, []int{1, 2, 3, 4, 5})
		want := []int{3, 10}

		checkSums(t, got, want)
	})
}
