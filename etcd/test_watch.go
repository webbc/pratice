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

	ch := cli.Watch(context.Background(), "/test", clientv3.WithRev(1))
	for {
		select {
		case watchResp := <-ch:
			for _, event := range watchResp.Events {
				fmt.Println(string(event.Kv.Key), string(event.Kv.Value), event.Kv.ModRevision)
			}
		}
	}
}
