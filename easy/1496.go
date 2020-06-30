package easy

func IsPathCrossing(path string) bool {
	type pos struct {
		x int
		y int
	}
	exist := map[pos]bool{{0,0}:true}
	x, y := 0,0
	pathByte := []byte(path)
	for _, v := range pathByte{
		if v == 'N' {
			y = y + 1
		} else if v == 'S' {
			y = y - 1
		} else if v == 'W' {
			x = x - 1
		} else if v == 'E' {
			x = x + 1
		}
		if exist[pos{x, y}] {
			return true
		}
		exist[pos{x, y}] = true
	}

	return false
}