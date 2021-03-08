package __100

/**
给定一个包含 m x n 个元素的矩阵（m 行, n 列），请按照顺时针螺旋顺序，返回矩阵中的所有元素。

示例 1:

输入:
[
 [ 1, 2, 3 ],
 [ 4, 5, 6 ],
 [ 7, 8, 9 ]
]
输出: [1,2,3,6,9,8,7,4,5]
示例 2:

输入:
[
  [1, 2, 3, 4],
  [5, 6, 7, 8],
  [9,10,11,12]
]
输出: [1,2,3,4,8,12,11,10,9,5,6,7]
 */
func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 {
		return []int{}
	}
	//边界
	top, right, bottom, left := 0, len(matrix[0])-1, len(matrix)-1, 0
	var res []int
	i, j := 0, 0
	dir := 1
	for top <= bottom && left <= right {
		res = append(res, matrix[i][j])
		switch dir {
		case 1:
			if j == right {
				dir=2
				top++
				i++
				continue
			}
			j++
		case 2:
			if i == bottom {
				dir = 3
				right--
				j--
				continue
			}
			i++
		case 3:
			if j==left {
				dir=4
				bottom--
				i--
				continue
			}
			j--
		case 4:
			if i== top {
				dir=1
				left++
				j++
				continue
			}
			i--
		}
	}
	return res
}






















