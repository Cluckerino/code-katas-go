package comdenom_test

import (
	"testing"

	"github.com/cluckerino/code-katas-go/comdenom"
	"github.com/stretchr/testify/assert"
)

func TestMultiLcm(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int
	}{
		{"", []int{6, 10, 12, 15}, 60},
		{"Coprimes", []int{3, 5, 7}, 105},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := comdenom.MultiLcm(tt.nums)
			assert.Equal(t, actual, tt.expected)
		})
	}
}