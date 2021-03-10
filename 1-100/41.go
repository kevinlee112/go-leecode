package __100


func firstMissingPositive(nums []int) int {
	n:=len(nums)
	//边界情况
	if n==0{
		return 1
	}
	for i:=range nums{
		for nums[i] != i+1 && nums[i] < n && nums[i] >= 1 && nums[i] != nums[nums[i]-1] { //小心死循环
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
		}
	}
	for i:=range nums{
		if nums[i]!=i+1{
			return i+1
		}
	}
	return nums[n-1]+1//数组已经是从1开始的非递减序列
}