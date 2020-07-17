package easy

func divisorGame(N int) bool {
	res := false
	for N > 1  {
		for i:=1;i<N;i++{
			if N%i == 0{
				res = !res
				N -= i
				break
			}
		}
	}
	return res
}