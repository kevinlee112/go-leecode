package basic

import (
	"fmt"
	"sync"
)

func QuickSort(data []int) []int {
	if len(data) <= 1 {
		return data
	}
	var wg sync.WaitGroup
	c := data[0]
	var s1, s2 []int

	for k, v := range data {
		if k == 0 {
			continue
		}
		if c > v {
			s2 = append(s2, v)
		} else {
			s1 = append(s1, v)
		}
	}

	wg.Add(2)
	go func() {
		s1 = QuickSort(s1)
		wg.Done()
	}()
	go func() {
		s2 = QuickSort(s2)
		wg.Done()
	}()
	wg.Wait()

	data = []int{c}
	if len(s1) > 0 {
		data = append(s1, data...)
	}
	if len(s2) > 0 {
		data = append(data, s2...)
	}
	return data
}

func main() {
	data := []int{3, 6, 23, 7, 2, 4, 9, 13}
	fmt.Println(QuickSort(data))
}

