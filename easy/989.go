package easy

func addToArrayForm(A []int, K int) []int {
	var i = len(A)-1
	for ;i>=0 && K>0;i-- {
		A[i],K=(A[i]+K)%10,(A[i]+K)/10
	}

	for ;K!=0;K/=10{
		var res []int
		res=append(res,K%10)
		res=append(res,A...)
		A=res
	}
	return A
}
