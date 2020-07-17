package easy

func bitwiseComplement(N int) int {
	if N == 0 {
		return 1
	}
	res := 0
	for i:=1; N > 0; i*=2 {
		res += (N&1^1) * i
		N >>= 1
	}
	return res
}