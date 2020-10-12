package easy

func isValid(s string) bool {
	m := map[byte]byte{')':'(', ']':'[', '}':'{'}
	stack := make([]byte, 0)
	if s == "" {
		return true
	}
	for i:=0;i<len(s);i++{
		if s[i] == '(' || s[i] == '[' || s[i] == '{' {
			stack = append(stack, s[i])
		}else if len(stack) > 0 && stack[len(stack)-1] == m[s[i]] {
			stack = stack[0:len(stack)-1]
		}else {
			return false
		}
	}
	return len(stack) == 0
}
