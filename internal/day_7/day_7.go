package day_7

import (
	"aoc-2025/internal/grid"
	"strconv"
)

func part1(lines []string) string {
	field := grid.New(lines, spliterParser)
	var splits int64
	for item := range field.Find(source) {
		field.Set(item.Location, tachyon)
		break
	}
	for {
		var found bool
		for beam := range field.FindWithoutTag(tachyon, "propagated") {
			found = true
			beam.Tag("propagated")
			if below, found := field.GetByLocation(beam.Location.Down()); !found {
				continue
			} else if below == splitter {
				field.Set(beam.Location.Down().Right(), tachyon)
				field.Set(beam.Location.Down().Left(), tachyon)
				splits++
			} else {
				field.Set(beam.Location.Down(), tachyon)
			}
		}
		if !found {
			break
		}
	}
	return strconv.FormatInt(splits, 10)
}

type object string

const (
	splitter object = "^"
	source   object = "S"
	empty    object = "."
	tachyon  object = "|"
)

func spliterParser(char int32) (object, bool) {
	if char == '^' {
		return splitter, true
	}
	if char == 'S' {
		return source, true
	}
	return empty, true
}

func part2(lines []string) string {
	var sum int64

	return strconv.FormatInt(sum, 10)
}
