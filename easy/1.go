package easy

/**
两数之和
给定一个整数数组 nums 和一个目标值 target，请你在该数组中找出和为目标值的那 两个 整数，并返回他们的数组下标。

你可以假设每种输入只会对应一个答案。但是，数组中同一个元素不能使用两遍。
 */

func twoSum(nums []int,target int) []int {
	m := make(map[int]int)
	for i,v := range nums {
		if value, exit := m[target - v]; exit {
			return []int{value, i}
		}else {
			m[v] = i
		}
	}
	return []int{}
}