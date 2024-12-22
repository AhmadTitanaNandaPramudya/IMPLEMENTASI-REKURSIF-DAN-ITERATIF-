func faktorial(n int) int {
	if n <= 1 {
		return 1
	}
	return n * faktorial(n-1)
}

func kombinasi(n, r int) int {
	return faktorial(n) / (faktorial(r) * faktorial(n-r))
}