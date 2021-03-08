package easy

func removeDuplicates(nums []int) int {
	if len(nums) == 0  {
		return 0
	}
	n := 0
	for i:=1;i<len(nums);i++{
		if nums[i] != nums[n] {
			n++
			nums[n] = nums[i]
		}
	}
	nums = nums[0:n+1]
	return len(nums)
}
