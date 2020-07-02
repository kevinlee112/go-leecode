package easy

func Reformat(s string) string {
	sByte := []byte(s)
	lens := len(sByte)
	res := make([]byte, 0)
	num, str := make([]byte, 0), make([]byte, 0)
	for i:=0;i<lens;i++{
		if sByte[i] > 58{
			num = append(num, sByte[i])
		}else {
			str = append(str, sByte[i])
		}
	}
	diff := len(num) - len(str)
	if diff > 0 && diff <2 {
		for i,v:=range str{
			res = append(res, v, num[i])
		}
		res = append([]byte{num[len(num) -1]}, res...)
	}else if diff <0 && diff > -2 {
		for i,v:=range num{
			res = append(res, v, str[i])
		}
		res = append([]byte{str[len(str) -1]}, res...)
	}else if diff == 0 {
		for i,v:=range num{
			res = append(res, v, str[i])
		}
	}

	return string(res)
}
