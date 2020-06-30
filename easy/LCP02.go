package easy



func fraction(cont []int) []int {
	lens := len(cont)
	n,m := 0,1
	for i := lens-1; i >= 0; i--{
		n, m = m, m*cont[i]+n
	}
	return []int {m,n}
}