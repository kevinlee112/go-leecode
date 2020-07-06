package easy

import "sort"

func arrayRankTransform(arr []int) []int {
	l := len(arr)
	dict := make(map[int]int, l)
	cArr := make([]int, l)
	copy(cArr, arr)
	sort.Ints(cArr)
	res := make([]int, 0)
	diff := 0
	for i, v := range cArr{
		if i > 0 && v == cArr[i-1] {
			diff ++
		}
		dict[v] = i + 1 - diff
	}
	for _, v := range arr{
		res = append(res, dict[v])
	}
	return res
}