package kperm

import (
	"testing"
	"github.com/tj/assert"
)


func TestFactorial(t *testing.T) {
	t.Run("test samples", func(t *testing.T) {
    assert.Equal(t, Factorial(5), uint64(120))
		assert.Equal(t, Factorial(0), uint64(1))
	})
}
