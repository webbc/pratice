package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"pratice/web_framework/framework"
	"syscall"
	"time"
)

func main() {
	core := framework.NewCore()
	registerRouter(core)
	server := &http.Server{Addr: ":9999", Handler: core}

	go func() {
		err := server.ListenAndServe()
		if err != nil {
			fmt.Println("ListenAndServe err:%v", err.Error())
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	timeoutCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(timeoutCtx); err != nil {
		fmt.Println("Shutdown err:%v", err.Error())
	}
}
