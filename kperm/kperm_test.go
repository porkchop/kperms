package kperm

import (
	"fmt"
	"testing"
)

func TestThreePermFive(t *testing.T) {
	kp, err := New(5, 3)
	if err != nil {
		fmt.Println(err)
		return
	}

	for !kp.Done() {
		fmt.Println(kp.Perm())
		kp.Next()
	}
	fmt.Println("Done", kp.MaxIndex(), kp.Index())
}
