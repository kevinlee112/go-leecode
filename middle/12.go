package middle

func intToRoman(num int) string {
	convertMap := map[int]string{
		1:    "I",
		4:    "IV",
		5:    "V",
		9:    "IX",
		10:   "X",
		40:   "XL",
		50:   "L",
		90:   "XC",
		100:  "C",
		400:  "CD",
		500:  "D",
		900:  "CM",
		1000: "M",
	}
	convertList := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	resultStr := ""
	for i := 0; i < len(convertList) && num != 0; i++ {
		for count := (num / convertList[i]);count != 0;count--{
			resultStr = resultStr + convertMap[convertList[i]]
		}
		num = num%convertList[i]
	}
	return resultStr
}
