package easy

func MaxNumberOfBalloons(text string) int {
	b := []byte("balon")
	s := []byte(text)
	dict := make(map[byte]int)
	for _,v := range b{
		dict[v] = 0
	}
	for _,v := range s{
		if _, ok := dict[v]; ok {
			dict[v]++
		}
	}
	res := dict[b[0]]
	for i,v := range dict{
		if i == 'l' || i == 'o' {
			if v/2 < res {
				res = v/2
			}
		} else {
			if v < res {
				res = v
			}
		}
	}
	return res
}