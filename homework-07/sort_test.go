package sort

import (
	"testing"
)

func Test_sortIntegers(t *testing.T) {
	nums := []int{3, 7, 1, 4, 2, 6, 5}
	got := sortIntegers(nums)
	want := []int{1, 2, 3, 4, 5, 6, 7}
	if got[0] != want[0] {
		t.Errorf("got %d, want %d", got, want)
	}
	t.Log("ABC")
}
