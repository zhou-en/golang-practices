package main

import "testing"

func TestLongestCommonPrefix(t *testing.T) {
	tests := []struct {
		name string
		strs []string
		want string
	}{
		{name: "common prefix exists", strs: []string{"flower", "flow", "flight"}, want: "fl"},
		{name: "no common prefix", strs: []string{"dog", "racecar", "car"}, want: ""},
		{name: "all strings are the same", strs: []string{"test", "test", "test"}, want: "test"},
		{name: "single string in array", strs: []string{"single"}, want: "single"},
		{name: "empty array of strings", strs: []string{}, want: ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestCommonPrefix(tt.strs); got != tt.want {
				t.Errorf("longestCommonPrefix(%v) = %q, want %q", tt.strs, got, tt.want)
			}
		})
	}
}
