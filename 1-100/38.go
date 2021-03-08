package __100

import "strconv"

func countAndSay(n int) string {
	if n==1 {
		return "1"
	}

	last := countAndSay(n-1)
	first := last[0]
	num := 1
	str := ""
	for i:=1;i<len(last);i++ {
		if last[i] == first {
			//相等,计数就可以
			num++
		}else {
			str += strconv.Itoa(num)
			str += string(first)
			first = last[i]
			num = 1
		}
	}
	str += strconv.Itoa(num)
	str += string(first)

	return str
}


