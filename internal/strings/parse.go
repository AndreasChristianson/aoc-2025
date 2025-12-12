package strings

import "strconv"

func MustParse(s string) int {
	if i, err := strconv.Atoi(s); err != nil {
		panic(err)
	} else {
		return i
	}
}

func MustParse64(s string) int64 {
	if i, err := strconv.ParseInt(s, 10, 64); err != nil {
		panic(err)
	} else {
		return i
	}
}
