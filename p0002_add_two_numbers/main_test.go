package main

import (
	"slices"
	"testing"
)

func TestAddTwoNumbers(t *testing.T) {
	tests := []struct {
		name string
		l1   []int
		l2   []int
		want []int
	}{
		{name: "342 + 465 = 807", l1: []int{2, 4, 3}, l2: []int{5, 6, 4}, want: []int{7, 0, 8}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := llistToSlic(addTwoNumbers(slictToLList(tt.l1), slictToLList(tt.l2)))
			if !slices.Equal(got, tt.want) {
				t.Errorf("addTwoNumbers(%v, %v) = %v, want %v", tt.l1, tt.l2, got, tt.want)
			}
		})
	}
}
