package day_10

import (
	strings2 "aoc-2025/internal/strings"
	"regexp"
	"strings"
)

func parseLines(lines []string) []*machine {
	ret := make([]*machine, len(lines))
	for i, line := range lines {
		ret[i] = parseLine(line)
	}
	return ret
}

var lineRegex = regexp.MustCompile(`^(\[.*]) (\(.*\)) (\{.*})$`)
var displayRegex = regexp.MustCompile(`^\[([#.]+)]$`)

func parseLine(line string) *machine {
	matches := lineRegex.FindStringSubmatch(line)
	if matches == nil || len(matches) != 4 {
		panic("bad line")
	}
	return &machine{
		display:                parseDisplay(matches[1]),
		buttons:                parseButtons(matches[2]),
		joltages:               parseJoltages(matches[3]),
		cachedHighJoltageIndex: -1,
	}
}

var joltageRegex = regexp.MustCompile(`^\{([,\d]+)}$`)

func parseJoltages(joltages string) []int {
	matches := joltageRegex.FindStringSubmatch(joltages)
	joltageCostIndices := strings.Split(matches[1], ",")
	ret := make([]int, len(joltageCostIndices))
	for i := range ret {
		ret[i] = strings2.MustParse(joltageCostIndices[i])
	}
	return ret
}

var buttonsRegex = regexp.MustCompile(`^\(([() ,\d]*)\)$`)

func parseButtons(buttonString string) [][]int {
	matches := buttonsRegex.FindStringSubmatch(buttonString)
	buttons := strings.Split(matches[1], ") (")
	ret := make([][]int, len(buttons))
	for i := range ret {
		ret[i] = parseButton(buttons[i])
	}
	return ret
}

var buttonRegex = regexp.MustCompile(`^\([\d,]+\)$`)

func parseButton(buttonString string) []int {
	buttonChangeIndices := strings.Split(buttonString, ",")
	ret := make([]int, len(buttonChangeIndices))
	for i := range ret {
		ret[i] = strings2.MustParse(buttonChangeIndices[i])
	}
	return ret
}

func parseDisplay(display string) []bool {
	matches := displayRegex.FindStringSubmatch(display)
	indicators := strings.Split(matches[1], "")
	ret := make([]bool, len(indicators))
	for i, indicator := range indicators {
		switch indicator {
		case "#":
			ret[i] = true
		case ".":
			ret[i] = false
		default:
			panic("invalid display format")
		}
	}
	return ret
}
