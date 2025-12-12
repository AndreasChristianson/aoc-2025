package day_12

import "aoc-2025/internal/grid"

type rectangle struct {
	width, length int
}

func (r *rectangle) area() int {
	return r.width * r.length
}

type treeConfiguration struct {
	presentShapes    []presentShape
	requiredPresents []int
	underTreeSize    rectangle
}

type presentEnum string

const (
	present = "#"
	na      = "na"
)

type presentShape struct {
	layout *grid.Grid[presentEnum]
}

func (s *presentShape) area() int {
	return len(s.layout.Items)
}
