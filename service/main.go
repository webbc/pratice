package main

import (
	"pratice/service/baseserv"
	"pratice/service/ntype"
	"pratice/service/service"
)

func main() {
	closeChan := make(chan struct{})
	InitEtcdService()
	<-closeChan
}

func InitEtcdService() bool {
	serv := service.NewEtcdService()
	if serv != nil {
		cfg := &baseserv.Config{
			Id:       ntype.SERVICE_ID_ETCD,
			TimeTick: ntype.COMMON_TIME_TICK,
			ChanSize: ntype.COMMON_CHAN_SIZE,
			Serv:     serv,
		}
		if baseserv.AddService(cfg) {
			return true
		}
		return false
	}
	return false
}
