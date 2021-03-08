package __100

/**
给定一个按照升序排列的整数数组 nums，和一个目标值 target。找出给定目标值在数组中的开始位置和结束位置。

你的算法时间复杂度必须是 O(log n) 级别。

如果数组中不存在目标值，返回 [-1, -1]。

示例 1:

输入: nums = [5,7,7,8,8,10], target = 8
输出: [3,4]
示例 2:

输入: nums = [5,7,7,8,8,10], target = 6
输出: [-1,-1]
 */
func searchRange(nums []int, target int) []int {
	if len(nums) < 1 {
		return []int{-1, -1}
	}
	index := findBinary(nums, 0, len(nums)-1, target)
	if index == -1 {
		return []int{-1, -1}
	}
	min,max:=index,index
	for {
		//注意不要下标溢出
		if min-1>=0 && nums[min-1]==nums[index]{
			min--
			continue
		}
		break
	}
	for {
		//注意不要下标溢出
		if max+1<=len(nums)-1 && nums[max+1]==nums[index]{
			max++
			continue
		}
		break
	}
	return []int{min,max}
}

func findBinary(nums []int, min, max, target int) int {
	if nums[min]==target{
		return min
	}
	if nums[max]==target{
		return max
	}
	if max-min<=1{
		return -1
	}
	mid := (min + max) / 2
	if nums[mid]>target{
		return findBinary(nums,min,mid-1,target)
	}else if nums[mid]< target {
		return findBinary(nums,mid+1,max,target)
	}else{
		return mid
	}
}