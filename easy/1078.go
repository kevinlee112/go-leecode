package easy

import "strings"

func findOcurrences(text string, first string, second string) []string {
	sSlice := strings.Split(text, " ")
	l := len(sSlice)
	res := make([]string, 0)
	if l<3 {
		return res
	}
	for i:=2;i<len(sSlice);i++{
		if sSlice[i-1]== second && sSlice[i-2] == first {
			res = append(res, sSlice[i])
		}
	}
	return res
}