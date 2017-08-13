package kperm

const LIM = 41

var facts [LIM]uint64

func Factorial(n int) (res uint64) {
	if facts[n] != 0 {
		res = facts[n]
		return res
	}

	if n > 0 {
		res = uint64(n) * Factorial(n-1)
		return res
	}

	return 1
}
