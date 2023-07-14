package arrays

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	numbers := []int{1, 2, 3}

	got := Sum(numbers)
	want := 6

	if got != want {
		t.Errorf("got '%d' want '%d' given, '%v'", got, want, numbers)
	}
}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2}, []int{0, 9})
	want := []int{3, 9}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got '%d' want '%d'", got, want)
	}
}

func TestSumAllTails(t *testing.T) {

	checkSums := func(t *testing.T, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got '%d' want '%d'", got, want)
		}
	}

	t.Run("make the sums of some slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{0, 9, 6})
		want := []int{2, 15}

		checkSums(t, got, want)
	})

	t.Run("safely sum empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{0, 9})
		want := []int{0, 9}

		checkSums(t, got, want)
	})

}

func BenchmarkSum(b *testing.B) {
	numbers := []int{1, 2, 3, 4, 5}

	for i := 0; i < b.N; i++ {
		Sum(numbers)
	}
}

func BenchmarkSumAll(b *testing.B) {
	first := []int{1, 2, 3, 4, 5}
	second := []int{1, 2, 3, 4, 5}
	third := []int{1, 2, 3, 4, 5}

	for i := 0; i < b.N; i++ {
		SumAll(first, second, third)
	}
}

func ExampleSum() {
	sum := Sum([]int{1, 2, 3})
	fmt.Println(sum)
	// Output: 6
}

func ExampleSumAll() {
	sum := SumAll([]int{1, 2, 3}, []int{2, 2, 5})
	fmt.Println(sum)
	// Output: [6 9]
}
