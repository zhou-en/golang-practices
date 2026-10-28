package main

import "fmt"

/*
 * @lc app=leetcode id=2 lang=golang
 *
 * [2] Add Two Numbers
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func slictToLList(s []int) *ListNode {
	dummyList := &ListNode{}
	current := dummyList
	for _, v := range s {
		current.Next = &ListNode{Val: v}
		current = current.Next
	}
	return dummyList.Next
}

func llistToSlic(ll *ListNode) []int {
	s := []int{}
	for current := ll; current != nil; current = current.Next {
		s = append(s, current.Val)
	}
	return s
}

// func addLLists(a, b *ListNode) *ListNode {
// }

func printLList(ll *ListNode) {

	for current := ll; current != nil; current = current.Next {
		fmt.Printf("\n%d", current.Val)
	}

	fmt.Printf("\n")
}

func revertLList(head *ListNode) *ListNode {
	var prev *ListNode
	current := head

	for current != nil {
		// copy the current next
		next := current.Next
		// swap the current next with previous
		current.Next = prev
		// move previous first
		prev = current
		// then move current foward
		current = next
	}
	// previous is now the head
	return prev

}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	return nil

}

// @lc code=end

func main() {
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

		fmt.Printf("Before:\n")
		printLList(node_1)

		fmt.Printf("After:\n")
		node_1 = revertLList(node_1)
		printLList(node_1)

		fmt.Printf("To Slice:\n")
		fmt.Printf("%v", llistToSlic(node_1))
	}
}
