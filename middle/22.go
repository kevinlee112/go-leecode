package middle

func generateParenthesis(n int) []string {
	res := new([]string)
	backtracking(n, n, "",  res)
	return *res
}

func backtracking(left, right int, tmp string, res *[]string) {
	if right == 0 {
		*res = append(*res, tmp)
		return
	}

	// 生成左括号
	if left > 0 {
		backtracking(left - 1, right, tmp + "(", res)
	}

	// 括号成对存在，有左括号才会有右括号
	if right > left {
		backtracking(left, right - 1, tmp + ")", res)
	}
}
