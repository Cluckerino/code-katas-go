package primesteps

// Given the bounds, find the smallest pair of primes that differ by the given gap/step.
// See: https://www.codewars.com/kata/steps-in-primes/go

// Step finds the smallest pair of prime (p, q) where p, q in [m, n],
// q > p, and q - p = g
// g: gap between the two primes
// m: inclusive lower limit of prime search
// n: inclusive upper limit of prime search
// returns: p, q in a slice, if found; otherwise, nil
func Step(g, m, n int) []int {
	// Check edge cases
	switch {
	case g == 1:
		// g==1 only acceptable for the first 2 primes.
		if m <= 2 && n >= 3 {
			return []int{2, 3}
		}
		fallthrough
	case g%2 == 1:
		fallthrough
	case m == n:
		return nil
	}
	// Start iterating on odd numbers:
	if m%2 == 0 {
		m = m + 1
	}
	// Store previously found primes here as we go up
	primeQueue := []int{}
	// Iterate through odd numbers
	for i := m; i <= n; i = i + 2 {
		// Skip non-primes
		if !IsPrime(i) {
			continue
		}
		keepIndex := 0
		// Iterate through past primes
		for _, p := range primeQueue {
			diff := i - p
			if diff == g { // diff == gap
				// Found the pair
				return []int{p, i}

			} else if diff > g { // diff too big
				// Increase index to ditch this prime
				keepIndex++
			}
		}
		// Ditch primes and append new one at the end
		primeQueue = append(primeQueue[keepIndex:], i)
	}
	return nil
}
