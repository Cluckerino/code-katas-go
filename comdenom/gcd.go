package comdenom

// Gcd gets the gcd of 2 numbers using the Euler algorithm
func Gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return Gcd(b, a%b)
}
