package easy

func plusOne(digits []int) []int {
	status := 1
	for i:=len(digits)-1;i>=0;i-- {
		digits[i] = digits[i] + status
		if digits[i] >= 10 {
			digits[i] = digits[i] % 10
			status = 1
		}else {
			status = 0
		}
	}
	if status == 1 {
		digits = append([]int{status}, digits...)
	}
	return digits
}