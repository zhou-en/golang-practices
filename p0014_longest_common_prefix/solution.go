package p0014longestcommonprefix

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
