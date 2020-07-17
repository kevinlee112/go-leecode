package easy

import "sort"

func twoCitySchedCost(costs [][]int) int {
	l := len(costs)
	diff := make([]int, l)
	res := 0
	for i:=0;i<l;i++ {
		res += costs[i][0]
		diff[i] = costs[i][1]  - costs[i][0]
	}
	sort.Ints(diff)
	for i:=0;i<l/2;i++ {
		res += diff[i]
	}
	return res
}
