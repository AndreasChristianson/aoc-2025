package day_9

import (
	"aoc-2025/internal/int_point/int_point_2d"
	"aoc-2025/internal/ints"
	aocslices "aoc-2025/internal/slices"
	"aoc-2025/internal/strings"
	"iter"
	"regexp"
	"slices"
	"strconv"
)

type rectangle struct {
	topLeft, bottomRight int_point_2d.Location
}

func newRectangle(corner1, corner2 int_point_2d.Location) *rectangle {
	return &rectangle{
		topLeft: int_point_2d.Location{
			Row: min(corner1.Row, corner2.Row),
			Col: min(corner1.Col, corner2.Col),
		},
		bottomRight: int_point_2d.Location{
			Row: max(corner1.Row, corner2.Row),
			Col: max(corner1.Col, corner2.Col),
		},
	}
}

func (r *rectangle) corners() []int_point_2d.Location {
	corner3 := int_point_2d.Location{
		Row: r.topLeft.Row,
		Col: r.bottomRight.Col,
	}
	corner4 := int_point_2d.Location{
		Row: r.bottomRight.Row,
		Col: r.topLeft.Col,
	}
	return []int_point_2d.Location{
		r.topLeft,
		r.bottomRight,
		corner3,
		corner4,
	}
}

func (r *rectangle) border() iter.Seq[int_point_2d.Location] {
	return func(yield func(int_point_2d.Location) bool) {
		for col := r.topLeft.Col + 1; col < r.bottomRight.Col; col++ {

			if !yield(int_point_2d.Location{
				Row: r.topLeft.Row + 1,
				Col: col,
			}) {
				return
			}
			if !yield(int_point_2d.Location{
				Row: r.bottomRight.Row - 1,
				Col: col,
			}) {
				return
			}
		}
		for row := r.topLeft.Row + 1; row < r.bottomRight.Row; row++ {

			if !yield(int_point_2d.Location{
				Row: row,
				Col: r.topLeft.Col + 1,
			}) {
				return
			}
			if !yield(int_point_2d.Location{
				Row: row,
				Col: r.bottomRight.Col - 1,
			}) {
				return
			}
		}
	}
}

type line struct {
	topLeft, bottomRight int_point_2d.Location
}

func newLine(point1, point2 int_point_2d.Location) *line {
	if point1.Col < point2.Col {
		return &line{
			topLeft:     point1,
			bottomRight: point2,
		}
	}
	if point1.Row < point2.Row {
		return &line{
			topLeft:     point1,
			bottomRight: point2,
		}
	}
	return &line{
		topLeft:     point2,
		bottomRight: point1,
	}
}

func (l line) horizontal() bool {
	return l.topLeft.Row == l.bottomRight.Row
}

func (l line) contains(point int_point_2d.Location) bool {
	if l.horizontal() {
		return l.topLeft.Row == point.Row && l.topLeft.Col <= point.Col && l.bottomRight.Col >= point.Col
	}
	return l.topLeft.Col == point.Col && l.topLeft.Row <= point.Row && l.bottomRight.Row >= point.Row

}

type tilePattern struct {
	redTiles         []int_point_2d.Location
	rectanglesBySize map[int64][]*rectangle
	rectangleSizes   []int64
	lines            []*line
	insideMap        map[int_point_2d.Location]bool
}

func (p *tilePattern) computeInside(point int_point_2d.Location) bool {
	var crossings int64
	for _, l := range p.lines {
		if l.contains(point) {
			return true
		}
		if l.horizontal() && l.topLeft.Row < point.Row && l.topLeft.Col < point.Col && l.bottomRight.Col >= point.Col {
			crossings++
		}
	}
	return crossings%2 == 1
}
func (p *tilePattern) inside(point int_point_2d.Location) bool {
	if value, found := p.insideMap[point]; found {
		return value
	} else {
		ret := p.computeInside(point)
		p.insideMap[point] = ret
		return ret
	}
}

func part1(lines []string) string {
	points := parse(lines)
	pattern := makePattern(points)
	var out int64
	out = pattern.rectangleSizes[len(pattern.rectangleSizes)-1]
	return strconv.FormatInt(out, 10)
}

func makePattern(points []int_point_2d.Location) *tilePattern {
	ret := &tilePattern{
		rectanglesBySize: make(map[int64][]*rectangle),
		rectangleSizes:   make([]int64, 0),
		redTiles:         points,
		lines:            joinPoints(points),
		insideMap:        make(map[int_point_2d.Location]bool),
	}
	for leftIndex := 0; leftIndex < len(points)-1; leftIndex++ {
		for rightIndex := leftIndex + 1; rightIndex < len(points); rightIndex++ {
			rightPoint := points[rightIndex]
			leftPoint := points[leftIndex]
			area := (ints.Abs(int64(rightPoint.Row-leftPoint.Row)) + 1) *
				(ints.Abs(int64(rightPoint.Col-leftPoint.Col)) + 1)
			rect := newRectangle(leftPoint, rightPoint)
			ret.rectangleSizes = append(ret.rectangleSizes, area)
			if rects, found := ret.rectanglesBySize[area]; found {
				ret.rectanglesBySize[area] = append(rects, rect)
			} else {
				ret.rectanglesBySize[area] = []*rectangle{rect}
			}
		}
	}
	ret.rectangleSizes = aocslices.Unique(ret.rectangleSizes)
	slices.Sort(ret.rectangleSizes)
	return ret
}

func joinPoints(points []int_point_2d.Location) []*line {
	lines := make([]*line, len(points))
	lines[0] = newLine(
		points[0],
		points[len(points)-1],
	)
	for i := 1; i < len(points); i++ {
		lines[i] = newLine(
			points[i],
			points[i-1],
		)
	}
	return lines
}

func parse(lines []string) []int_point_2d.Location {
	ret := make([]int_point_2d.Location, len(lines))
	for i, line := range lines {
		ret[i] = parseLine(line)
	}
	return ret
}

var lineRegex = regexp.MustCompile(`^(\d+),(\d+)$`)

func parseLine(line string) int_point_2d.Location {
	matches := lineRegex.FindStringSubmatch(line)
	return int_point_2d.Location{
		Col: strings.MustParse(matches[1]),
		Row: strings.MustParse(matches[2]),
	}
}

func part2(lines []string) string {
	points := parse(lines)
	pattern := makePattern(points)
	var out int64
main:
	for sizeIndex := len(pattern.rectangleSizes) - 1; sizeIndex >= 0; sizeIndex-- {
		size := pattern.rectangleSizes[sizeIndex]
		rects := pattern.rectanglesBySize[size]
	nextRect:
		for _, rect := range rects {
			for point := range rect.border() {
				if !pattern.inside(point) {
					continue nextRect
				}
			}
			out = size
			break main
		}

	}

	return strconv.FormatInt(out, 10)
}
