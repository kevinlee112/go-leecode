package middle

import "sort"

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	var ans [][]int
	for i := 0; i < len(nums)-2; i++ {
		if nums[i] > 0 {
			break
		}
		for i < len(nums)-3 && nums[i] == nums[i+1] {
			i++
		}
		p1, p2 := i+1, len(nums)-1
		if nums[p2] < 0 {
			break
		}
		for p1 < p2 {
			for p1 < p2 && nums[p1] == nums[p1+1] {
				p1++
			}
			for p1 < p2 && nums[p2] == nums[p2-1] {
				p2--
			}
			sum := nums[i] + nums[p1] + nums[p2]
			if sum < 0 {
				p1++
			}else if sum > 0 {
				p2--
			}else{
				ans = append(ans, []int{nums[i], nums[p1], nums[p2]})
				// 移动到重复元素的最后一个位置上
				p1++
				p2--
			}
		}
	}
	return ans
}