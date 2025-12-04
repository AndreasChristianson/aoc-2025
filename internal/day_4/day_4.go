package day_4

import (
	"aoc-2025/internal/grid"
	"strconv"
)

func part1(lines []string) string {
	floor := grid.New(lines, parseRoll)
	count := 0
	for item := range floor.Values() {
		if accessibleToForkLifts(item) {
			count++
		}
	}

	return strconv.Itoa(count)
}

func accessibleToForkLifts(i *grid.Item[*roll]) bool {
	neighborRollCount := 0
	for neighbor := range i.Neighbors(1) {
		if neighbor.Value.isRoll() {
			neighborRollCount++
		}
	}
	if neighborRollCount < 4 {
		return true
	}
	return false
}

type roll struct {
	value string
}

func (i *roll) isRoll() bool {
	return i.value == "@"
}

func parseRoll(char int32) (*roll, bool) {
	asString := string(char)
	if asString == "." {
		return nil, false
	}
	return &roll{value: asString}, true

}

func part2(lines []string) string {
	floor := grid.New(lines, parseRoll)

	count := 0
	for {
		changed := false
		for item := range floor.Values() {
			if accessibleToForkLifts(item) {
				count++
				floor.Remove(item)
				changed = true
			}
		}
		if !changed {
			break
		}
	}

	return strconv.Itoa(count)
}
