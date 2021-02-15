package ds

// Primes lists primes till N
func Primes(n int) []int {
	prime := make([]bool, n+1)
	for x := 2; x <= n; x++ {
		prime[x] = true
	}
	primes := make([]int, 0)
	for x := 2; x <= n; x++ {
		if !prime[x] {
			continue
		}
		primes = append(primes, x)
		for y := x + x; y <= n; y += x {
			prime[y] = false
		}
	}
	return primes
}
