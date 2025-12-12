package day_12

import (
	"strconv"
)

func part1(lines []string) (output string) {
	trees := parseLines(lines)
	var result int64
	for _, tree := range trees {
		if tree.presentsFit() {
			result++
		}
	}
	return strconv.FormatInt(result, 10)
}
