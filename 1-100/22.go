package __100

/**
数字 n 代表生成括号的对数，请你设计一个函数，用于能够生成所有可能的并且 有效的 括号组合。
示例：

输入：n = 3
输出：[
       "((()))",
       "(()())",
       "(())()",
       "()(())",
       "()()()"
     ]
 */


func generateParenthesis(n int) []string {
	res := &[]string{}
	backtracking(n, n, "", res)
	return *res
}

func backtracking(l, r int, tmp string, res *[]string)  {
	if r == 0 {
		*res = append(*res, tmp)
		return
	}
	if l > 0 {
		backtracking(l-1, r, tmp + "(", res)
	}
	if r > l {
		backtracking(l, r-1, tmp + ")", res)
	}
}