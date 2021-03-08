package __100
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
//func permute(nums []int) [][]int {
//	if len(nums) == 1 {
//		return [][]int{nums}
//	}
//	var res [][]int
//	for i, num := range nums{
//		tmp := nums[0:i]
//		tmp = append(tmp, nums[i+1:]...)
//		sub := permute(tmp)
//		for _,s := range sub {
//			res = append(res, append(s, num))
//		}
//	}
//	return res
//}
//
//// 最终结果
//var result [][]int
//
//// 回溯核心
//// nums: 原始列表
//// pathNums: 路径上的数字
//// used: 是否访问过
//func backtrack(nums, pathNums []int, used[]bool) {
//	if len(nums) == len(pathNums) {
//		tmp := make([]int, len(nums))
//		copy(tmp, pathNums)
//		fmt.Println(pathNums)
//		fmt.Println(tmp)
//		result = append(result, tmp)
//		return
//	}
//
//	for i:=0; i<len(nums); i++ {
//		// 检查是否访问过
//		if !used[i] {
//			used[i] = true
//			pathNums = append(pathNums, nums[i])
//			backtrack(nums,pathNums,used)
//			pathNums = pathNums[:len(pathNums) -1]
//			used[i] = false
//		}
//	}
//}
//
//func Permute2(nums []int) [][]int {
//	var pathNums []int
//	var used = make([]bool, len(nums))
//	result = [][]int{}
//	backtrack(nums, pathNums, used)
//	return result
//}
//
//
//
//
//
//
//
