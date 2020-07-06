package easy

func checkIfExist(arr []int) bool {
	lens := len(arr)
	for i:=0;i<lens;i++ {
		for j:=i+1;j<lens;j++ {
			if arr[i] == arr[j] * 2 || arr[j] == arr[i] * 2 {
				return true
			}
		}
	}
	return false
}
