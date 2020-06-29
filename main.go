package main

import (
	"fmt"
	"summery/leecode/easy"
)

func main()  {
	//1431
	//candies := []int{2,3,5,1,3}
	//ret :=easy.KidsWithCandies(candies, 3)
	//
	//fmt.Println(ret)

	//1436
	//paths := [][]string{{"B","C"}, {"D","B"}, {"C","A"}}
	//easy.DestCity(paths)

	//target := []int{1, 3}
	//ret := easy.BuildArray(target, 3)
	//fmt.Println(ret)

	//str := "abbcccddddeeeeedcba"
	//pow := easy.MaxPower(str)
	//fmt.Println(pow)
	nums:= []int{1,2,3,4}
	ret := easy.Shuffle(nums, 2)
	fmt.Println(ret)
}
