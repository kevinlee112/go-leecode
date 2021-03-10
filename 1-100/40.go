package __100

import "sort"

func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	return dfs2(candidates, target)
}

func dfs2(candidates []int, target int) [][]int {
	ret := [][]int{}
	for i, d := range candidates {
		if i>0&& candidates[i-1] == d {
			continue
		} else if d > target {
			break
		} else if d == target {
			ret = append(ret, []int{d})
		}else {
			for _, v := range dfs(candidates[i+1:], target-d){
				ret = append(ret, append([]int{d}, v...))
			}
		}
	}
	return ret
}
