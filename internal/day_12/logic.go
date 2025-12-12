package day_12

func (t *treeConfiguration) presentsFit() bool {
	space := t.underTreeSize.area()
	var minNeededSpace int
	for i, requiredPresentCount := range t.requiredPresents {
		minNeededSpace += requiredPresentCount * t.presentShapes[i].area()
	}
	return space >= minNeededSpace
}
