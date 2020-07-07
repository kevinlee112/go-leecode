package easy


func shiftGrid(grid [][]int, k int) [][]int {
	n := len(grid)
	m := len(grid[0])
	ret := make([][]int, n)
	//可以整行操作
	if k%m == 0 {
		for i := 0; i < n; i++ {
			//计算新矩阵行与旧矩阵的对应关系
			x := (i + k/m) % n
			ret[x] = grid[i]
		}
	} else {
		for i := 0; i < n; i++ {
			ret[i] = make([]int, m)
		}
		for i := 0; i < n; i++ {
			for j := 0; j < m; j++ {
				//计算新矩阵元素与旧矩阵元素对应关系
				y := (j + k) % m
				x := (i + (j+k)/m) % n
				ret[x][y] = grid[i][j]
			}
		}
	}
	return ret
}