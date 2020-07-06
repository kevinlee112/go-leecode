package easy

func GetNoZeroIntegers(n int) []int {
	for i:=1; i<n; i++ {
		if checkNum(i) {
			if checkNum(n-i) {
				return []int{i, n-i}
			}
		}
	}

	return []int{}
}

func checkNum(num int) bool {
	for num > 0 {
		if num % 10 == 0 {
			return false
		}
		num = num / 10
	}

	return true
}