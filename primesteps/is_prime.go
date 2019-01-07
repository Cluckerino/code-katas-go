package primesteps

// IsPrime uses the 6k +/-1 algorithm when testing for primes.
// See: https://en.wikipedia.org/wiki/Primality_test#Simple_methods
func IsPrime(num int) bool {
	switch {
	// Numbers <= 1 aren't prime
	case num <= 1:
		return false
	// 2 and 3 are prime
	case num <= 3:
		return true
	// 6k method should eliminate multiples of 2/3
	case num%2 == 0 || num%3 == 0:
		return false
	}
	for i := 5; i*i <= num; i += 6 {
		if num%i == 0 || num%(i+2) == 0 {
			return false
		}
	}
	return true
}
