package easy

func sortArrayByParityII(A []int) []int {
	res:= make([]int, len(A))
	i:=0
	for len(A) >0 {
		for j,v:=range A{
			if i%2 == v%2 {
				res[i]=v
				i++
				A = append(A[0:j], A[j+1:]...)
				break
			}
		}
	}
	return res
}
