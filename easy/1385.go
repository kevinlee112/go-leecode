package easy

import "math"

func FindTheDistanceValue(arr1 []int, arr2 []int, d int) int {
	res := 0
	for _, x := range arr1{
		res++
		for _,y := range arr2 {
			if int(math.Abs(float64(x - y))) > d {
				res--
				break
			}
		}

	}

	return res
}