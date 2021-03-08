package basic

/**
存在数组，例如 A [1,2,3,3,4,4,4,5,5,5,6,6,6,11,13,21,22,24,24,24,33]
这样有序且存在重复元素的数组，写出求该数组中某个元素第一次出现的序号。例如 A 数组中的元素 4 第一次出现的序号为 4 （序号从 0 开始记）

1
21 ； 12
123 132 213 312 321 231


 */

func findBinary(arr []int, a int, left int, right int) int {
	if left > right {
		 return -1
	}
	mid := (left + right)/2
	if arr[mid] == a && ( mid ==0 || arr[mid-1] != a ) {
		return mid
	}else if arr[mid] > a {
		return findBinary(arr, a, mid + 1, right)
	}else{
		return findBinary(arr, a, left, mid-1)
	}
}

func valid(arr []string)  {

}



















