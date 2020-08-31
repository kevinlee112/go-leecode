package easy

func diStringMatch(S string) []int {
	left, right, result := 0, len(S), []int{}
	for _, v := range S {
		if v == 'I' {
			result = append(result, left)
			left++
		} else {
			result = append(result, right)
			right--
		}
	}
	result = append(result, left)
	return result
}