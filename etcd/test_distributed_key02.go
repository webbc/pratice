package main

import (
	"context"
	"fmt"
	clientv3 "go.etcd.io/etcd/client/v3"
	concurrency "go.etcd.io/etcd/client/v3/concurrency"
	"log"

	"time"
)

func main() {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"47.100.15.39:12379", "47.100.15.39:22379", "47.100.15.39:32379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		log.Fatal(err)
	}

	// 创建session
	s2, err := concurrency.NewSession(cli, concurrency.WithTTL(120))
	if err != nil {
		log.Fatal(err)
	}
	defer s2.Close()

	fmt.Println("s2开始尝试加锁")

	m2 := concurrency.NewMutex(s2, "/lockkey")
	if err := m2.Lock(context.TODO()); err != nil {
		log.Fatal(err)
	}

	fmt.Println("s2已得分布式锁")
	fmt.Println("s2正在进行业务处理...")
	time.Sleep(time.Second * 20)
	fmt.Println("s2业务处理完成")

	if err := m2.Unlock(context.TODO()); err != nil {
		log.Fatal(err)
	}
}
