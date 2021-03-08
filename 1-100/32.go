package __100
/**
给定一个只包含 '(' 和 ')' 的字符串，找出最长的包含有效括号的子串的长度。

示例 1:

输入: "(()"
输出: 2
解释: 最长有效括号子串为 "()"
示例 2:

输入: ")()())"
输出: 4
解释: 最长有效括号子串为 "()()"
))
 */


func longestValidParentheses(s string) int {
	maxAns := 0
	dp := make([]int, len(s))
	for i:=1;i<len(s);i++ {
		if s[i] == ')' {
			if s[i-1]=='(' {
				if i>=2 {
					dp[i] = dp[i-2] + 2
				}else{
					dp[i]=2
				}
			}else if i > dp[i-1] && s[i - dp[i-1] -1] == '(' {
				if i-dp[i-1] >= 2 {
					dp[i] = dp[i-1] + dp[i-dp[i-1]-2]+2
				}else {
					dp[i] = dp[i-1] + 2
				}
			}
			if maxAns < dp[i] {
				maxAns = dp[i]
			}
		}
	}
	return maxAns
}


func longestValidParentheses2(s string) int {
	maxAns := 0
	stack := []int{}
	stack = append(stack, -1)
	for i:=0;i<len(s);i++ {
		if s[i]=='(' {
			stack =append(stack, i)
		}else {
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				stack = append(stack, i)
			}else {
				if i -stack[len(stack)] > maxAns {
					maxAns = i -stack[len(stack)]
				}
			}
		}
	}
	return maxAns
}

func longestValidParentheses3(s string) int {
	left, right, maxLength := 0, 0, 0
	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			left++
		} else {
			right++
		}
		if left == right {
			if maxLength  < 2 * right{
				maxLength = 2 * right
			}
		}else if right > left {
			left, right = 0, 0
		}
	}
	left, right = 0, 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '(' {
			left++
		} else {
			right++
		}
		if left == right {
			if maxLength  < 2 * right{
				maxLength = 2 * right
			}
		} else if left > right {
			left, right = 0, 0
		}
	}
	return maxLength
}



























