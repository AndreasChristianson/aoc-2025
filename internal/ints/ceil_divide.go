package ints

func CeilDivide(numerator, denominator int) int {
	if numerator%denominator == 0 {
		return numerator / denominator
	}
	return numerator/denominator + 1
}
