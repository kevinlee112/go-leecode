package middle

import "fmt"

var result [][]int

// 回溯核心
// nums: 原始列表
// pathNums: 路径上的数字
// used: 是否访问过
func backtrack(nums, pathNums []int, used[]bool) {
	if len(nums) == len(pathNums) {
		result = append(result, pathNums)
		fmt.Println(result)
		return
	}

	for i:=0; i<len(nums); i++ {
		// 检查是否访问过
		if !used[i] {
			used[i] = true
			pathNums = append(pathNums, nums[i])
			backtrack(nums,pathNums,used)
			pathNums = pathNums[:len(pathNums) -1]
			used[i] = false
		}
	}
}

func Permute(nums []int) [][]int {
	var pathNums []int
	var used = make([]bool, len(nums))
	result = [][]int{}
	backtrack(nums, pathNums, used)
	return result
}