package easy

func RunningSum(nums []int) []int {
	lens := len(nums)
	if lens <= 1 {
		return nums
	}

	for i:=1;i<lens;i++{
		nums[i] = nums[i-1] + nums[i]
	}

	return nums
}