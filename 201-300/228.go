package _01_300

import "strconv"

/**
给定一个无重复元素的有序整数数组 nums 。

返回 恰好覆盖数组中所有数字 的 最小有序 区间范围列表。也就是说，nums 的每个元素都恰好被某个区间范围所覆盖，并且不存在属于某个范围但不属于 nums 的数字 x 。

列表中的每个区间范围 [a,b] 应该按如下格式输出：

"a->b" ，如果 a != b
"a" ，如果 a == b
 

示例 1：

输入：nums = [0,1,2,4,5,7]
输出：["0->2","4->5","7"]
解释：区间范围是：
[0,2] --> "0->2"
[4,5] --> "4->5"
[7,7] --> "7"

 */
func summaryRanges(nums []int) []string {
	res := []string{}
	n := len(nums)
	if n == 0 {
		return res
	}
	l, r := 0, 0
	for r < n {
		for r < n-1 && nums[r+1] == nums[r] + 1{
			r++
		}
		if l == r {
			res = append(res, strconv.Itoa(l))
		}else if l < r {
			res = append(res, strconv.Itoa(nums[l]) + "->" + strconv.Itoa(nums[r]))
		}
		r++
		l = r
	}
	return res
}