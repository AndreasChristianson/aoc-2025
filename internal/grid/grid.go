package grid

import (
	"iter"
)

type Location struct {
	row, col int
}

func At(row, col int) Location {
	return Location{row: row, col: col}
}

type Item[V any] struct {
	Value    V
	Location Location
	parent   *Grid[V]
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
				location := At(
					rowDelta+i.Location.row,
					colDelta+i.Location.col,
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

type Grid[V any] struct {
	Items  map[Location]*Item[V]
	Width  int
	Height int
}

func New[V any](lines []string, categorizer func(int32) (V, bool)) *Grid[V] {
	grid := &Grid[V]{
		Items:  make(map[Location]*Item[V]),
		Height: len(lines),
	}
	width := -1
	for row, line := range lines {
		width = max(width, len(line))
		for col, char := range line {
			if parsed, ok := categorizer(char); ok {
				location := At(row, col)
				grid.Items[location] = &Item[V]{
					parent:   grid,
					Location: location,
					Value:    parsed,
				}
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

func (g *Grid[V]) Locations() iter.Seq[Location] {
	return func(yield func(location Location) bool) {
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
	return g.GetByLocation(At(row, col))
}

func (g *Grid[V]) GetByLocation(at Location) (V, bool) {
	if val, ok := g.Items[at]; ok {
		return val.Value, ok
	} else {
		return *new(V), false
	}

}
