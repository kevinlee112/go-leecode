package _01_300

func computeArea(A int, B int, C int, D int, E int, F int, G int, H int) int {
	w := min(C,G)-max(A,E)
	h := min(D,H)-max(B,F)
	return (C - A) * (D - B) + (G - E) * (H - F) - max(w,0) * max(h,0)
}

func min(x,y int) int{
	if x>y{
		return y
	}
	return x
}