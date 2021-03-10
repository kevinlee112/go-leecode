package _01_400

/**
给定一个正整数 n ，你可以做如下操作：

如果 n 是偶数，则用 n / 2替换 n 。
如果 n 是奇数，则可以用 n + 1或n - 1替换 n 。
n 变为 1 所需的最小替换次数是多少？

 

示例 1：

输入：n = 8
输出：3
解释：8 -> 4 -> 2 -> 1
示例 2：

输入：n = 7
输出：4
解释：7 -> 8 -> 4 -> 2 -> 1
或 7 -> 6 -> 3 -> 2 -> 1
示例 3：

输入：n = 4
输出：2
 */


func integerReplacement(n int) int {
	return dp(n)
}

func dp(n int) int {
	if n == 1 {
		return 0
	}
	if n%2 == 0 {
		return dp(n/2) + 1
	} else {
		return min(dp(n-1), dp(n+1)) + 1
	}
}
