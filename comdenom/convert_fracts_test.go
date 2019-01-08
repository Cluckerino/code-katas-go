package comdenom_test

import (
	"testing"

	"github.com/cluckerino/code-katas-go/comdenom"
	"github.com/stretchr/testify/assert"
)

// Tests for convert fracts
func TestConvertFracts(t *testing.T) {
	tests := []struct {
		name     string
		input    [][]int
		expected string
	}{
		{
			"KataEx1",
			[][]int{{1, 2}, {1, 3}, {1, 4}},
			"(6,12)(4,12)(3,12)",
		},
		{
			"KataEx2",
			[][]int{{69, 130}, {87, 1310}, {30, 40}},
			"(18078,34060)(2262,34060)(25545,34060)",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := comdenom.ConvertFracts(tt.input)
			assert.Equal(t, tt.expected, actual)
		})
	}
}
