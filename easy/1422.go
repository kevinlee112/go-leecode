package easy

import "fmt"

func MaxScore(s string) int {
	sByte := []byte(s)
	lens := len(sByte)
	res, sum := 0, 0
	for i:=0;i<lens;i++ {
		if sByte[i] == '1' {
			sum ++
		}
	}

	fmt.Println(sum)
	left := 0
	for j:=0;j<lens-1;j++ {
		if sByte[j] == '0'{
			left ++
		}else {
			sum --
		}
		fmt.Println(left+sum)
		if res < (sum + left) {
			res = sum + left
		}
	}



	return res
}