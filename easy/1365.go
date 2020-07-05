package easy

func SmallerNumbersThanCurrent(nums []int) []int {
	lens := len(nums)
	res := make([]int, lens)
	for i:=0;i<lens;i++ {
		for j:=0;j<lens;j++ {
			if nums[j] < nums[i] {
				res[i]++
			}

		}
	}

	return res
}