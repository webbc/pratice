package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		time.Sleep(time.Second * 10)
	}()

	go func() {
		fmt.Println("g1开始Wait")
		wg.Wait()
		fmt.Println("g1结束Wait")
	}()

	fmt.Println("main开始Wait")
	wg.Wait()
	fmt.Println("main结束Wait")
	time.Sleep(time.Second * 2)
}
