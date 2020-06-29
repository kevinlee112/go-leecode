package easy

import "fmt"

func CanBeEqual(target []int, arr []int) bool {
	for _, v := range target{
		find := false
		for i:=0;i<len(arr);i++ {
			fmt.Println(arr)
			if arr[i] == v {
				arr = append(arr[:i], arr[i+1:]...)
				find = true
				break
			}
		}
		if !find {
			return false
		}
	}
	return true
}