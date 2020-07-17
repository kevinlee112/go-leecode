package easy

func canThreePartsEqualSum(A []int) bool {
	sum, subSum, count := 0, 0, 0
	for i:=0;i<len(A);i++ {
		sum+= A[i]
	}
	if sum % 3 != 0 {
		return false
	}else {
		subSum = sum / 3
	}
	for i:=0;i<len(A)-1;i++ {
		subSum -= A[i]
		if subSum == 0 {
			count++
			subSum = sum / 3
		}
		if count == 2 {
			return true
		}
	}

	return false
}