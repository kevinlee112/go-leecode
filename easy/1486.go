package easy

func XorOperation(n int, start int) int {
	ret := 0
	for i:=0;i<n;i++ {
		ret ^= start + n*2
	}

	return ret
}
