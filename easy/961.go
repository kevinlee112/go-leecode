package easy

import "sort"

func repeatedNTimes(A []int) int {
	sort.Ints(A)
	count := 1
	for i:=1;i<len(A);i++ {
		if A[i] == A[i-1] {
			count++
		}else {
			count=1
		}
		if count == len(A)/2 {
			return A[i]
		}
	}
	return 0
}