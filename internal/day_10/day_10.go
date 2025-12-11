package day_10

import (
	"strconv"
)

func part1(lines []string) string {
	machines := parseLines(lines)
	var out int64
	for _, m := range machines {
		out += m.findMinButtons()
	}
	return strconv.FormatInt(out, 10)
}

func part2(lines []string) string {
	machines := parseLines(lines)
	var out int64
	for _, m := range machines {
		out += m.findMinJoltageButtons()
	}
	return strconv.FormatInt(out, 10)
}
