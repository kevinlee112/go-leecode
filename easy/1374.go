package easy


func GenerateTheString(n int) string {
	res := ""
	m := n
	if n%2 == 0 {
		res += "b"
		m = n -1
	}
	for i:=0;i<m;i++ {
		res += "a"
	}
	return res
}