package easy

import "math"

func isAlienSorted(words []string, order string) bool {
	m := make(map[int32]int)
	for k, v := range order {
		m[v] = k
	}
	for i := 1; i < len(words); i++ {
		if !alienBigger(words[i], words[i-1], m) {
			return false
		}
	}
	return true
}

func alienBigger(a, b string, m map[int32]int) bool {
	la, lb := len(a), len(b)
	l := int(math.Min(float64(la), float64(lb)))
	for i := 0; i < l; i++ {
		if m[int32(a[i])] > m[int32(b[i])]{
			return true
		} else if m[int32(a[i])] < m[int32(b[i])] {
			return false
		}
	}
	if la > lb {
		return true
	}
	return false
}
