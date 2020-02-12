package sum

import (
	"reflect"
	"testing"
)

func checkSums(t *testing.T, got, want []int) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestSum(t *testing.T) {
	numbers := []int{1, 2, 3}

	got := Sum(numbers)
	want := 6

	if got != want {
		t.Errorf("got '%d' want '%d', given %v", got, want, numbers)
	}
}

func TestAll(t *testing.T) {
	got := All([]int{1, 2, 3}, []int{4, 5, 6})
	want := []int{6, 15}
	checkSums(t, got, want)
}

func TestAllTails(t *testing.T) {

	t.Run("make the sums of tails of", func(t *testing.T) {
		got := AllTails([]int{1, 2}, []int{0, 9})
		want := []int{2, 9}
		checkSums(t, got, want)
	})

	t.Run("safely sum empty slices", func(t *testing.T) {
		got := AllTails([]int{}, []int{3, 4, 5})
		want := []int{0, 9}
		checkSums(t, got, want)
	})

}
