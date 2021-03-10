package __100

import "sort"

func subsetsWithDup(nums []int) (result [][]int) {
	n := len(nums)
	sort.Ints(nums)
	var dfs func(temp []int, idx int)
	dfs = func(temp []int, idx int) {
		result = append(result, append([]int(nil), temp...))
		for i := idx; i < n; i++ {
			if i > idx && nums[i] == nums[i-1] {
				continue
			}
			temp = append(temp, nums[i])
			dfs(temp, i+1)
			temp = temp[:len(temp)-1]
		}
	}
	dfs([]int{}, 0)
	return
}
