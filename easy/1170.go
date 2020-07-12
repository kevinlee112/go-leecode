package easy

import "fmt"

func NumSmallerByFrequency(queries []string, words []string) []int {
	num := make([]int, len(words))
	for i,v := range words{
		num[i] = f(v)
	}
	fmt.Println(num)
	res := make([]int, len(queries))
	for i, v := range queries{
		for _, k := range num{
			if f(v) < k {
				res[i]++
			}
		}
	}

	return res
}

func f(s string) int {
	b := []byte(s)
	min := b[0]
	res := 1
	for i:=1;i<len(b);i++{
		if b[i] == min {
			res++
		}else if b[i] < min {
			res = 1
			min = b[i]
		}
	}
	fmt.Println(res)
	return res
}
