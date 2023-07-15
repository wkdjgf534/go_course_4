package sort

import (
	"reflect"
	"testing"
)

func Test_sortIntegers(t *testing.T) {
	nums := []int{3, 7, 1, 4, 2, 6, 5}
	got := sortIntegers(nums)
	want := []int{1, 2, 3, 4, 5, 6, 7}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %d, want %d", got, want)
	}
	t.Log("When passed an unsorted slice of integers")
}

func Test_sortStrings(t *testing.T) {
	type args struct {
		strs []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "when passed an unsortered slice of strings",
			args: args{strs: []string{"Go", "Bravo", "Gopher", "Alpha", "Grin", "Delta"}},
			want: []string{"Alpha", "Bravo", "Delta", "Go", "Gopher", "Grin"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortStrings(tt.args.strs); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortStrings() = %v, want %v", got, tt.want)
			}
		})
	}
}
