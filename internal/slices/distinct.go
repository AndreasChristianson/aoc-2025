package slices

func Unique[V comparable](input []V) []V {
	seen := make(map[V]bool)
	result := make([]V, 0)

	for _, item := range input {
		if _, ok := seen[item]; !ok {
			seen[item] = true
			result = append(result, item)
		}
	}
	return result
}
