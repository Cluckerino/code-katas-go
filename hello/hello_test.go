package hello_test

import (
	"testing"

	"github.com/cluckerino/code-katas-go/hello"
	"github.com/stretchr/testify/assert"
)

// Test Hello World
func TestHelloWorld(t *testing.T) {
	actual, expected := hello.World(), "Hello World!"
	assert.Equal(t, expected, actual)
}
