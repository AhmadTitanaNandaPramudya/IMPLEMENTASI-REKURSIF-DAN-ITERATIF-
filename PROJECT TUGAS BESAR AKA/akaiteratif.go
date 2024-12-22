func faktorialIteratif(n int) int {
	hasil := 1
	for i := 1; i <= n; i++ {
		hasil *= i
	}
	return hasil
}

func kombinasiIteratif(n, r int) int {
	return faktorialIteratif(n) / (faktorialIteratif(r) * faktorialIteratif(n-r))
}