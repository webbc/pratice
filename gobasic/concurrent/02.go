package main

import (
	"container/list"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Pool struct {
	lock    sync.Mutex // 锁
	clients list.List  // 连接
	cond    *sync.Cond // cond实例
	close   bool       // 是否关闭
}

type Client struct {
	id int32
}

func NewClient() *Client {
	return &Client{
		id: rand.Int31n(100000),
	}
}

func (this *Client) Close() {
	fmt.Printf("Client:%d 正在关闭", this.id)
}

func NewPool(maxConnNum int) *Pool {
	pool := new(Pool)
	pool.cond = sync.NewCond(&pool.lock)

	// 创建连接
	for i := 0; i < maxConnNum; i++ {
		client := NewClient()
		pool.clients.PushBack(client)
	}

	return pool
}

func (this *Pool) Pull() *Client {
	this.lock.Lock()
	defer this.lock.Unlock()

	// 已关闭
	if this.close {
		fmt.Println("Pool is closed")
		return nil
	}

	// 如果连接池没有连接 需要阻塞
	for this.clients.Len() <= 0 {
		this.cond.Wait()
	}

	// 从链表中取出头节点，删除并返回
	ele := this.clients.Remove(this.clients.Front())
	return ele.(*Client)
}

func (this *Pool) Push(client *Client) {
	this.lock.Lock()
	defer this.lock.Unlock()

	if this.close {
		fmt.Println("Pool is closed")
		return
	}

	// 向链表尾部插入一个连接
	this.clients.PushBack(client)

	// 唤醒一个正在等待的goruntine
	this.cond.Signal()
}

func (this *Pool) Close() {
	this.lock.Lock()
	defer this.lock.Unlock()

	// 关闭连接
	for e := this.clients.Front(); e != nil; e = e.Next() {
		client := e.Value.(*Client)
		client.Close()
	}

	// 重置数据
	this.close = true
	this.clients.Init()
}

func main() {

	var wg sync.WaitGroup

	pool := NewPool(3)
	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go func(index int) {

			defer wg.Done()

			// 获取一个连接
			client := pool.Pull()

			fmt.Printf("Time:%s | 【goruntine#%d】获取到client[%d]\n", time.Now().Format("15:04:05"), index, client.id)
			time.Sleep(time.Second * 5)
			fmt.Printf("Time:%s | 【goruntine#%d】使用完毕，将client[%d]放回池子\n", time.Now().Format("15:04:05"), index, client.id)

			// 将连接放回池子
			pool.Push(client)
		}(i)
	}

	wg.Wait()
}
