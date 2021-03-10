package __100

func numTrees(n int) int {
	dp :=make([]int,n+1)
	if n==0{
		return 0
	}
	dp[0]=1
	dp[1]=1        //寻找出边界
	for i:=2;i<=n;i++{
		for j:=0;j<=i-1;j++{
			dp[i]+=dp[j]*dp[i-1-j]   //递推关系式
		}
	}
	return dp[n]
}