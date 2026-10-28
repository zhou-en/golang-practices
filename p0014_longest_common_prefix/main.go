package main

import (
	"cmp"
	"fmt"
	"slices"
	"strings"
)

func sortSlice(s []string) []string {
	slices.SortFunc(s, func(a, b string) int {
		// 1. Primary sort: length
		if len(a) != len(b) {
			return cmp.Compare(len(a), len(b))
		}
		// 2. Secondary sort (tie-breaker): alphabetical order
		return cmp.Compare(a, b)
	})

	return s
}

func longestCommonPrefix(s []string) string {
	if len(s) == 0 {
		return ""
	}
	fmt.Printf("before: %v", s)
	sorted := sortSlice(s)
	fmt.Printf("after: %v", s)
	prefix := s[0]
	for i := 1; i < len(sorted); i++ {
		for !strings.HasPrefix(s[i], prefix) {
			prefix = prefix[:len(prefix)-1]
			if prefix == "" {
				return ""
			}
		}
	}
	return prefix
}

func main() {
	tests := []struct {
		title string
		strs  []string
		want  string
	}{
		{title: "common prefix exists", strs: []string{"flower", "flow", "flight"}, want: "fl"},
		{title: "no common prefix", strs: []string{"dog", "racecar", "car"}, want: ""},
		{title: "all strings are the same", strs: []string{"test", "test", "test"}, want: "test"},
		{title: "single string in array", strs: []string{"single"}, want: "single"},
		{title: "empty array of strings", strs: []string{}, want: ""},
	}
	for _, tt := range tests {
		got := longestCommonPrefix(tt.strs)
		fmt.Printf("\nlongestCommonPrefix() = %v, want %v", got, tt.want)
	}
}
