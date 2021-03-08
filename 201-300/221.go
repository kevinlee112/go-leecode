package _01_300


func maximalSquare(matrix [][]byte) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}
	ans, d := 0, make([][]int, len(matrix))
	for i, m, n := 0, len(matrix), len(matrix[0]); i < m; i++ {
		d[i] = make([]int, n)
		for j := 0; j < n; j++ {
			if matrix[i][j] == '1' {
				if i == 0 || j == 0 {
					d[i][j] = 1
				} else {
					d[i][j] = min(d[i - 1][j], min(d[i][j - 1], d[i - 1][j - 1])) + 1
				}
			}
			if d[i][j] > ans {
				ans = d[i][j]
			}
		}
	}
	return ans * ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
