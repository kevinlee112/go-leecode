package __100


/**
给你一个升序排列的整数数组 nums ，和一个整数 target 。

假设按照升序排序的数组在预先未知的某个点上进行了旋转。（例如，数组 [0,1,2,4,5,6,7] 可能变为 [4,5,6,7,0,1,2] ）。

请你在数组中搜索 target ，如果数组中存在这个目标值，则返回它的索引，否则返回 -1 。

 */
func search(nums []int, target int) int {
	left, right :=0, len(nums)-1
	for left <= right {
		mid := (right + left)/2
		if nums[mid] == target {
			return mid
		}
		if nums[left] == target {
			return left
		}
		if nums[right] == target {
			return right
		}
		if nums[left] < nums[mid] {
			if target < nums[mid] && target > nums[left] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else if nums[mid] < nums[left] {
			if target > nums[mid] && target < nums[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}
	return  -1
}