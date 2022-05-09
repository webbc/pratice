package main

import (
	"context"
	"fmt"
	"golang.org/x/sync/semaphore"
	"log"
	"runtime"
	"time"
)

func main() {
	ctx := context.TODO()

	var (
		maxWorkers = runtime.GOMAXPROCS(0)                    // 获取CPU核数作为worker数量
		sem        = semaphore.NewWeighted(int64(maxWorkers)) // 信号量
		out        = make([]int, 32)                          // 输出
	)

	for i := range out {

		// 如果没有worker可用，会阻塞在这里，直到某个worker被释放
		if err := sem.Acquire(ctx, 1); err != nil {
			log.Printf("Failed to acquire semaphore: %v", err)
			break
		}

		go func(i int) {
			defer sem.Release(1)

			time.Sleep(time.Second * 1)

			out[i] = i
		}(i)

	}

	// 请求所有的worker,这样能确保前面的worker都执行完
	if err := sem.Acquire(ctx, int64(maxWorkers)); err != nil {
		log.Printf("Failed to acquire semaphore: %v", err)
	}

	fmt.Println(out)
}
