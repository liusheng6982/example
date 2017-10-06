package service

import (
	"fmt"
	consulapi "github.com/hashicorp/consul/api"
	"log"
	"net"
	"os"
	"hiyuncms/config"
)

var ip string

func getLocalPort() int {
	return config.GetInt("hiyuncms.server.backend.port")
}

func init()  {
	addresses, err := net.InterfaceAddrs()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, address := range addresses {
		// 检查ip地址判断是否回环地址
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				ip = ipnet.IP.String()
				break
			}
		}
	}
}

func getLocalIp() string{
	return  ip
}

func getConfig()  *consulapi.Config{
	value := consulapi.Config{
			Address: fmt.Sprintf("%s:%d",config.GetValue("consul.server.ip"),config.GetInt("consul.server.port")),
			Scheme: config.GetValue("consul.server.scheme")}

	return &value
}

func getServiceId() string {
	return fmt.Sprintf("%s:%d", getLocalIp(), getLocalPort())
}

func isRegistry() bool {
	return config.GetBool("consul.server.registry.registry")
}

func UnRegService() {
	if !isRegistry(){
		return
	}
	agentClient, _ := consulapi.NewClient(getConfig())
	agentClient.Agent().ServiceDeregister(getServiceId())
}

func RegService() {
	if !isRegistry(){
		return
	}
	appName := config.GetValue("hiyuncms.application.name")
	var err error = nil
	agentClient, err := consulapi.NewClient( getConfig() )
	if err != nil {
		log.Fatal("consul client error : ", err)
	}
	//创建一个新服务。
	registration := new(consulapi.AgentServiceRegistration)
	registration.ID = getServiceId()
	registration.Name = appName
	registration.Port = getLocalPort()
	registration.Tags = []string{appName}
	registration.Address = getLocalIp()

	//增加check。
	check := new(consulapi.AgentServiceCheck)
	check.HTTP = fmt.Sprintf("http://%s:%d%s", registration.Address, registration.Port, "/ping")
	//设置超时 5s。
	check.Timeout = "5s"
	//设置间隔 5s。
	check.Interval = "5s"
	//注册check服务。
	registration.Check = check

	err = agentClient.Agent().ServiceRegister(registration)

	if err != nil {
		log.Fatal("register server error : ", err)
	}

}
