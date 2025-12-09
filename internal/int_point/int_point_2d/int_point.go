package int_point_2d

type Location struct {
	Row, Col int
}

func (l Location) Down() Location {
	return At(l.Row+1, l.Col)
}

func (l Location) Left() Location {
	return At(l.Row, l.Col-1)
}

func (l Location) Right() Location {
	return At(l.Row, l.Col+1)
}

func At(row, col int) Location {
	return Location{Row: row, Col: col}
}
