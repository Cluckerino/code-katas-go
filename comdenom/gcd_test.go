package comdenom_test

import (
	"testing"

	"github.com/cluckerino/code-katas-go/comdenom"
	"github.com/stretchr/testify/assert"
)

// Test Gcd
func TestGcd(t *testing.T) {
	tests := []struct {
		name           string
		a, b, expected int
	}{
		{"Fwd", 25, 40, 5},
		{"Rev", 40, 25, 5},
		{"Coprimes", 32, 35, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := comdenom.Gcd(tt.a, tt.b)
			assert.Equal(t, actual, tt.expected)
		})
	}
}
