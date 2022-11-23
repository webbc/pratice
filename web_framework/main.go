package main

import (
	"net/http"
	"pratice/web_framework/framework"
)

func main() {
	core := framework.NewCore()
	registerRouter(core)
	server := &http.Server{Addr: ":9999", Handler: core}
	server.ListenAndServe()
}
