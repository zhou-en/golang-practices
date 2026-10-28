package main

import "fmt"

func twoSum(nums []int, target int) []int {
	seen := make(map[int]int)
	for i, num := range nums {
		if idx, found := seen[target-num]; found {
			return []int{idx, i}
		}
		seen[num] = i
	}
	return nil
}

func main() {
	tests := []struct {
		nums   []int
		target int
		want   []int
	}{
		{nums: []int{2, 7, 11, 15}, target: 9, want: []int{0, 1}},
		{nums: []int{3, 2, 4}, target: 6, want: []int{1, 2}},
		{nums: []int{3, 3}, target: 6, want: []int{0, 1}},
	}

	for _, tt := range tests {
		got := twoSum(tt.nums, tt.target)
		fmt.Printf("\ntwoSum(%v, %d) = %v; want %v", tt.nums, tt.target, got, tt.want)
	}
}
