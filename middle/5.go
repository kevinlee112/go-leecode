package middle

/**
给定一个字符串 s，找到 s 中最长的回文子串。你可以假设 s 的最大长度为 1000。

示例 1：

输入: "babad"
输出: "bab"
注意: "aba" 也是一个有效答案。
示例 2：

输入: "cbbd"
输出: "bb"
 */

func longestPalindrome(s string) string {
	if len(s) <= 1 {
		return s
	}
	dp := make([][]bool, len(s))
	start, end := 0, 0
	for i:=1;i<len(s);i++ {
		dp[i] = make([]bool, len(s))
		dp[i][i] = true
		for j:=0;j<i;j++ {
			if s[i] == s[j] && (dp[j+1][i-1] || i-j <=2) {
				dp[j][i] = true
				if i - j > end - start {
					start, end = j, i
				}
			}else {
				dp[j][i] = false
			}
		}
	}

	return s[start:end+1]
}






