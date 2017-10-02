package service

import (
	"fmt"
	consulapi "github.com/hashicorp/consul/api"
	"log"
	"net"
	"os"
)

var ip string
var Port int
var serviceId string
var serviceNname string

var agentClinet *consulapi.Client

func init() {
	addrs, err := net.InterfaceAddrs()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, address := range addrs {
		// 检查ip地址判断是否回环地址
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				ip = ipnet.IP.String()
				fmt.Println(ipnet.IP.String())
				break
			}

		}
	}
}

func UnRegService() {
	agentClinet.Agent().ServiceDeregister(serviceId)
}

func RegService() {

	serviceNname = "newCar2"

	serviceId = fmt.Sprintf("%s:%d", ip, Port)

	config := consulapi.Config{Address: "127.0.0.1:8500", Scheme: "http"}
	var err error = nil
	agentClinet, err = consulapi.NewClient(&config)
	if err != nil {
		log.Fatal("consul client error : ", err)
	}
	//创建一个新服务。
	registration := new(consulapi.AgentServiceRegistration)
	registration.ID = serviceId
	registration.Name = serviceNname
	registration.Port = Port
	registration.Tags = []string{serviceNname}
	registration.Address = ip

	//增加check。
	check := new(consulapi.AgentServiceCheck)
	check.HTTP = fmt.Sprintf("http://%s:%d%s", registration.Address, registration.Port, "/ping")
	//设置超时 5s。
	check.Timeout = "5s"
	//设置间隔 5s。
	check.Interval = "5s"
	//注册check服务。
	registration.Check = check
	log.Println("get check.HTTP:", check)

	err = agentClinet.Agent().ServiceRegister(registration)

	if err != nil {
		log.Fatal("register server error : ", err)
	}

}
