package easy

func replaceElements(arr []int) []int {
	l := len(arr)
	res := make([]int, l)
	for i:=0;i<l;i++ {
		max := arr[i+1]
		for j:=i+2;j<l;j++ {
			if max < arr[j] {
				max = arr[j]
			}
		}
		res[i] = max
	}
	return res
}