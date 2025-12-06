package day_5

import (
	"aoc-2025/internal/int_range"
	"regexp"
	"sort"
	"strconv"
)

func part1(lines []string) string {
	freshRanges, items := parse(lines)
	count := 0
	for _, item := range items {
		for _, freshRange := range freshRanges {
			if freshRange.Contains(item) {
				count++
				break
			}
		}
	}

	return strconv.Itoa(count)
}

func parse(lines []string) ([]*int_range.IntRange, []int64) {
	var intMode bool
	ranges := make([]*int_range.IntRange, 0)
	ints := make([]int64, 0)
	for _, line := range lines {
		if line == "" {
			intMode = true
			continue
		}
		if intMode == false {
			ranges = append(ranges, parseRange(line))
		} else {
			ints = append(ints, parseId(line))
		}
	}
	return ranges, ints
}

func parseId(line string) int64 {
	if id, err := strconv.ParseInt(line, 10, 64); err != nil {
		panic(err)
	} else {
		return id
	}
}

var rangeRegex = regexp.MustCompile("(\\d+)-(\\d+)")

func parseRange(line string) *int_range.IntRange {
	matches := rangeRegex.FindStringSubmatch(line)
	if minFresh, err := strconv.ParseInt(matches[1], 10, 64); err != nil {
		panic(err)
	} else if maxFresh, err := strconv.ParseInt(matches[2], 10, 64); err != nil {
		panic(err)
	} else {
		return int_range.New(minFresh, maxFresh)
	}
}

func part2(lines []string) string {
	freshRanges, _ := parse(lines)
	combinedRanges := make([]*int_range.IntRange, 0)
	sort.Slice(freshRanges, func(i, j int) bool {
		return freshRanges[i].Min < freshRanges[j].Min
	})

	for len(freshRanges) > 0 {
		unprocessed := make([]*int_range.IntRange, 0)
		current := freshRanges[0]
		for i := 1; i < len(freshRanges); i++ {
			toBeProcessed := freshRanges[i]
			if newCombination, ok := current.Combine(toBeProcessed); ok {
				current = newCombination
			} else {
				unprocessed = append(unprocessed, toBeProcessed)
			}
		}
		combinedRanges = append(combinedRanges, current)
		freshRanges = unprocessed
	}

	var count int64
	for _, freshRange := range combinedRanges {
		count += freshRange.Size()
	}

	return strconv.FormatInt(count, 10)
}
