package kperm

import (
	"github.com/tj/assert"
	"reflect"
	"testing"
)

func TestPerms(t *testing.T) {
	assertDigitsCorrect := func(perm []int, n int) {
		// entries contain valid digits
		for _, v := range perm {
			assert.True(t, v >= 0)
			assert.True(t, v < n)
		}
		assert.True(t, len(perm) > 0)

		// digits never repeat
		for i := 0; i < len(perm); i++ {
			for j := i + 1; j < len(perm); j++ {
				assert.True(t, perm[i] != perm[j])
			}
		}
	}

	assertPermsNeverRepeat := func(perms [][]int) {
		for i := 0; i < len(perms); i++ {
			for j := i + 1; j < len(perms); j++ {
				assert.True(t, reflect.DeepEqual(perms[i], perms[j]))
			}
		}
	}

	assertValidPerm := func(n int, k int) {
		kp, err := New(n, k)
		assert.NoError(t, err, "init")

		currentPerm := kp.Perm()
		assertDigitsCorrect(currentPerm, n)
		perms := [][]int{currentPerm}
		assert.Equal(t, uint64(len(perms)-1), kp.Index())
		for !kp.Done() {
			kp.Next()
			currentPerm = kp.Perm()
			assertDigitsCorrect(currentPerm, n)
			perms = append(perms, currentPerm)
			assert.Equal(t, uint64(len(perms)-1), kp.Index())
		}

		// correct number of permutations
		assert.Equal(t, uint64(len(perms)), Factorial(n)/Factorial(n-k))

		// none of them repeat
		assertPermsNeverRepeat(perms)
	}

	assertInvalidPerm := func(n int, k int) {
		_, err := New(n, k)
		assert.Error(t, err, "init")
	}

	t.Run("sample n and k values", func(t *testing.T) {
		assertValidPerm(5, 3)
		assertValidPerm(5, 5)
		assertValidPerm(5, 1)
		assertValidPerm(8, 4)
	})

	t.Run("invalid n and k values", func(t *testing.T) {
		assertInvalidPerm(0, -1)
		assertInvalidPerm(5, 0)
		assertInvalidPerm(5, 6)
	})
}
