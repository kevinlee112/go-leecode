package easy

import (
	"sort"
)

func KWeakestRows(mat [][]int, k int) []int {
	m := len(mat)
	n := len(mat[0])
	dict := make(map[int]int, m)
	s := make([]int, n)
	res := make([]int, 0)
	for i:=0;i<m;i++ {
		sum := 0
		for j:=0;j<n;j++ {
			if mat[i][i] == 1{
				sum++
			}
		}
		dict[sum*100 + i] = i
		s[i] = sum*100 + i
	}
	sort.Ints(s)
	for i, v:= range s{
		res = append(res, dict[v])
		if i>=k-1 {
			break
		}
	}
	return res
}