package easy


func decompressRLElist(nums []int) []int {
	res := make([]int, 0)
	for i, v := range nums{
		if i%2 == 0 {
			for j:=0;j<v;j++ {
				res = append(res, nums[i+1])
			}
		}
	}
	return res
}