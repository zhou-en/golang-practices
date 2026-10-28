// Given a string s containing just the characters '(', ')', '{', '}', '[' and ']',
// determine if the input string is valid.

// An input string is valid if:

//     Open brackets must be closed by the same type of brackets.
//     Open brackets must be closed in the correct order.
//     Every close bracket has a corresponding open bracket of the same type.

// Example 1:

// Input: s = "()"

// Output: true

// Example 2:

// Input: s = "()[]{}"

// Output: true

// Example 3:

// Input: s = "(]"

// Output: false

// Example 4:

// Input: s = "([])"

// Output: true

// Example 5:

// Input: s = "([)]"

// Output: false

// Constraints:

//     1 <= s.length <= 104
//     s consists of parentheses only '()[]{}'.

package main

import "fmt"

func isValid(s string) bool {

	bracketMap := map[string]string{
		"(": ")",
		"[": "]",
		"{": "}",
	}

	seen := make([]string, 0, len(s))

	// fmt.Printf("%v\n", string(s[0]))
	for i, v := range s {
		if i == 0 {
			seen = append(seen, string(s[0]))
			continue
		}

		if len(seen) == 0 {
			seen = append(seen, string(v))
			continue
		}

		// seen = append(seen, string(v))
		lastSeenInd := len(seen) - 1
		if bracketMap[string(seen[lastSeenInd])] == string(v) {
			// fmt.Println(" i m here")
			seen = seen[:lastSeenInd]
			// fmt.Printf("%v", seen)
			continue
		}

		seen = append(seen, string(v))

	}
	if len(seen) == 0 {
		return true
	}
	fmt.Printf("%v\n", seen)
	return false
}

func longestValid(s string) int {
	fmt.Printf("\ns = %s", s)
	if s == "" {
		return 0
	}

	bracketMap := map[string]string{
		"(": ")",
		"[": "]",
		"{": "}",
	}

	seen := make([]string, 0, len(s))
	seen = append(seen, string(s[0]))
	validIndices := make([]int, 0, len(s))
	for i, v := range s {
		if i == 0 {
			continue
		}
		// popped in the last iteratioin
		if len(seen) == 0 {
			seen = append(seen, string(v))
			continue
		}
		fmt.Printf("\nseen = %v", seen)

		// seen = append(seen, string(v))
		lastSeenInd := len(seen) - 1
		if bracketMap[string(seen[lastSeenInd])] == string(v) {
			validIndices = append(validIndices, i)
			seen = seen[:lastSeenInd]
			continue
		}

		seen = append(seen, string(v))
	}

	if len(validIndices) == 0 {
		return 0
	}

	if len(validIndices) == 1 {
		return 2
	}
	fmt.Printf("\nvalid indices: %v", validIndices)
	maxLongest := 2
	longestCount := 2
	for i, v := range validIndices {
		if i == len(validIndices)-1 {
			break
		}
		fmt.Printf("\ni = %d", i)
		if validIndices[i+1]-v >= 1 && validIndices[i+1]-v <= 2 {
			longestCount += 2
			if longestCount >= maxLongest {
				maxLongest = longestCount
			}
		} else {
			longestCount = 0
		}

	}

	// fmt.Printf("\nvalid indices: %v", validIndices)
	return maxLongest
}

func main() {

	testCases := []struct {
		name   string
		input  string
		result int
	}{
		{
			name:   "(()",
			input:  "(()",
			result: 2,
		},
		{
			name:   ")()())",
			input:  ")()())",
			result: 4,
		},
		{
			name:   "",
			input:  "",
			result: 0,
		},
		{
			name:   "()(()",
			input:  "()(()",
			result: 2,
		},
		{
			name:   "()(())",
			input:  "()(())",
			result: 6,
		},
	}

	for _, tc := range testCases {
		result := longestValid(tc.input)
		fmt.Printf("\ninput: %v, expected: %d, result: %d", tc.input, tc.result, result)
	}

}
