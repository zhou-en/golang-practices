package main

import "testing"

func TestLongestValid(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{name: "(()", input: "(()", want: 2},
		{name: ")()())", input: ")()())", want: 4},
		{name: "empty", input: "", want: 0},
		{name: "()(()", input: "()(()", want: 2},
		{name: "()(())", input: "()(())", want: 6},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestValid(tt.input); got != tt.want {
				t.Errorf("longestValid(%q) = %d, want %d", tt.input, got, tt.want)
			}
		})
	}
}
