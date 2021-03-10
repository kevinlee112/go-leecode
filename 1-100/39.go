package __100

import "sort"

func combinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	return dfs(candidates, target)
}

func dfs(candidates []int, target int) [][]int {
	ret := [][]int{}
	for i, d := range candidates {
		if d > target {
			break
		} else if d == target {
			ret = append(ret, []int{d})
		}else {
			for _, v := range dfs(candidates[i:], target-d){
				ret = append(ret, append([]int{d}, v...))
			}
		}
	}
	return ret
}




