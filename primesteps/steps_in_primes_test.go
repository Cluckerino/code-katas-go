package primesteps_test

import (
	"testing"

	"github.com/cluckerino/code-katas-go/primesteps"
	"github.com/stretchr/testify/assert"
)

type TestData struct {
	g, m, n  int
	expected []int
}

var kataExamples = []TestData{
	{2, 100, 110, []int{101, 103}},
	{4, 100, 110, []int{103, 107}},
	{6, 100, 110, []int{101, 107}},
	{8, 300, 400, []int{359, 367}},
	{10, 300, 400, []int{307, 317}},
}

var oneGapTests = []TestData{
	{1, 0, 4, []int{2, 3}},
	{1, 10, 30, []int(nil)},
}

func ExecuteTest(t *testing.T, data TestData) {
	actual, expected := primesteps.Step(data.g, data.m, data.n), data.expected
	assert.Equal(t, expected, actual)
}

func TestStepKataExamples(t *testing.T) {
	for _, td := range kataExamples {
		ExecuteTest(t, td)
	}
}

// Test odd gaps for range min > 2
func TestStepOddGap(t *testing.T) {
	actual, expected := primesteps.Step(11, 30000, 100000), []int(nil)
	assert.Equal(t, expected, actual)
}

// Test range min == range max
func TestStepZeroRange(t *testing.T) {
	actual, expected := primesteps.Step(10, 30000, 30000), []int(nil)
	assert.Equal(t, expected, actual)
}

// Test gap = 1
func TestStepGapIsOne(t *testing.T) {
	for _, td := range oneGapTests {
		ExecuteTest(t, td)
	}
}

// Test empty result
func TestStepValidInputsNoResult(t *testing.T) {
	actual, expected := primesteps.Step(10, 100, 112), []int(nil)
	assert.Equal(t, expected, actual)
}
