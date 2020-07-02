package easy

func MinCount(coins []int) int {
	ret := 0
	for _, v := range coins {
		if v%2 == 0 {
			ret += v / 2
		} else {
			ret += v/2 + 1
		}
	}

	return ret
}
