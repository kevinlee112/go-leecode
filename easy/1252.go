package easy

func oddCells(n int, m int, indices [][]int) int {
	grip :=make([][]int,n)
	for i:=0;i<n;i++{
		grip[i] = make([]int,m)
	}
	for i:=0;i<len(indices);i++{
		r := indices[i][0]
		for j:=0;j<len(grip[r]);j++{
			grip[r][j]++
		}
		c := indices[i][1]
		for j:=0;j<len(grip);j++{
			grip[j][c]++
		}
	}
	var res int
	for i:=0;i<n;i++{
		for j:=0;j<m;j++{
			if grip[i][j]%2!=0{
				res++
			}
		}
	}
	return res
}
