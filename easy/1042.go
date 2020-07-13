package easy

import "fmt"

func GardenNoAdj(N int, paths [][]int) []int {
	res := make([]int, 0)
	if len(paths) == 0 {
		for i:=0;i<N;i++{
			res = append(res, 1)
		}
		return res
	}
	dict := make(map[int]int)
	flag := true
	for flag {
		flag = false
		for i:=0;i<len(paths);i++ {
			if _, ok := dict[paths[i][0]];!ok{
				dict[paths[i][0]] = 1
			}
			if _, ok := dict[paths[i][1]];!ok{
				dict[paths[i][1]] = 2
			}
			if dict[paths[i][1]] == dict[paths[i][0]] {
				if dict[paths[i][1]] < 4 {
					dict[paths[i][1]]++
				}else {
					dict[paths[i][0]] = 1
				}
				flag = true
			}
		}
	}
	fmt.Println(dict)

	for i:=0;i<N;i++{
		if _, ok :=dict[i+1];ok {
			res = append(res, dict[i+1])
		}else {
			res = append(res, 1)
		}

	}
	return res
}