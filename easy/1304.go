package easy

func sumZero(n int) []int {
	res := make([]int, n)
	sum := 0
	for i:=0;i<n-1;i++ {
		res[i] = i
		sum+=i

	}
	res[n-1] = 0-sum
	return res
}