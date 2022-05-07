package main

import (
	"fmt"
	"sync"
)

var once sync.Once
var instance *Singleton

type Singleton struct {
}

func GetInstance() *Singleton {
	once.Do(func() {
		instance = &Singleton{}
	})
	return instance
}

func main() {
	i1 := GetInstance()
	i2 := GetInstance()

	fmt.Printf("i1:%p\n", i1)
	fmt.Printf("i2:%p\n", i2)
}
