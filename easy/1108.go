package easy


func defangIPaddr(address string) string {
	sByte := []byte(address)
	res := ""
	for _,v := range sByte{
		if v == '.' {
			res += "[.]"
		}else {
			res += string(v)
		}
	}

	return res
}