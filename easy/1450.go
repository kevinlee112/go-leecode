package easy

func BusyStudent(startTime []int, endTime []int, queryTime int) int {
	ret := 0
	d := len(startTime)
	for i:=0;i<d;i++ {
		if startTime[i] <= queryTime && endTime[i] >= queryTime {
			ret++
		}
	}

	return ret
}
