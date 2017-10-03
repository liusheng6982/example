package main


import (
	"fmt"
	"syscall"
	"os"
	"net/http"
	"os/signal"
	"hiyuncms/service"
	"hiyuncms/routes"
	"hiyuncms/config"
)


func main() {
	runFrontendServer()
	runBackendServer()
	RegService()
}

func runBackendServer()  {
	go http.ListenAndServe(fmt.Sprintf(":%d", config.GetInt("hiyuncms.server.backend.port")), routes.BackendRoute)
}


func runFrontendServer(){
	go http.ListenAndServe(fmt.Sprintf(":%d", config.GetInt("hiyuncms.server.frontend.port")), routes.FrontendRoute)
}

func RegService(){
	service.RegService()
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	<-sig
	fmt.Println("服务器以优雅的停止")
	service.UnRegService()
}