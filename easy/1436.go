package easy

func DestCity(paths [][]string) string {
	dst := paths[0][1]
	for len(paths) > 0 {
		n := len(paths)
		for i:=1;i<n;i++ {
			if dst == paths[i][0] {
				dst = paths[i][1]
				paths = append(paths[:i], paths[i+1:]...)
			}
		}
	}

	return dst
}