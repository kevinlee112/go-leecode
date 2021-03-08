package __100


/**
给你一个字符串 s 和一个字符规律 p，请你来实现一个支持 '.' 和 '*' 的正则表达式匹配。

'.' 匹配任意单个字符
'*' 匹配零个或多个前面的那一个元素
所谓匹配，是要涵盖 整个 字符串 s的，而不是部分字符串。

 
示例 1：

输入：s = "aa" p = "a"
输出：false
解释："a" 无法匹配 "aa" 整个字符串。
示例 2:

输入：s = "aa" p = "a*"
输出：true
解释：因为 '*' 代表可以匹配零个或多个前面的那一个元素, 在这里前面的元素就是 'a'。因此，字符串 "aa" 可被视为 'a' 重复了一次。
示例 3：

输入：s = "ab" p = ".*"
输出：true
解释：".*" 表示可匹配零个或多个（'*'）任意字符（'.'）。
示例 4：

输入：s = "aab" p = "c*a*b"
输出：true
解释：因为 '*' 表示零个或多个，这里 'c' 为 0 个, 'a' 被重复一次。因此可以匹配字符串 "aab"。
示例 5：

输入：s = "mississippi" p = "mis*is*p*."
输出：false
 */

func isMatch(s, p string) bool {
	if s==p{
		return true
	}
	if len(p) == 0 && len(s) != 0 {
		return false
	}
	dp := make([][]bool, len(s) + 1)
	for i := 0; i <= len(s); i++ {
		dp[i] = make([]bool, len(p)+1)
	}
	dp[0][0] = true
	if len(s) != 0 && (p[0] == s[0] || p[0] == '.') {
		dp[1][1] = true
	}
	if p[0] == '*' {
		dp[0][1] = true
	}
	for i := 1; i < len(p); i++ {
		if p[i] == '*' && dp[0][i-1]{
			dp[0][i + 1] = true
		}
	}
	for i := 0; i < len(s); i++{
		for j := 1; j < len(p); j++ {
			if p[j] == '.' || p[j] == s[i] {
				dp[i + 1][j + 1] = dp[i][j]
			}
			if p[j] == '*' {
				if p[j - 1] != s[i] && p[j - 1] != '.' { //如果前一个元素不匹配 且不为任意元素
					dp[i + 1][j + 1] = dp[i + 1][j - 1]
				} else {
					dp[i + 1][j + 1] = (dp[i][j + 1] || dp[i + 1][j - 1]) // || dp[i + 1][j]不需要
				}
			}
		}
	}
	return dp[len(s)][len(p)]
}












