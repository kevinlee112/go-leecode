package basic


/**
描述  二维数组0 1 判断从左上角 走到 右下角 是否能走到
返回 true false
 */

func checkThrow(a [][]int) bool {
	if len(a) == 0 {
		return true
	}
	m, n := len(a), len(a[0])
	if n == 0 {
		return true
	}
	dp := make([][]int, m)
	for i:=0;i<m;i++ {
		if a[i][0] == 1{
			dp[i][0] = 1
		}else{
			dp[i][0] = 0
		}
	}
	for i:=0;i<n;i++ {
		if a[0][i] == 1{
			dp[0][i] = 1
		}else{
			dp[0][i] = 0
		}
	}
	for i:=1;i<m;i++ {
		for j:=1;j<n;j++ {
			dp[i][j] = dp[i-1][j] + dp[i][j-1]
		}
	}

	return dp[m-1][n-1] > 0
}
