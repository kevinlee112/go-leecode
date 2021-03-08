package _01_500

/**
494. 目标和
给定一个非负整数数组，a1, a2, ..., an, 和一个目标数，S。现在你有两个符号 + 和 -。对于数组中的任意一个整数，你都可以从 + 或 -中选择一个符号添加在前面。

返回可以使最终数组和为目标数 S 的所有添加符号的方法数。
示例：
输入：nums: [1, 1, 1, 1, 1], S: 3
输出：5
解释：

-1+1+1+1+1 = 3
+1-1+1+1+1 = 3
+1+1-1+1+1 = 3
+1+1+1-1+1 = 3
+1+1+1+1-1 = 3

一共有5种方法让最终目标和为3。
 */

func findTargetSumWays(nums []int, S int) int {
	sum := 0
	for e :=range nums{
		sum = sum + nums[e]
	}
	if S>sum{
		return 0
	}
	if (sum+S)&1!=0{
		return 0
	}
	sumAdd:=(sum+S)/2
	dp:=make([]int,sumAdd+1)
	dp[0] = 1
	for _, num := range nums{
		for i:=sumAdd;i>=num;i--{
			dp[i]+=dp[i-num]
		}
	}
	return dp[sumAdd]
}