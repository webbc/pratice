package main

import (
	"fmt"
	"time"
)

func main() {

	go func() {
		var a, b chan int

		select {
		case a <- 1:
		case b <- 1:
		}

		fmt.Println("22222")
	}()

	time.Sleep(time.Second * 5)
	fmt.Println(1111)
}
