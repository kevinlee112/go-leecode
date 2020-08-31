package easy

func minDeletionSize(A []string) int {
	var x,y = len(A),len(A[0])
	if x==1{
		return y
	}
	var res int
	for i:=0;i<y;i++{
		for j:=1;j<x;j++{
			if A[j][i]<A[j-1][i]{
				res++
				break
			}
		}
	}
	return res
}