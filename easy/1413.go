package easy

func MinStartValue(nums []int) int {
	ret, sum := 0, 0
	for _,v := range nums{
		sum += v
		if sum < ret {
			ret = sum
		}
	}
	return -ret + 1
}


