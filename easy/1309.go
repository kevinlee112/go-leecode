package easy

func freqAlphabets(s string) string {
	sByte := []byte(s)
	n := len(sByte)
	res := ""
	for i:=n-1;i>=0;i-- {
		if sByte[i] == '#' {
			res = switchChar(string(sByte[i-2]) + string(sByte[i-1])) + res
			i = i - 2

		}else {
			res = switchChar(string(sByte[i])) + res
		}
	}
	return res
}

func switchChar(char string) string {
	var res string

	switch char {
	case "1":
		res = "a"
	case "2":
		res = "b"
	case "3":
		res = "c"
	case "4":
		res = "d"
	case "5":
		res = "e"
	case "6":
		res = "f"
	case "7":
		res = "g"
	case "8":
		res = "h"
	case "9":
		res = "i"
	case "10":
		res = "j"
	case "11":
		res = "k"
	case "12":
		res = "l"
	case "13":
		res = "m"
	case "14":
		res = "n"
	case "15":
		res = "o"
	case "16":
		res = "p"
	case "17":
		res = "q"
	case "18":
		res = "r"
	case "19":
		res = "s"
	case "20":
		res = "t"
	case "21":
		res = "u"
	case "22":
		res = "v"
	case "23":
		res = "w"
	case "24":
		res = "x"
	case "25":
		res = "y"
	case "26":
		res = "z"
	}

	return res
}