package basic

func findBinary1(s []int, k int) int {
	l, r := 0, len(s) -1
	for l <= r {
		m := (l + r) >> 1
		if s[m] < k {
			l = m + 1
		}else if s[m]> k {
			r = m -1
		}else {
			return m
		}
	}
	return -1
}

func findBinary2(s []int, k , l, r int) int {
	if len(s) < 1 {
		return -1
	}
	m := (l + r) >> 1
	if s[m] < k {
		return findBinary2(s, k, m + 1, r)
	}else if s[m]> k {
		return findBinary2(s, k, l, m -1)
	}else {
		return m
	}

}
