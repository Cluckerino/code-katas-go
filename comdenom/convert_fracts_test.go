package comdenom_test

import (
	"testing"

	"github.com/cluckerino/code-katas-go/comdenom"
	"github.com/stretchr/testify/assert"
)

func RunTest(t *testing.T, input [][]int, expected string) {
	actual := comdenom.ConvertFracts(input)
	assert.Equal(t, expected, actual)
}

// Easy example from kata
func TestConvertFractsKataEx1(t *testing.T) {
	RunTest(t,
		[][]int{{1, 2}, {1, 3}, {1, 4}},
		"(6,12)(4,12)(3,12)",
	)
}

// Harder example from kata
func TestConvertFractsKataEx2(t *testing.T) {
	RunTest(t,
		[][]int{{69, 130}, {87, 1310}, {30, 40}},
		"(18078,34060)(2262,34060)(25545,34060)",
	)
}
