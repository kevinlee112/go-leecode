package easy

import "sort"

func DistanceBetweenBusStops(distance []int, start int, destination int) int {
	l := len(distance)
	sum1 := 0
	sum2 := 0
	s := []int{start, destination}
	sort.Ints(s)
	for i:=0; i< l; i++ {
		if i>=s[0] && i < s[1] {
			sum1 += distance[i]
		} else {
			sum2 += distance[i]
		}
	}
	if sum1 > sum2 {
		return sum2
	}else {
		return sum1
	}
}