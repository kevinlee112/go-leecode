package basic

func MergeSort(intList []int) []int {
	if len(intList) <=1 {
		return intList
	}
	mid := len(intList)/2
	left, right := MergeSort(intList[:mid]), MergeSort(intList[mid:])

	return merge(left, right)
}

func merge(left, right []int) []int {
	i,j:= 0, 0;
	temp := make([]int, 0)
	for i< len(left) && j<len(right){
		if left[i] < right[j] {
			temp = append(temp, left[i])
			i++
		}else {
			temp = append(temp, right[j])
			j++
		}
	}

	if i< len(left){
		temp = append(temp, left[i:]...)
	}else if j< len(right){
		temp = append(temp, right[i:]...)
	}
	return temp
}
