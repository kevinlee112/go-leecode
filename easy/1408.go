package easy

import "strings"

func StringMatching(words []string) []string {
	var res []string
	for i, n := range words{
		for j, m := range words{
			if i == j || len(m) <= len(n) {
				continue
			}
			if strings.Contains(m, n) {
				res = append(res, n)
				break
			}
		}
	}

	return res
}