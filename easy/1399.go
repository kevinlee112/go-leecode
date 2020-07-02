package easy

func CountLargestGroup(n int) int {
	dict := make(map[int]int)
	max := 0
	res := 0
	for i:=1;i<n+1;i++ {
		dict[Sum(i)]++
		if max < dict[Sum(i)] {
			res = 1
			max = dict[Sum(i)]
		}else if max == dict[Sum(i)] {
			res++
		}
	}

	return res
}

func Sum(num int) int {
	res := 0
	res += num % 10
	for num/10  > 0 {
		num = num/10
		res += num % 10
	}
	return res
}