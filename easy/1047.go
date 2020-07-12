package easy

func RemoveDuplicates(S string) string {
	sBYte := []byte(S)
	for i := 0; i < len(sBYte)-1; i++ {
		for  i < len(sBYte)-1 && sBYte[i] == sBYte[i+1] {
			sBYte = append(sBYte[:i], sBYte[i+2:]...)
			if i>0 {
				i--
			}
		}
	}
	return string(sBYte)
}
