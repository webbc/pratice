package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	var wg sync.WaitGroup

	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go func(index int) {
			defer wg.Done()
			fmt.Printf("【goruntine#%d】开始工作\n", index)
			time.Sleep(time.Second * 2)
			fmt.Printf("【goruntine#%d】结束工作\n", index)
		}(i)
	}

	wg.Wait()
	fmt.Printf("所有goruntine全部结束，可以处理数据了")
}
