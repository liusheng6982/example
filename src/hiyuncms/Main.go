package main


import (
	"fmt"
	"syscall"
	"os"
	"net/http"
	"os/signal"
	_"hiyuncms/models"
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
	bServer :=func () {
		err := http.ListenAndServe(fmt.Sprintf(":%d", config.GetInt("hiyuncms.server.backend.port")), routes.BackendRoute)
		if err != nil {
			fmt.Printf("init runBackendServer error%s\n", err)
		}
	}
	go bServer()
}


func runFrontendServer(){
	fServer :=func (){
		err := http.ListenAndServe(fmt.Sprintf(":%d", config.GetInt("hiyuncms.server.frontend.port")), routes.FrontendRoute)
		if err != nil {
			fmt.Printf("%c[0;40;31m%s%s%c[0m\n", 0x1B,  "init runFrontendServer error:",err, 0x1B)
		}
	}
	go fServer()
}

func RegService(){
	service.RegService()
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	<-sig
	fmt.Println("服务器以优雅的停止")
	service.UnRegService()
}