package baseserv

import (
	"fmt"
	"runtime/debug"
	"time"
)

type Service struct {
	id        uint32
	timeTick  uint32
	messages  chan interface{}
	serv      IService
	servM     *ServiceManager
	closeChan chan struct{}
}

func NewService(cfg *Config, serverManager *ServiceManager) *Service {
	return &Service{
		id:        cfg.Id,
		timeTick:  cfg.TimeTick,
		serv:      cfg.Serv,
		servM:     serverManager,
		closeChan: make(chan struct{}),
		messages:  make(chan interface{}, cfg.ChanSize),
	}
}

func (this *Service) Start() {
	defer func() {
		if err := recover(); err != nil {
			if this.serv.Panic(err) {
				panic(err)
			} else {
				fmt.Printf("%v", debug.Stack())
			}
		} else {
			this.Stop()
		}
	}()

	// 启动失败
	if !this.serv.Start() {
		fmt.Printf("Service Start fail %d\n", this.id)
		return
	}

	ticker := time.NewTicker(time.Duration(this.timeTick) * time.Millisecond)
	for {
		select {
		case msg := <-this.messages:
			this.serv.Process(msg)
		case now := <-ticker.C:
			this.serv.Update(now.Unix())
		case <-this.closeChan:
			return
		}
	}
}

func (this *Service) Stop() {
	for more := true; more; {
		select {
		case msg := <-this.messages:
			this.serv.Process(msg)
		default:
			more = false
		}
	}
	this.stopService()
}

func (this *Service) stopService() {
	this.serv.Stop()
	if this.servM != nil {
		this.servM.DelService(this.id)
	}
}
