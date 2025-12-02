package strings

import (
	"aoc-2025/internal/ints"
)

func IsPalindrome(input string) bool {
	return input[:len(input)/2] == Reverse(input[ints.CeilDivide(len(input), 2):])
}
