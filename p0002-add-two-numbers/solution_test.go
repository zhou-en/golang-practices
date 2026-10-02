package p0002addtwonumbers

import (
	"fmt"
	"testing"
)

func TestAddTwoNumbers(t *testing.T) {
	tests := []struct {
		title    string
		input_1  []int
		input_2  []int
		expected []int
	}{
		{
			input_1:  []int{2, 4, 3},
			input_2:  []int{5, 6, 4},
			expected: []int{7, 0, 8},
		},
	}

	for _, test := range tests {
		node_1 := slictToLList(test.input_1)
		// node_2 := slictToLList(test.input_2)
		t.Run(test.title, func(t *testing.T) {
			fmt.Printf("Before:\n")
			printLList(node_1)

			fmt.Printf("After:\n")
			node_1 = revertLList(node_1)
			printLList(node_1)

			fmt.Printf("To Slice:\n")
			fmt.Printf("%v", llistToSlic(node_1))
		})
	}
}
