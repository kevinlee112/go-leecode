package __100

import "fmt"

/**
给定两个大小分别为 m 和 n 的正序（从小到大）数组 nums1 和 nums2。请你找出并返回这两个正序数组的 中位数 。

示例 1：

输入：nums1 = [1,3], nums2 = [2]
输出：2.00000
解释：合并数组 = [1,2,3] ，中位数 2
示例 2：

输入：nums1 = [1,2], nums2 = [3,4]
输出：2.50000
解释：合并数组 = [1,2,3,4] ，中位数 (2 + 3) / 2 = 2.5
*/

func FndMedianSortedArrays(nums1 []int, nums2 []int) (result float64) {
	m, n := len(nums1), len(nums2)
	ls := []int{}
	var i, j int
	for i<m || j <n {
		if i<m && j <n {
			if nums1[i] < nums2[j] {
				ls = append(ls, nums1[i])
				i++
			}else {
				ls = append(ls, nums2[j])
				j++
			}
		}else if i < m {
			ls = append(ls, nums1[i])
			i++
		} else if j < n {
			ls = append(ls, nums2[j])
			j++
		}
	}
	fmt.Println(ls)
	if (m+n)%2 == 1 {
		return float64(ls[(m+n) / 2])
	}else {
		fmt.Println(ls[(m+n) / 2], ls[(m+n-2)/2])
		return float64(ls[(m+n) / 2] + ls[(m+n-2)/2])/2
	}
}
