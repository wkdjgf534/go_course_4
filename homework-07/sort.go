package sort

import "sort"

func sortIntegers(nums []int) []int {
	sort.Ints(nums)
	return nums
}

func sortStrings(strs []string) []string {
	sort.Strings(strs)
	return strs
}

func sortFloats(nums []float64) []float64 {
	sort.Float64s(nums)
	return nums
}
