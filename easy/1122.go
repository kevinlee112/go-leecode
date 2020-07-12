package easy

import "sort"

func relativeSortArray(arr1 []int, arr2 []int) []int {
	m := len(arr2)
	dict := make(map[int]int)
	res := make([]int, 0)
	for _, v := range arr1{
		dict[v]++
	}
	for i:=0;i<m;i++ {
		for j:=0;j<dict[arr2[i]];j++ {
			res = append(res, arr2[i])
		}
		delete(dict, arr2[i])
	}
	last := make([]int, 0)
	for k, v := range dict{
		for j:=0;j<v;j++ {
			last = append(last, k)
		}
	}
	sort.Ints(last)
	return  append(res, last...)
}


