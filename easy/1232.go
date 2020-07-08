package easy


func checkStraightLine(coordinates [][]int) bool {
	l := len(coordinates)
	diff := []int{coordinates[1][0] - coordinates[0][0], coordinates[1][1] - coordinates[0][1]}
	for i:=2;i<l;i++ {
		if (coordinates[i][0] - coordinates[0][0]) * diff[1] != (coordinates[i][1] - coordinates[0][1]) * diff[0] {
			return false
		}
	}
	return true
}