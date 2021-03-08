package __100

import "sort"

/**
给出一个无重叠的 ，按照区间起始端点排序的区间列表。

在列表中插入一个新的区间，你需要确保列表中的区间仍然有序且不重叠（如果有必要的话，可以合并区间）。

 

示例 1：

输入：intervals = [[1,3],[6,9]], newInterval = [2,5]
输出：[[1,5],[6,9]]
示例 2：

输入：intervals = [[1,2],[3,5],[6,7],[8,10],[12,16]], newInterval = [4,8]
输出：[[1,2],[3,10],[12,16]]
解释：这是因为新的区间 [4,8] 与 [3,5],[6,7],[8,10] 重叠。
 */
func insert(intervals [][]int, newInterval []int) [][]int {
	intervals = append(intervals, newInterval)
	sort.Slice(intervals, func(i,j int) bool{
		return intervals[i][0] < intervals[j][0]
	})
	var result [][]int
	for i:=0;i<len(intervals);i++ {
		if len(result) == 0 {
			result = append(result, intervals[i])
		}else if intervals[i][0] > result[len(result)-1][1] {
			result = append(result, intervals[i])
		}else if intervals[i][1] > result[len(result)-1][1] {
			result[len(result)-1][1] = intervals[i][1]
		}

	}
	return result
}





















