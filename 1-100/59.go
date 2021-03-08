package __100
/**
给定一个正整数 n，生成一个包含 1 到 n2 所有元素，且元素按顺时针顺序螺旋排列的正方形矩阵。

示例:

输入: 3
输出:
[
 [ 1, 2, 3 ],
 [ 8, 9, 4 ],
 [ 7, 6, 5 ]
]
 */

func generateMatrix(n int) [][]int {
	matrix := make([][]int, n)
	for i:=0;i<n;i++ {
		matrix[i] = make([]int, n)
	}
	i , j := 0, 0
	top, right, bottom, left := 0, n-1, n-1, 0
	dir := 1
	count := 1
	for top <= bottom && left <= right {
		matrix[i][j] = count
		count++
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

	return matrix
}
