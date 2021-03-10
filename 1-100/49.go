package __100

import "sort"

func groupAnagrams(strs []string) [][]string {
	res := [][]string{}
	m := map[string][]string{}
	for _, v := range strs {
		bytes := []byte(v)
		sort.Slice(bytes, func(i, j int) bool {return bytes[i] < bytes[j]})
		m[string(bytes)] = append(m[string(bytes)], v)
	}
	for _,v := range m {
		res = append(res, v)
	}
	return res
}
