package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var a int32

	var wg sync.WaitGroup

	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 1; j <= 10000; j++ {
				atomic.AddInt32(&a, 1)
			}
		}()
	}
	wg.Wait()
	fmt.Println(a)
}
