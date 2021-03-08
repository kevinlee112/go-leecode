package middle

import "sort"

func nextPermutation(nums []int)  {
	if len(nums) < 2 {
		return
	}
	i := len(nums)-2
	for i >= 0&&nums[i]>=nums[i+1] {
		i--
	}
	if i==-1 {
		sort.Ints(nums)
	}else {
		j := len(nums)-1
		for j>=0&&nums[j]<=nums[i]{
			j--
		}
		nums[i], nums[j] = nums[j], nums[i]
		sort.Ints(nums[i+1:])

	}
}