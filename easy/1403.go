package easy

import "sort"

func MinSubsequence(nums []int) []int {
	res := make([]int, 0)
	sum, subSum := 0, 0
	sort.Ints(nums)
	for _, v := range nums {
		sum += v
	}
	lens := len(nums)
	for i:=0; i< lens;i++ {
		subSum += nums[lens-1-i]
		res = append(res, nums[lens-1-i])
		if subSum > sum/2 {
			break
		}
	}
	return  res
}