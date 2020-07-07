package easy

func findSpecialInteger(arr []int) int {
	l := len(arr) / 4
	v := arr[0]
	for i:=1;i<len(arr);i++ {
		if v == arr[i] {
			l--
		}else{
			v = arr[i]
			l = len(arr) / 4
		}
		if l <0 {
			break
		}
	}
	return v
}
