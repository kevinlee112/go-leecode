package __100

/**
编辑距离

给你两个单词 word1 和 word2，请你计算出将 word1 转换成 word2 所使用的最少操作数 。

你可以对一个单词进行如下三种操作：

插入一个字符
删除一个字符
替换一个字符
 

示例 1：

输入：word1 = "horse", word2 = "ros"
输出：3
解释：
horse -> rorse (将 'h' 替换为 'r')
rorse -> rose (删除 'r')
rose -> ros (删除 'e')
示例 2：

输入：word1 = "intention", word2 = "execution"
输出：5
解释：
intention -> inention (删除 't')
inention -> enention (将 'i' 替换为 'e')
enention -> exention (将 'n' 替换为 'x')
exention -> exection (将 'n' 替换为 'c')
exection -> execution (插入 'u')
 */

func minDistance(word1 string, word2 string) int {
	var dp = make([][]int, len(word1)+1)
	s1 := len(word1)
	s2 := len(word2)
	for i := 0; i <= s1; i++ {
		dp[i] = make([]int, s2+1)
	}
	dp[0][0] = 0
	for i := 1; i <= s1; i++ {
		dp[i][0] = i
	}
	for i := 1; i <= s2; i++ {
		dp[0][i] = i
	}
	for i := 1; i <= s1; i++ {
		for j := 1; j <= s2; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]              // 如果 word1[i-1] 与 word2[i-1]相等
			} else {
				dp[i][j] = dp[i-1][j-1] + 1          // 替换代价
			}
			dp[i][j] = min(dp[i][j], dp[i][j-1] + 1) // 插入代价
			dp[i][j] = min(dp[i][j], dp[i-1][j] + 1) // 删除代价
		}
	}
	return dp[s1][s2]
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
