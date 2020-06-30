package easy

func Average(salary []int) float64 {
	min,max,sum:=salary[0],salary[0],0
	for _,i:=range salary{
		sum+=i
		if i<min{
			min=i
		}
		if i>max{
			max=i
		}
	}
	return float64(sum-min-max)/float64( len(salary)-2)
}