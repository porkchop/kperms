package main

import (
	"./src"
	"fmt"
)

func main() {
	work()
}

func work(){
	kp, err := kperm.New(5, 3)
	if err!=nil{
		fmt.Println(err)
		return
	}

	for !kp.Done() {
		fmt.Println(kp.Perm())
		kp.Next()
	}
	fmt.Println("Done", kp.MaxIndex(), kp.Index())
}
