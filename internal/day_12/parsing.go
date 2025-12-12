package day_12

import (
	"aoc-2025/internal/grid"
	"aoc-2025/internal/strings"
	"regexp"
	strings2 "strings"
)

var partNumberRegex = regexp.MustCompile("^(\\d+):$")

func parseLines(lines []string) []treeConfiguration {
	presentShapes := parsePresentShapes(lines)
	ret := make([]treeConfiguration, len(lines)-30)
	for i, line := range lines[30:] {
		ret[i] = treeConfiguration{
			presentShapes:    presentShapes,
			underTreeSize:    parseRect(line),
			requiredPresents: parseRequiredPresents(line),
		}
	}
	return ret
}

var presentsRegex = regexp.MustCompile("^\\d+x\\d+: ([ \\d]+)$")

func parseRequiredPresents(line string) []int {
	matches := presentsRegex.FindStringSubmatch(line)
	presentRequirements := strings2.Split(matches[1], " ")
	ret := make([]int, len(presentRequirements))
	for i, presentRequirement := range presentRequirements {
		ret[i] = strings.MustParse(presentRequirement)
	}
	return ret
}

func parsePresentShapes(lines []string) []presentShape {
	ret := make([]presentShape, 6)
	for i := 0; i < 30; i += 5 {
		number := partNumberRegex.FindStringSubmatch(lines[i])[1]
		gridLines := lines[i+1 : i+4]
		ret[strings.MustParse(number)] = presentShape{
			layout: grid.New(gridLines, parsePresentShape),
		}
	}
	return ret
}

var rectRegex = regexp.MustCompile("^(\\d+)x(\\d+): .*$")

func parseRect(line string) rectangle {
	matches := rectRegex.FindStringSubmatch(line)
	return rectangle{
		width:  strings.MustParse(matches[1]),
		length: strings.MustParse(matches[2]),
	}
}

func parsePresentShape(char int32) (presentEnum, bool) {
	switch char {
	case '#':
		return present, true
	default:
		return na, false
	}
}
