package grid

import (
	"iter"
	"slices"
)

type Location struct {
	row, col int
}

func (l Location) Down() Location {
	return At(l.row+1, l.col)
}

func (l Location) Left() Location {
	return At(l.row, l.col-1)
}

func (l Location) Right() Location {
	return At(l.row, l.col+1)
}

func At(row, col int) Location {
	return Location{row: row, col: col}
}

type Item[V comparable] struct {
	Value    V
	Location Location
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

func (i *Item[V]) Tag(tag string) {
	i.tags = append(i.tags, tag)
}

type Grid[V comparable] struct {
	Items  map[Location]*Item[V]
	Width  int
	Height int
}

func New[V comparable](lines []string, categorizer func(int32) (V, bool)) *Grid[V] {
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

func (g *Grid[V]) Set(location Location, value V) bool {
	if oldItem, ok := g.Items[location]; ok && oldItem.Value == value {
		return false
	}
	g.Items[location] = g.NewItem(location, value)
	return true
}

func (g *Grid[V]) NewItem(location Location, value V) *Item[V] {
	return &Item[V]{
		Value:    value,
		Location: location,
		parent:   g,
		tags:     make([]string, 0),
	}
}
