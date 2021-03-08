package __100

import "sort"

/**
1. 两数之和
给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那 两个 整数，并返回他们的数组下标。
你可以假设每种输入只会对应一个答案。但是，数组中同一个元素不能使用两遍。
示例:
给定 nums = [2, 7, 11, 15], target = 9
因为 nums[0] + nums[1] = 2 + 7 = 9
所以返回 [0, 1]
[3,2,4]
6
 */

func towSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, v:= range nums{
		if value, exit := m[target - v];exit {
			return []int{value, i}
		}else {
			m[v] = i
		}
	}

	return []int{}
}

func towSum2(nums []int, target int) []int {
	sort.Ints(nums)
	l, r := 0, len(nums) - 1
	for l < r {
		if nums[l] + nums[r] > target {
			r--
		}else if nums[l] + nums[r] < target {
			l++
		}else {
			return []int{l, r}
		}
	}
	return []int{}
}