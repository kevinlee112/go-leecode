package _01_400

import "sort"

/**
给定一些标记了宽度和高度的信封，宽度和高度以整数对形式 (w, h) 出现。当另一个信封的宽度和高度都比这个信封大的时候，这个信封就可以放进另一个信封里，如同俄罗斯套娃一样。

请计算最多能有多少个信封能组成一组“俄罗斯套娃”信封（即可以把一个信封放到另一个信封里面）。

说明:
不允许旋转信封。

示例:

输入: envelopes = [[5,4],[6,4],[6,7],[2,3]]
输出: 3
解释: 最多信封的个数为 3, 组合为: [2,3] => [5,4] => [6,7]。
 */

func maxEnvelopes(envelopes [][]int) int {
	sort.Slice(envelopes, func(i, j int) bool {
		return envelopes[i][0] < envelopes[j][0] || (envelopes[i][0] == envelopes[j][0] && envelopes[i][1] > envelopes[j][1])
	})
	max := make([]int, 0)
	for i:=0;i<len(envelopes);i++ {
		if len(max) == 0 || envelopes[i][1] > max[len(max)-1] {
			max = append(max,envelopes[i][1])
		}else {
			for j := 0 ; j < len(max); j++ {
				if envelopes[i][1] <= max[j] {
					max[j] = envelopes[i][1]
					break
				}
			}
		}
	}

	return len(max)
}