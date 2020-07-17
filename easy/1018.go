package easy


func prefixesDivBy5(A []int) []bool {
	l := len(A)
	res := make([]bool, l)
	num := 0
	for i := 0; i < l; i++ {
		num = num << 1 + A[i]
		if num % 5 == 0 {
			res[i] = true
		}
	}
	num %= 5
	return res
}













