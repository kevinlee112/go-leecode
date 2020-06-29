package easy

import "strings"

func IsPrefixOfWord(sentence string, searchWord string) int {
	sentes := strings.Split(sentence, " ")

	for k, v := range sentes {
		if strings.HasPrefix(v, searchWord) {
			return k+1
		}
	}

	return -1
}