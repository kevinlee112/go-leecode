package easy

import "sort"

func sortedSquares(A []int) []int {
	res := make([]int, len(A))
	for i:=0;i<len(A);i++ {
		res[i] = A[i]*A[i]
	}
	sort.Ints(res)
	return res
}