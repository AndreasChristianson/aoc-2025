package int_range

type IntRange struct {
	Min, Max int64
}

func (r *IntRange) Contains(value int64) bool {
	return value >= r.Min && value <= r.Max
}

func (r *IntRange) Size() int64 {
	return r.Max - r.Min + 1
}

func (r *IntRange) Combine(other *IntRange) (*IntRange, bool) {
	if r.Contains(other.Min) && r.Contains(other.Max) {
		return r, true
	}
	if other.Contains(r.Min) && other.Contains(r.Max) {
		return other, true
	}
	if r.Contains(other.Min) && other.Contains(r.Max) {
		return New(r.Min, other.Max), true
	}
	if r.Contains(other.Max) && other.Contains(r.Min) {
		return New(other.Min, r.Max), true
	}
	return nil, false
}

func New(minRange int64, maxRange int64) *IntRange {
	if maxRange < minRange {
		panic("max range must be greater than min range")
	}
	return &IntRange{minRange, maxRange}
}
