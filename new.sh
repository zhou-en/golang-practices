#!/usr/bin/env bash
# Create a folder for a new LeetCode problem.
# Usage: ./new.sh <number> <snake_case_title>
# Example: ./new.sh 15 3sum   -> p0015_3sum/
set -euo pipefail

if [ $# -ne 2 ]; then
	echo "usage: $0 <number> <snake_case_title>   (example: $0 15 3sum)" >&2
	exit 1
fi

num=$1
title=$2

if ! [[ $num =~ ^[0-9]+$ ]]; then
	echo "error: number must be digits only, got '$num'" >&2
	exit 1
fi
if ! [[ $title =~ ^[a-z0-9_]+$ ]]; then
	echo "error: title must be lowercase snake_case, got '$title'" >&2
	exit 1
fi

cd "$(dirname "$0")"
dir=$(printf 'p%04d_%s' "$((10#$num))" "$title")

if [ -e "$dir" ]; then
	echo "error: $dir already exists" >&2
	exit 1
fi

# LeetCode URLs use the title in kebab-case. Check the link: a few titles differ.
slug=${title//_/-}

# Get the real title (for example "3Sum") from the LeetCode API.
# If the lookup fails, build one from the folder name: two_sum -> Two Sum.
response=$(curl -s -m 10 https://leetcode.com/graphql \
	-H 'Content-Type: application/json' -H 'Referer: https://leetcode.com' \
	-d "{\"query\":\"query q(\$s:String!){question(titleSlug:\$s){questionFrontendId title}}\",\"variables\":{\"s\":\"$slug\"}}" || true)
real_title=$(printf '%s' "$response" | sed -n 's/.*"title":"\([^"]*\)".*/\1/p')
real_id=$(printf '%s' "$response" | sed -n 's/.*"questionFrontendId":"\([^"]*\)".*/\1/p')

if [ -n "$real_title" ]; then
	if [ "$real_id" != "$((10#$num))" ]; then
		echo "warning: LeetCode says '$slug' is problem $real_id, not $((10#$num))" >&2
	fi
else
	real_title=$(printf '%s' "$title" | awk -F_ '{for (i = 1; i <= NF; i++) $i = toupper(substr($i, 1, 1)) substr($i, 2); print}')
	echo "warning: LeetCode has no problem at /problems/$slug/ (or the lookup failed). Fix the title and link in description.md." >&2
fi

mkdir "$dir"

cat >"$dir/description.md" <<EOF
# $((10#$num)). $real_title

Source: https://leetcode.com/problems/$slug/

<!-- Paste the problem statement, examples, and constraints here. -->
EOF

cat >"$dir/main.go" <<'EOF'
package main

import "fmt"

// Paste the LeetCode starter function here and solve it.

func main() {
	fmt.Println("run the tests: go test .")
}
EOF

cat >"$dir/main_test.go" <<'EOF'
package main

import "testing"

func TestSolution(t *testing.T) {
	tests := []struct {
		name string
		// Add input and want fields that match the starter function.
	}{
		// Add one case per LeetCode example, for example:
		// {name: "example 1", input: ..., want: ...},
	}

	if len(tests) == 0 {
		t.Fatal("no test cases yet: add the LeetCode examples")
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// got := solve(tt.input)
			// if got != tt.want {
			// 	t.Errorf("solve(%v) = %v, want %v", tt.input, got, tt.want)
			// }
		})
	}
}
EOF

echo "created $dir/"
echo "  $dir/description.md  problem statement"
echo "  $dir/main.go         solution"
echo "  $dir/main_test.go    test cases"
echo "next: go test ./$dir"
