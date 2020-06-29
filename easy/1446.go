package easy

func MaxPower(s string) int {
	ret, max := 1, 1
	str := []byte(s)
	lens := len(str)
	for i := 1; i < lens; i++ {
		if str[i] == str[i-1] {
			max ++
		}else {
			max =1
		}
		if max > ret {
			ret = max
		}
	}
	return ret
}