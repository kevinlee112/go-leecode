package __100

import (
	"sort"
)

/*
给你一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？请你找出所有满足条件且不重复的三元组。

注意：答案中不可以包含重复的三元组。

 

示例：

给定数组 nums = [-1, 0, 1, 2, -1, -4]，

满足要求的三元组集合为：
[
  [-1, 0, 1],
  [-1, -1, 2]
]
 */

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	var res [][]int
	for i:=0;i<len(nums);i++{
		if nums[i]>0 {
			break
		}
		p1, p2 := i+1, len(nums)-1
		for p1 < p2 {
			sum := nums[i] + nums[p1] + nums[p2]
			if sum > 0 {
				p2--
			}else if sum < 0 {
				p1++
			}else {
				res = append(res, []int{nums[i], nums[p1], nums[p2]})
				p1++
				p2--
			}
		}
		for p1 < p2 && nums[p1] == nums[p1+1] {
			p1++
		}
		for p1 < p2 && nums[p2] == nums[p2-1] {
			p2--
		}
		for i < len(nums)-3 && nums[i] == nums[i+1] {
			i++
		}
	}
	return res
}
