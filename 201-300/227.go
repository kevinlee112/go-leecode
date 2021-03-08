package _01_300


func calculate(s string) int {
	stack := make([]int, 0, len(s))
	op := make([]int, 0, len(s))
	num := 0
	for i:=0;i<len(s);i++ {
		switch s[i] {
		case '1','2','3','4','5','6','7','8','9','0':
			num = 0
			for i < len(s) && s[i]>='0'&&s[i]<='9'{
				num = num*10 + int(s[i]-'0')
				i++
			}
			if len(op) > 0 && op[len(op) - 1] > 2 {
				if op[len(op)-1] == 3 {
					stack[len(stack)-1] *= num
				}else {
					stack[len(stack)-1] /= num
				}
				op = op[:len(op)-1]
			}else {
				stack = append(stack, num)
			}
			i--
		case '+':
			op = append(op, 1)
		case '-':
			op = append(op, -1)
		case '*':
			op = append(op, 3)
		case '/':
			op = append(op, 4)
		}
	}
	for len(op) > 0 {
		stack[1] = stack[0] + op[0]*stack[1]
		op = op[1:]
		stack = stack[1:]
	}

	return stack[0]
}