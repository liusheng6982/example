package util


import (
	"fmt"
	"github.com/KenmyZhang/aliyun-communicate/app"
)
var (
	gatewayUrl = "http://dysmsapi.aliyuncs.com/"
	accessKeyId = "LTAIDICSO5JOkFV2"
	accessKeySecret = "wL64iUsD1CRm7hme4J5ATuWxfApaL3"
	phoneNumbers = "13918015069"
	signName = "阿里云短信测试专用"
	templateCode = "SMS_118250052"
	templateParam = "{\"code\":\"1234\"}"
)

func TestSms() {
	smsClient := app.NewSmsClient(gatewayUrl)
	if result, err := smsClient.Execute(accessKeyId, accessKeySecret, phoneNumbers, signName, templateCode, templateParam); err != nil {
		fmt.Println("error:", err.Error())
	} else {
		for key, value := range result {
			fmt.Println("key:", key, " value:",value)
		}
	}

}