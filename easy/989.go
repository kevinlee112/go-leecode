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
		//    A=append(A,K%10,A...)  语法错误，不能A...不能与其他的一起
		A=res
	}
	return A
}
