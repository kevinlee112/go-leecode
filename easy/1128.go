package easy

func numEquivDominoPairs(dominoes [][]int) int {
	l := len(dominoes)
	dict := make(map[[2]int]int)
	res := 0
	for i := 0; i < l; i++ {
		if dict[[2]int{dominoes[i][0], dominoes[i][1]}] > 0 {
			res += dict[[2]int{dominoes[i][0], dominoes[i][1]}]
			dict[[2]int{dominoes[i][0], dominoes[i][1]}]++
		}else if dict[[2]int{dominoes[i][1], dominoes[i][0]}] > 0 {
			res += dict[[2]int{dominoes[i][1], dominoes[i][0]}]
			dict[[2]int{dominoes[i][1], dominoes[i][0]}]++
		}else {
			dict[[2]int{dominoes[i][0], dominoes[i][1]}]++
		}

	}

	return res
}