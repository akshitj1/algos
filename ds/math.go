package ds

import "math"

// Pow give x^q for ints
func Pow(x, q int) int64 {
	if q == 0 {
		return 1
	}
	y := Pow(x, q>>1)
	y *= y
	if q&1 == 1 {
		y *= int64(x)
	}
	return y
}

// PowMod gives (x^q)%mod
func PowMod(x, q, mod int) int {
	if q == 0 {
		return 1
	}

	y := int64(PowMod(x, q>>1, mod))
	y = (y * y) % int64(mod)
	if q&1 == 1 {
		y = (y * int64(x)) % int64(mod)
	}
	return int(y)
}

// Fib return nth fibonacci number using golden ratio
func Fib(n int) int64 {
	goldenRatio := (1 + math.Sqrt(5)) / 2
	nthFib := math.Pow(goldenRatio, float64(n)) / math.Sqrt(5)
	return int64(math.Round(nthFib))
}
