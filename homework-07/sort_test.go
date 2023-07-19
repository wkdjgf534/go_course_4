package sort

import (
	"math/rand"
	"reflect"
	"sort"
	"testing"
)

func Test_sortInts(t *testing.T) {
	nums := []int{3, 7, 1, 4, 2, 6, 5}
	sort.Ints(nums)
	want := []int{1, 2, 3, 4, 5, 6, 7}
	if !reflect.DeepEqual(nums, want) {
		t.Errorf("got %d, want %d", nums, want)
	}
	t.Log("When passed an unsorted slice of integers")
}

func Test_sortStrings(t *testing.T) {
	tests := []struct {
		name string
		got  []string
		want []string
	}{
		{
			name: "when passed an unsortered slice of strings #1",
			got:  []string{"Go", "Bravo", "Gopher", "Alpha", "Grin", "Delta"},
			want: []string{"Alpha", "Bravo", "Delta", "Go", "Gopher", "Grin"},
		},
		{
			name: "when passed an unsortered slice of strings #2",
			got:  []string{"5", "3", "4", "1", "2", "6"},
			want: []string{"1", "2", "3", "4", "5", "6"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sort.Strings(tt.got)
			if !reflect.DeepEqual(tt.got, tt.want) {
				t.Errorf("sortStrings() = %v, want %v", tt.got, tt.want)
			}
		})
	}
}

func Benchmark_sortInts(b *testing.B) {
	nums := rand.Perm(20)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sort.Ints(nums)
	}
}

func Benchmark_sortFloat64s(b *testing.B) {
	min, max := 0.10, 100.10
	cap := rand.Intn(20)
	nums := make([]float64, cap)
	for i := range nums {
		nums[i] = min + rand.Float64()*(max-min)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sort.Float64s(nums)
	}
}
