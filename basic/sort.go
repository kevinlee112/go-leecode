package basic

//冒泡
func SortBubble(arr []int) []int {
	for i:=1;i<len(arr);i++ {
		for j:=0;j<len(arr)-i;j++ {
			if arr[j]>arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}

//选择
func sortSelect(arr []int) []int {
	for i:=0;i<len(arr);i++ {
		minIndex := i
		for j:=i+1;j<len(arr);j++ {
			if arr[i] > arr[j] {
				minIndex = j
			}
		}
		arr[i], arr[minIndex] = arr[minIndex], arr[i]
	}
	return arr
}

//快排
func sortQuick(arr []int, left int, right int)  {
	if left < right {
		i, j := left, right
		key := arr[(left + right)/2]
		for i<=j {
			for arr[i] < key {
				i++
			}
			for arr[j] > key {
				j--
			}
			if i <= j {
				arr[i], arr[j] = arr[j], arr[i]
				i++
				j--
			}
		}
		if left < j {
			sortQuick(arr, left, j)
		}
		if right > i {
			sortQuick(arr, i, right)
		}
	}
	
}
