package main

import (
	"context"
	"fmt"
	clientv3 "go.etcd.io/etcd/client/v3"
	"time"
)

func main() {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"47.100.15.39:12379", "47.100.15.39:22379", "47.100.15.39:32379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	resp, err := cli.Get(ctx, "/test", clientv3.WithPrefix())
	cancel()
	if err != nil {
		fmt.Printf("err:%v", err)
		return
	}

	fmt.Println(resp.Header.Revision)
	fmt.Println(string(resp.Kvs[0].Key), string(resp.Kvs[0].Value))

	defer cli.Close()
}
