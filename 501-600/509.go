package _01_600

/**
斐波那契数
 */

func fib(N int) int {
	if N == 0 || N == 1 {
		return N
	}
	return fib(N-1) + fib(N-2)
}