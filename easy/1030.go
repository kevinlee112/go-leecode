package easy

func allCellsDistOrder(R int, C int, r0 int, c0 int) [][]int {
	hash := make(map[int][][]int)
	res := make([][]int, R*C)
	maxDist := 0
	for i := 0; i < R; i++ {
		for j := 0; j < C; j++ {
			dist := distance(i, j, r0, c0)
			hash[dist] = append(hash[dist], []int{i, j})
			if dist > maxDist {
				maxDist = dist
			}
		}
	}

	k := 0
	for i := 0; i <= maxDist; i++ {
		for _, value := range hash[i] {
			res[k] = value
			k++
		}
	}
	return res
}

func distance(a int, b int, a1 int, b1 int) int {
	t1 := 0
	t2 := 0
	if a > a1 {
		t1 = a-a1
	} else {
		t1 = a1-a
	}
	if b > b1 {
		t2 = b-b1
	} else {
		t2 = b1-b
	}
	return t1 + t2
}
