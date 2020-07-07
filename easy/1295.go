package easy

func findNumbers(nums []int) int {
	res := 0
	for _,v := range nums{
		if checkNums(v) {
			res++
		}
	}
	return res
}

func checkNums(n int) bool {
	res := true
	for n > 0 {
		res = !res
		n /= 10
	}
	return res
}