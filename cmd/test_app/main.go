package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("i am ok")
	for i := 1; i < 10000; i++ {
		time.Sleep(1000 * time.Millisecond)
		fmt.Printf("i = %d\n", i)
	}
}
