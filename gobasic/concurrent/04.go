package main

import (
	"sync"
	"sync/atomic"
)

type RedisConn struct {
	once sync.Once
	init uint32
}

func (this *RedisConn) Init() {
	this.once.Do(func() {

		// do redis connection

		atomic.StoreUint32(&this.init, 1)
	})
}
func (this *RedisConn) IsConnect() bool { // 另外一个goroutine
	return atomic.LoadUint32(&this.init) != 0
}
