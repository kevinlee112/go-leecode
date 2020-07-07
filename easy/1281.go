package easy

func subtractProductAndSum(n int) int {
	sum, qr := 0, 0
	for n > 0 {
		sum +=n%10
		qr *=n%10
		n/=10
	}
	return qr - sum
}
