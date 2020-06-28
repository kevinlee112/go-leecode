package main

import (
	"fmt"
	"summery/leecode/easy"
)

func main()  {
	candies := []int{2,3,5,1,3}
	ret :=easy.KidsWithCandies(candies, 3)

	fmt.Println(ret)
}
