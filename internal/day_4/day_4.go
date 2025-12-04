package day_4

import (
	"iter"
	"strconv"
)

type gridLocation struct {
	row int
	col int
}

type item struct {
	parent   *printFloor
	location gridLocation
	value    string
}

func (i *item) accessibleToForkLifts() bool {
	if !i.isRoll() {
		return false
	}
	neighborRollCount := 0
	for neighbor := range i.neighbors(1) {
		if neighbor.isRoll() {
			neighborRollCount++
		}
	}

	if neighborRollCount < 4 {
		return true
	}
	return false
}

func (i *item) isRoll() bool {
	return i.value == "@"
}

func (i *item) neighbors(distance int) iter.Seq[item] {
	return func(yield func(item) bool) {
		for rowDelta := -distance; rowDelta <= distance; rowDelta++ {
			for colDelta := -distance; colDelta <= distance; colDelta++ {
				if rowDelta == 0 && colDelta == 0 {
					continue
				}
				location := gridLocation{
					row: rowDelta + i.location.row,
					col: colDelta + i.location.col,
				}
				if neighbor, ok := i.parent.grid[location]; !ok {
					continue
				} else {
					if !yield(neighbor) {
						return
					}
				}
			}
		}
	}
}

type printFloor struct {
	grid map[gridLocation]item
}

func (f *printFloor) items() iter.Seq[item] {
	return func(yield func(item) bool) {
		for _, value := range f.grid {
			if !yield(value) {
				return
			}
		}
	}
}

func (f *printFloor) remove(i item) {
	delete(f.grid, i.location)
}

func parseLines(lines []string) *printFloor {
	floor := &printFloor{
		grid: make(map[gridLocation]item),
	}
	for col, line := range lines {
		for row, char := range line {
			if char == '.' {
				continue
			}
			location := gridLocation{row: row, col: col}
			floor.grid[location] = item{
				parent:   floor,
				location: location,
				value:    string(char),
			}
		}
	}
	return floor
}

func part1(lines []string) string {
	floor := parseLines(lines)
	count := 0
	for item := range floor.items() {
		if item.accessibleToForkLifts() {
			count++
		}
	}

	return strconv.Itoa(count)
}

func part2(lines []string) string {
	floor := parseLines(lines)
	count := 0
	for {
		changed := false
		for item := range floor.items() {
			if item.accessibleToForkLifts() {
				count++
				floor.remove(item)
				changed = true
			}
		}
		if !changed {
			break
		}
	}

	return strconv.Itoa(count)
}
