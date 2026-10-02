package p0014longestcommonprefix

import (
	// "reflect"
	"testing"
)

func TestLongestCommonPrefix(t *testing.T) {
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
		t.Run(tt.title, func(t *testing.T) {
			if got := longestCommonPrefix(tt.strs); got != tt.want {
				t.Errorf("longestCommonPrefix() = %v, want %v", got, tt.want)
			}
		})
	}
}
