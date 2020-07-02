package easy

var (
	dict map[int][]int
	res int
)

func NumWays(n int, relation [][]int, k int) int {
	dict = make(map[int][]int,0)
	for i := 0 ; i < n ; i++{
		dict[i] = make([]int,0)
	}
	res = 0
	for _,v := range relation{
		dict[v[0]] = append(dict[v[0]], v[1])
	}
	dfs(0,k)
	return res
}

func dfs (curIndex,step int){
	if step == 0 {
		if curIndex == len(dict) -1{
			res ++
		}
		return
	}
	for _,next := range dict[curIndex] {
		temp := step - 1
		dfs(next,temp)
	}
}

