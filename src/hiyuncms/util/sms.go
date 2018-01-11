package util


import (
	"fmt"
	"github.com/KenmyZhang/aliyun-communicate/app"
)
var (
	gatewayUrl = "http://dysmsapi.aliyuncs.com/"
	accessKeyId = "LTAIDICSO5JOkFV2"
	accessKeySecret = "wL64iUsD1CRm7hme4J5ATuWxfApaL3"
	signName = "e医链"
	templateCode = "SMS_119086797"
)	

func SendSms( code,mobile string ) error {
	templateParam1 := fmt.Sprintf("{\"code\":\"%s\"}", code)
	smsClient := app.NewSmsClient(gatewayUrl)
	result, err := smsClient.Execute(accessKeyId, accessKeySecret, mobile, signName, templateCode, templateParam1)
	if err != nil {
		fmt.Println("error:", err.Error())
	} else {
		for key, value := range result {
			fmt.Println("key:", key, " value:",value)
		}
	}
	return  err
}

func TestSms() {
	SendSms("1235", "13918015069")
}