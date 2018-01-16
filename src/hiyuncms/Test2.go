package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-vgo/robotgo"
)

func main() {
	fmt.Println("=-=-=-==-=-=-=-=-=-\nController-PC start...\nPC端占用端口号为:9090\n=-=-=-==-=-=-=-=-=-")

	//192.168.30.12
	http.HandleFunc("/", receiveClientRequest)

	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}

func receiveClientRequest(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	//  fmt.Println("收到客户端请求: ", r.Form)

	var key string = r.FormValue("key")
	fmt.Println("received key: ", key)

	robotgo.TypeStr(key)

}