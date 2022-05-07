package baseserv

import (
	"fmt"
	"sync"
)

type ServiceManager struct {
	servs sync.Map
}

func NewServiceManager() *ServiceManager {
	return &ServiceManager{}
}

func (this *ServiceManager) GetService(id uint32) *Service {
	if serv, ok := this.servs.Load(id); ok {
		return serv.(*Service)
	}
	return nil
}

func (this *ServiceManager) AddService(cfg *Config) bool {

	// 判断service是否存在
	if this.GetService(cfg.Id) != nil {
		fmt.Printf("service already exist %d\n", cfg.Id)
		return false
	}

	return this.startService(cfg)
}

func (this *ServiceManager) DelService(id uint32) {
	this.servs.Delete(id)
}

func (this *ServiceManager) startService(cfg *Config) bool {
	service := NewService(cfg, this)
	if !service.serv.Init() {
		fmt.Printf("ServiceManager startService fail %d\n", cfg.Id)
		return false
	}
	go service.Start()
	this.servs.Store(cfg.Id, service)
	return true
}
