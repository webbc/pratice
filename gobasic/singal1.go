package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

type SignalHandler interface {
	HandleSigHup()
	HandleSigInt()
}

type MySignalHandler struct {
	SignalHandler
}

func (this *MySignalHandler) HandleSigHup() {
	fmt.Println("MySignalHandler HandleSigHup")
}

func (this *MySignalHandler) HandleSigInt() {
	fmt.Println("MySignalHandler HandleSigInt")
}

func signalRun(handler SignalHandler) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM, syscall.SIGHUP)
	for {
		select {
		case sign := <-c:
			switch sign {
			case syscall.SIGHUP:
				if handler != nil {
					handler.HandleSigHup()
				}
			case syscall.SIGINT:
				if handler != nil {
					handler.HandleSigInt()
				}
			default:
				return
			}
		}
	}

}

func main() {
	signalRun(&MySignalHandler{})
}
