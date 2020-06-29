package easy

func Shuffle(nums []int, n int) []int {
	ret := make([]int, 2 * n)
	arr1 := nums[0 : n]
	arr2 := nums[n : len(nums)]
	for i := 0; i < n; i++ {
		ret[2*i] = arr1[i]
		ret[2*i+1] = arr2[i]
	}
	return ret
}