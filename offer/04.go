package offer


func findNumberIn2DArray(matrix [][]int, target int) bool {
	m, n := len(matrix), len(matrix[0])
	if m==0 || n==0 {
		return false
	}
	start, end := 0, 0
	for i:=0;i<min(m, n);i++ {
		if matrix[i][i] < target {
			start = i
		}else if matrix[i][i] > target {
			end = i
			break
		}else {
			return true
		}
	}

	if start == m-1 {
		for j:= 0;j<n;j++ {
			if matrix[m-1][j] == target{
				return true
			}
		}

	}else if start == n {
		for i:= 0;i<m;i++ {
			if matrix[i][n-1] == target{
				return true
			}
		}
	}else {

	}

	for i:=start+1;i<end;i++ {
		for j:=start+1;j<end;i++ {
			if matrix[i][j] == target {
				return true
			}
		}
	}

	return false
}

func min(i, j int) int {
	if i > j {
		return j
	}
	return i
}




