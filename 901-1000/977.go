package _01_1000


/**
给你一个按 非递减顺序 排序的整数数组 nums，返回 每个数字的平方 组成的新数组，要求也按 非递减顺序 排序。
输入：nums = [-4,-1,0,3,10]
输出：[0,1,9,16,100]
解释：平方后，数组变为 [16,1,0,9,100]
排序后，数组变为 [0,1,9,16,100]
*/

func sortedSquares(A []int) []int {
	l := len(A)
	res := make([]int, l)
	for left, right:= 0, l-1; left <= right; {
		if A[left] < 0 && -A[left] > A[right] {
			res[right-left] = A[left] * A[left]
			left++
		}else {
			res[right-left] = A[right] * A[right]
			right--
		}
	}
	return res
}
