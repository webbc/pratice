package service

import (
	"fmt"
	"pratice/service/baseserv"
)

type EtcdService struct {
	*baseserv.CommonService
}

func NewEtcdService() *EtcdService {
	return &EtcdService{}
}

func (this *EtcdService) Init() bool {
	fmt.Println("EtcdService Init...")
	return true
}

func (this *EtcdService) Start() bool {
	fmt.Println("EtcdService Start...")
	return true
}

func (this *EtcdService) Update(now int64) {
	fmt.Printf("EtcdService Update now:%v\n", now)
}

func (this *EtcdService) Stop() {
	fmt.Println("EtcdService Stop...")
}

func (this *EtcdService) Process(msg interface{}) bool {
	fmt.Printf("EtcdService Process msg:%v\n", msg)
	return true
}

func (this *EtcdService) Panic(msg interface{}) bool {
	fmt.Printf("EtcdService Panic msg:%v\n", msg)
	return true
}
