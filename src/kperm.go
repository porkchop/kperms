package kperm

import (
	"errors"
)


type KPerm struct {
	current []int
  index int
  k int
  n int
}

func New(n int, k int) (*KPerm, error) {
  if n < 1 {
    return nil, errors.New("n must be >= 1")
  }
  if k > n {
    return nil, errors.New("k must be <= n")
  }
  if k < 1 {
    return nil, errors.New("k must be >= 1")
  }

	firstperm := make([]int, n)
  for i := 0; i < n; i++ {
    firstperm[i] = i
  }
  kp := &KPerm{current: firstperm, index: 0, n: n, k: k}
	return kp, nil
}

func (kp KPerm) Perm() []int {
	return kp.current[0: kp.k]
}

func factorial(n int) int {
	i := 1
	for ;n > 0; n-- {
		i = i * n
	}
	return i
}

func (kp KPerm) Index() int {
  return kp.index
}

func (kp KPerm) NumPerms() int {
  return factorial(kp.n) / factorial(kp.n - kp.k)
}

func (kp KPerm) MaxIndex() int {
  return kp.NumPerms() - 1
}

func (kp KPerm) Done() bool {
  return kp.MaxIndex() == kp.index
}

func (kp KPerm) Reset() {
  for i := 0; i < kp.n; i++ {
    kp.current[i] = i
  }
  kp.index = 0
}

func swap(a []int, i int, j int) {
  a[i], a[j] = a[j], a[i]
}

func reverse(a []int, i int, j int) {
	r := a[i:j+1]
	for i, j = 0, j - i; i < j; i, j = i + 1, j - 1 {
		r[i], r[j] = r[j], r[i]
	}
}

func (kp *KPerm) Next() []int {
  a := kp.current
  edge := kp.k - 1
  n := kp.n
  k := kp.k

  // find j in (k…n-1) where aj > aedge
  j := k
  for j < n && a[edge] >= a[j] {
    j++
  }

  if j < n {
    swap(a, edge, j)
  } else {
    reverse(a, k, n-1)

    // find rightmost ascent to left of edge
    i := edge - 1
    for i >= 0 && a[i] >= a[i+1] {
      i--
    }

    if i < 0 {
      // no more permutations
      return nil
    }

    // find j in (n-1…i+1) where aj > ai
    j = n - 1
    for j > i && a[i] >= a[j] {
      j--
    }

    swap(a, i, j)
    reverse(a, i+1, n-1)
  }

  kp.index++

  return a[0:k-1]
}
