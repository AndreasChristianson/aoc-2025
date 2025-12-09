package grid

import (
	"aoc-2025/internal/int_point/int_point_2d"
	"iter"
	"slices"
)

type Item[V comparable] struct {
	Value    V
	Location int_point_2d.Location
	parent   *Grid[V]
	tags     []string
}

func (i *Item[V]) Remove() {
	i.parent.Remove(i)
}

func (i *Item[V]) Neighbors(distance int) iter.Seq[*Item[V]] {
	return func(yield func(*Item[V]) bool) {
		for rowDelta := -distance; rowDelta <= distance; rowDelta++ {
			for colDelta := -distance; colDelta <= distance; colDelta++ {
				if rowDelta == 0 && colDelta == 0 {
					continue
				}
				location := int_point_2d.At(
					rowDelta+i.Location.Row,
					colDelta+i.Location.Col,
				)
				if neighbor, ok := i.parent.Items[location]; !ok {
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

func (i *Item[V]) Tag(tag string) {
	i.tags = append(i.tags, tag)
}

type Grid[V comparable] struct {
	Items  map[int_point_2d.Location]*Item[V]
	Width  int
	Height int
}

func New[V comparable](lines []string, categorizer func(int32) (V, bool)) *Grid[V] {
	grid := &Grid[V]{
		Items:  make(map[int_point_2d.Location]*Item[V]),
		Height: len(lines),
	}
	width := -1
	for row, line := range lines {
		width = max(width, len(line))
		for col, char := range line {
			if parsed, ok := categorizer(char); ok {
				location := int_point_2d.At(row, col)
				grid.Items[location] = grid.NewItem(location, parsed)
			}
		}
	}
	grid.Width = width
	return grid
}

func (g *Grid[V]) Values() iter.Seq[*Item[V]] {
	return func(yield func(*Item[V]) bool) {
		for _, value := range g.Items {
			if !yield(value) {
				return
			}
		}
	}
}

func (g *Grid[V]) Locations() iter.Seq[int_point_2d.Location] {
	return func(yield func(location int_point_2d.Location) bool) {
		for key, _ := range g.Items {
			if !yield(key) {
				return
			}
		}
	}
}

func (g *Grid[V]) Remove(i *Item[V]) {
	delete(g.Items, i.Location)
}

func (g *Grid[V]) Get(row int, col int) (V, bool) {
	return g.GetByLocation(int_point_2d.At(row, col))
}

func (g *Grid[V]) GetByLocation(at int_point_2d.Location) (V, bool) {
	if val, ok := g.Items[at]; ok {
		return val.Value, ok
	} else {
		return *new(V), false
	}

}

func (g *Grid[V]) Find(value V) iter.Seq[*Item[V]] {
	return func(yield func(item *Item[V]) bool) {
		for _, item := range g.Items {
			if item.Value == value {
				if !yield(item) {
					return
				}
			}
		}
	}
}

func (g *Grid[V]) FindWithoutTag(value V, tag string) iter.Seq[*Item[V]] {
	return func(yield func(item *Item[V]) bool) {
		for item := range g.Find(value) {
			if !slices.Contains(item.tags, tag) {
				if !yield(item) {
					return
				}
			}
		}
	}

}

func (g *Grid[V]) Set(location int_point_2d.Location, value V) bool {
	if oldItem, ok := g.Items[location]; ok && oldItem.Value == value {
		return false
	}
	g.Items[location] = g.NewItem(location, value)
	return true
}

func (g *Grid[V]) NewItem(location int_point_2d.Location, value V) *Item[V] {
	return &Item[V]{
		Value:    value,
		Location: location,
		parent:   g,
		tags:     make([]string, 0),
	}
}
