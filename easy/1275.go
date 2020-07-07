package easy

func Tictactoe(moves [][]int) string {
	pan := [3][3]int{}
	l := len(moves)
	if l<5 {
		return "Pending"
	}
	for i:=l-1;i>=0;i-=2 {
		pan[moves[i][0]][moves[i][1]] = 1
	}
	if  (pan[0][1] + pan[0][2] + pan[0][0]) == 3 ||
		(pan[1][1] + pan[1][2] + pan[1][0]) ==3 ||
		(pan[2][1] + pan[2][2] + pan[2][0])==3 ||

		(pan[0][0] + pan[1][0] + pan[2][0]) == 3 ||
		(pan[0][1] + pan[1][1] + pan[2][1]) == 3 ||
		(pan[0][2] + pan[1][2] + pan[2][2]) == 3 ||

		(pan[0][0] + pan[1][1] + pan[2][2]) == 3 ||
		(pan[0][2] + pan[1][1] + pan[2][0]) == 3 {
		if l % 2 == 0 {
			return "B"
		}else {
			return "A"
		}
	}else {
		if l == 9 {
			return "Draw"
		}else {
			return "Pending"
		}
	}
}