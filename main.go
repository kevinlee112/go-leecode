package main

import (
	"fmt"
	"summery/leecode/easy"
)

func main()  {
	n := [][]int{
		{1,0},
		{1,0},
		{10},
		{1,1,0,0,0},
		{1,1,1,1,1},
	}

	fmt.Println(easy.KWeakestRows(n, 3))

}
