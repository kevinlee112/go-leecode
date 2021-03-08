package __100

import "sort"

/**
给定一个包括 n 个整数的数组 nums 和 一个目标值 target。找出 nums 中的三个整数，使得它们的和与 target 最接近。返回这三个数的和。假定每组输入只存在唯一答案。

 

示例：

输入：nums = [-1,2,1,-4], target = 1
输出：2
解释：与 target 最接近的和是 2 (-1 + 2 + 1 = 2) 。

 */
func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	res := target
	for i:=0;i<len(nums);i++ {
		p1, p2 := i+1, len(nums)-1
		diff := target
		for p1<p2 {
			diff = target - nums[i] - nums [p1] - nums[p2]
			if diff >0 {
				p1++
				if res > diff  {
					res = diff
				}
			}else if diff < 0 {
				p2++
				if res + diff > 0 {
					res = diff * -1
				}
			}else {
				return 0
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
	}
	return res
}