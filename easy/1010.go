package easy

func numPairsDivisibleBy60(time []int) int {
	res := 0
	dict := make(map[int]int)
	for i:=0;i<len(time);i++{
		if dict[60 - time[i] % 60] > 0 {
			res += dict[60 - time[i] % 60]
		}
		dict[time[i] % 60]++
	}
	return res
}