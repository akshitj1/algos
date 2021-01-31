package ds

// Pow give x^q for ints
func Pow(x, q int) int64 {
	if q == 0 || x == 0 {
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
