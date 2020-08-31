package easy

func validMountainArray(A []int) bool {
	ALen := len(A)
	if ALen < 3 {
		return false
	}
	//是否上坡
	isUp := false
	if A[0] < A[1] {
		isUp = true
	} else {
		return false
	}

	for i := 1; i < ALen && i+1 < ALen; i++ {
		if isUp {
			//平地，返回false
			if A[i+1] == A[i] {
				return false
			}
			//下坡了
			if A[i+1] < A[i] {
				isUp = false
				continue
			}
			continue
		}
		//下坡
		if A[i+1] >= A[i] {
			return false
		}
	}
	if isUp == true {
		return false
	}
	return true
}