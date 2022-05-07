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

	// 创建连接
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"47.100.15.39:12379", "47.100.15.39:22379", "47.100.15.39:32379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		log.Fatal(err)
	}

	// 创建session s1
	s1, err := concurrency.NewSession(cli, concurrency.WithTTL(120))
	if err != nil {
		log.Fatal(err)
	}

	// session.Close会Revoke当前session创建的lease
	defer s1.Close()

	fmt.Println("s1开始尝试加锁")

	// 创建锁
	m1 := concurrency.NewMutex(s1, "/lockkey")

	// 加锁
	if err := m1.Lock(context.TODO()); err != nil {
		log.Fatal(err)
	}

	fmt.Println("s1已得分布式锁")
	fmt.Println("s1正在进行业务处理...")
	time.Sleep(time.Second * 20)
	fmt.Println("s1业务处理完成")

	// 解锁
	if err := m1.Unlock(context.TODO()); err != nil {
		log.Fatal(err)
	}
}
