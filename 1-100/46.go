package __100

import "sort"

//
//import "fmt"
//
///**
//46. 全排列
//给定一个 没有重复 数字的序列，返回其所有可能的全排列。
//
//示例:
//
//输入: [1,2,3]
//输出:
//[
//  [1,2,3],
//  [1,3,2],
//  [2,1,3],
//  [2,3,1],
//  [3,1,2],
//  [3,2,1]
//]
// */


func permute(nums []int) (ans [][]int) {
	sort.Ints(nums)
	var pathNums []int
	n := len(nums)
	used := make([]bool, n)
	var backtrach func(int)
	backtrach = func(idx int) {
		if idx == n{
			ans = append(ans, append( []int{}, pathNums...))
			return
		}
		for i:=0;i<n;i++{
			if !used[i] {
				used[i] = true
				pathNums = append(pathNums, nums[i])
				backtrach(idx+1)
				pathNums = pathNums[:len(pathNums) -1]
				used[i] = false
			}
		}
	}
	backtrach(0)
	return
}





