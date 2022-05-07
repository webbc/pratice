package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

type Config struct {
	version int32
}

func main() {
	var config atomic.Value

	go func() {
		for {
			config.Store(&Config{version: rand.Int31n(1000)})
			time.Sleep(time.Millisecond * 1)
		}
	}()

	var wg sync.WaitGroup

	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 1; j <= 100; j++ {
				c, ok := config.Load().(*Config)
				if !ok {
					continue
				}
				fmt.Println(c.version)
			}
		}()
	}

	wg.Wait()
}
