package easy

func balancedStringSplit(s string) int {
	res := 0
	sByte := []byte(s)
	count := 1
	value := sByte[0]
	for i:=1;i<len(sByte);i++ {
		if count == 0 {
			value = sByte[i]
			count++
			continue
		}
		if sByte[i] != value {
			count--
			if count == 0 {
				res++
			}
		}else {
			count++
		}
	}
	return res
}
