package easy


func findJudge(N int, trust [][]int) int {
	trs:=make([]int,N)
	usd:=make([]int,N)
	for _,v:=range trust{
		usd[v[0]-1]=1
		trs[v[1]-1]++
	}
	for k,v:=range trs{
		//trust[i] 是完全不同的，所以v==N-1可以作为判断法官的依据
		if v==N-1{
			//再来判断法官候选人是否不相信任何人
			if usd[k]==0{
				return k+1
			}else{
				return -1
			}
		}
	}
	return -1
}