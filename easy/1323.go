package easy

import "strconv"

func maximum69Number (num int) int {
	s := strconv.Itoa(num)
	sByte := []byte(s)
	for i, v := range sByte{
		if v == '6' {
			sByte[i] = '9'
			break
		}
	}
	res, _ := strconv.Atoi(string(sByte))
	return res
}