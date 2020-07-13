package easy

import "sort"

func numMovesStones(a int, b int, c int) []int {
	s := []int{a, b, c}
	sort.Ints(s)
	res := make([]int, 2)
	if (s[2] - s[1] == 1) && (s[1] - s[0] == 1) {
		res[0] = 0
	}else if (s[2] - s[1] > 2) && (s[1] - s[0] == 1) {
		res[0] = 1
	} else if (s[1] - s[0] > 2) && (s[2] - s[1] == 1) {
		res[0] = 1
	} else if (s[1] - s[0] == 2) && (s[2] - s[1] == 1) {
		res[0] = 1
	} else if (s[2] - s[1] > 2) && (s[1] - s[0] == 1) {
		res[0] = 1
	} else {
		res[0] = 2
	}
	for i:=1;i<3;i++ {
		res[1] += s[i]-s[i-1] -1
	}
	return res
}
