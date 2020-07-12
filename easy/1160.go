package easy

import "fmt"

func CountCharacters(words []string, chars string) int {
	dict := make(map[byte]int)
	res := 0
	b := []byte(chars)
	fmt.Println(dict)
	for _, v := range words{
		for _,v := range b{
			dict[v]++
		}
		w := []byte(v)
		l := len(w)
		num := 0
		for _,k := range w{
			if dict[k] > 0 {
				num++
				dict[k]--
			}else {
				break
			}
		}
		if num == l {
			res += l
		}
		dict = make(map[byte]int)
	}
	return res
}