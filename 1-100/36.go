package __100

import "math"

/**
判断一个 9x9 的数独是否有效。只需要根据以下规则，验证已经填入的数字是否有效即可。

数字 1-9 在每一行只能出现一次。
数字 1-9 在每一列只能出现一次。
数字 1-9 在每一个以粗实线分隔的 3x3 宫内只能出现一次。


上图是一个部分填充的有效的数独。

数独部分空格内已填入了数字，空白格用 '.' 表示。

示例 1:

输入:
[
  ["5","3",".",".","7",".",".",".","."],
  ["6",".",".","1","9","5",".",".","."],
  [".","9","8",".",".",".",".","6","."],
  ["8",".",".",".","6",".",".",".","3"],
  ["4",".",".","8",".","3",".",".","1"],
  ["7",".",".",".","2",".",".",".","6"],
  [".","6",".",".",".",".","2","8","."],
  [".",".",".","4","1","9",".",".","5"],
  [".",".",".",".","8",".",".","7","9"]
]
输出: true
示例 2:

输入:
[
  ["8","3",".",".","7",".",".",".","."],
  ["6",".",".","1","9","5",".",".","."],
  [".","9","8",".",".",".",".","6","."],
  ["8",".",".",".","6",".",".",".","3"],
  ["4",".",".","8",".","3",".",".","1"],
  ["7",".",".",".","2",".",".",".","6"],
  [".","6",".",".",".",".","2","8","."],
  [".",".",".","4","1","9",".",".","5"],
  [".",".",".",".","8",".",".","7","9"]
]
输出: false
解释: 除了第一行的第一个数字从 5 改为 8 以外，空格内其他数字均与 示例1 相同。
     但由于位于左上角的 3x3 宫内有两个 8 存在, 因此这个数独是无效的。
*/

func isValidSudoku(board [][]byte) bool {
	var hang [9]int       //行
	var lie [9]int        //列
	var hanglie [3][3]int //9个九宫格
	for index1, val1 := range board {
		for index2, _ := range val1 {
			if board[index1][index2] == '.' { //如果是.直接跳出
				continue
			}
			//用与运算判断行中是否出现某该元素，若出现直接返回false
			if hang[index1]&int(math.Pow(2, float64(board[index1][index2]-'0'))) == int(math.Pow(2, float64(board[index1][index2]-'0'))) {
				return false
			} else {
				//用或运算将该元素对应的位填充为1
				hang[index1] = hang[index1] | int(math.Pow(2, float64(board[index1][index2]-'0')))
			}
			//同上
			if lie[index2]&int(math.Pow(2, float64(board[index1][index2]-'0'))) == int(math.Pow(2, float64(board[index1][index2]-'0'))) {
				return false
			} else {
				lie[index2] = lie[index2] | int(math.Pow(2, float64(board[index1][index2]-'0')))
			}
			//同上
			if hanglie[index1/3][index2/3]&int(math.Pow(2, float64(board[index1][index2]-'0'))) == int(math.Pow(2, float64(board[index1][index2]-'0'))) {
				return false
			} else {
				hanglie[index1/3][index2/3] = hanglie[index1/3][index2/3] | int(math.Pow(2, float64(board[index1][index2]-'0')))
			}
		}
	}
	return true
}
