package primesteps_test

import (
	"testing"

	"github.com/cluckerino/code-katas-go/primesteps"
	"github.com/stretchr/testify/assert"
)

var primes = []int{2, 3, 5, 7, 11, 857, 4231, 65323}
var nonPrimes = []int{-1, 0, 1, 4, 35, 799709}

func TestIsPrimeWithPrimes(t *testing.T) {
	for _, p := range primes {
		assert.True(t, primesteps.IsPrime(p))
	}
}

func TestIsPrimeWithNonPrimes(t *testing.T) {
	for _, np := range nonPrimes {
		assert.False(t, primesteps.IsPrime(np))
	}
}
