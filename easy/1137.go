package easy

func Tribonacci(n int) int {
	if n < 2{
		return n
	}
	if n == 2{
		return 1
	}
	dp := make([]int, n + 1)
	dp[0] = 0
	dp[1] = 1
	dp[2] = 1
	for i:=3; i<=n; i++ {
		dp[i] = dp[i-1] + dp[i-2] + dp[i-3]
	}
	return dp[n]
}