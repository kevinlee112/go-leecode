package easy



func removeOuterParentheses(S string) string {
	res := ""
	num := 0
	index := 0
	for k, v := range S {
		if string(v) == "(" {
			num++
		} else {
			num--
		}
		if num == 1 && string(v) == "(" {
			index = k
		}
		if num == 0 {
			res += S[index+1:k]
		}
	}
	return res
}
