package yy

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"hiyuncms/models/yy"
	"hiyuncms/controllers/frontend"
	"log"
	"time"
	"hiyuncms/models"

	"github.com/smartwalle/alipay"
	"github.com/astaxie/beego/logs"
	"net/http"
)

var a = []byte(`MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAs6bAiqNYKhLQ/U4ecM6v
 PYXV6dvfk1giS5ulPUe1OzJQnXqZmtKcEIfBOF+NuSWNus34Pm+yn946Ir7S8yNG
 KzY0Anwg3clw6KXj6O17SI9Buk2SZJ9GmBOgJkAnYh4hUVk4sRw3Z+hlijFvuWL6
 XB3rVBAhwlxDCzTioc2Mu7g5sy2Tm6F2G19p1sGepLM+wtg5mskuhftVYAYWuLWz
 Gou1DTQ09Xmp7REzdnVRO/LYDiwpjOCBNUEcTyWARSzFQHVXKixXsbVM4xdA4kAS
 DIaOvhqsBzvh1aZqW8iNoFn3m1Iz+ZA38ftjU9MAjvqsLNBNrzaV3dGtK8OArB1q
 NwIDAQAB`)

var b = []byte(`-----BEGIN RSA PRIVATE KEY-----
MIIEowIBAAKCAQEAs6bAiqNYKhLQ/U4ecM6vPYXV6dvfk1giS5ulPUe1OzJQnXqZmtKcEIfBOF+NuSWNus34Pm+yn946Ir7S8yNGKzY0Anwg3clw6KXj6O17SI9Buk2SZJ9GmBOgJkAnYh4hUVk4sRw3Z+hlijFvuWL6XB3rVBAhwlxDCzTioc2Mu7g5sy2Tm6F2G19p1sGepLM+wtg5mskuhftVYAYWuLWzGou1DTQ09Xmp7REzdnVRO/LYDiwpjOCBNUEcTyWARSzFQHVXKixXsbVM4xdA4kASDIaOvhqsBzvh1aZqW8iNoFn3m1Iz+ZA38ftjU9MAjvqsLNBNrzaV3dGtK8OArB1qNwIDAQABAoIBAECf0I1Omw1vfVxReKPNxb4U4dFhNbjUMGoQUE2N+QSVYeh0TMMj1d4gZ4I25U1f0+8J3q3fEltt3m3XRR0PaFNtCSKHsm714rbdzfFhVELSvasd8nZd0VAtZyO7Wi9ydTvFI56abtfjAnvGstD2aOcwIBI6R1qaQ8fJO2lG/sQbs8hhLQmoYWS0cUQzZuIkaF89uo7UmUVy8hz1ACajmaA1oCZjwnMaYbUKrhv4cNI/L2z/3OlF5gQy731jSHSuqsCPKB7AOcaMEqcFCIHyxxn7UicgJYu1NK3UCnJ1hLLFSygM60MsHSrNHv0kink9G8vdHoGYv40FThzF8tsQ4qECgYEA3pS+Qm/8ABI3oF6rgdT6HvhZ1EmlN2khtZnxXkiLj0I5g3EUuJ4uwR6+cgrAr3A8uhoh5LWHx90SjLuphXcPbyISmi2mWA+Bg35ZpWxfvyiI2XdJPzadvWsVi+tcRYECsafAiQYG5VLKRKkDFPxWqQ4TeUrKOvNG1FaNFdGwRxECgYEAzp/wcoJiF257bxJ2AdIskjGnG7KvUTV/YUrXY0wEkKzymTRE1NfPbGEfELSuxTtfv5Y/hbB+ytkZmHdRkPldUAvL0CG7u92D/SKtmzX9ahaDYWWtDRV6I/4OKt60+xMwmCv4Oi4Vdv1K4SjCcALXCb6Mw42ZAW4UGlKJAPprbMcCgYByq+cpi1AlKT2HXb62cOc7tW9yM07vMTawvNLhZDaiY9gFo+itBLHJxPERCAElYYmnx3bWwb9mdLrtznET1bcZ5k/3JrWggLyU5i+BTkg1z8hRYWdXLeguglDjeSpclI6ywF4tOfGri++xV/HCig6LojjeMG3n2RYQp1agext6QQKBgQCpmsmBBRtFho/VbX7mIcIqQo2cA8E61MH5d7hzLnv00bHVJf12BKujl9krGlT3WrROjCMaNvTsxuXmq9KNQNNimDw1XOs/2yWzjFqas+eOxGoVcaNpwP5gOvMgJ2zBR1A1KKp5/0fpQyLKzW1FCl+/BOWAw2MbtGLV9He3ENdLEwKBgCCL/fVN8AFUj0mUIn7ih9J+glH4DVfLMmTKZTt5ad4ht9bxoUXDE9TVAs1hO+F8Q6qs3wliacERm0W/9L8/VUP+ztfzq65wM4M4Bse69bewfwKiBVMb0jIAE0JHIfCMHJLAuEhkASLXavGQlrdNtPWA3AuIPu2NaHPMbOXVX8/q
-----END RSA PRIVATE KEY-----`)


func AliPayNotify(c *gin.Context){
	alipayClient := alipay.New("2016091200494527", "2088102175304454",[]byte(a),[]byte(b), false)
	alipayClient.AliPayPublicKey = a
	alipayClient.GetTradeNotification(c.Request)
}

func AliPrePay(c *gin.Context){
	vipLevel,_:= c.GetPostForm("VipLevel")

	VipLevel, err:= strconv.Atoi( vipLevel )
	if err != nil {
		log.Printf("支付时，vipvel错误：%s", err.Error() )
	}
	sessionInfo := frontend.GetSessionInfo(c)
	if sessionInfo.UserName == "" {
		log.Printf("支付时，用户还没有登录" )
	}
	payment := yy.YyPayment{}
	payment.VipLevel = VipLevel
	payment.PayStatus = 0
	payment.CompanyId = sessionInfo.CompanyId
	payment.UserId = sessionInfo.UserId
	payment.OrderInfo = ""
	payment.OrderNo = ""
	payment.PayTime = models.Time(time.Now())
	payment.PayAmount = 1//120000

	yy.SavePayment( &payment )


	alipayClient := alipay.New("2016091200494527", "2088102175304454",[]byte(a),[]byte(b), false)
	alipayClient.AliPayPublicKey = a

	var p = alipay.AliPayTradeWapPay{}
	p.NotifyURL = "xxx"
	p.Subject = "标题"
	p.OutTradeNo = "传递一个唯一单号"
	p.TotalAmount = "10.00"
	p.ProductCode = "商品编码"

	url, err := alipayClient.TradeWapPay(p)
	if err != nil {
		log.Printf( err.Error() )
	}
	logs.Info( url )


	if url != nil {
		c.Redirect(http.StatusFound, url.String())
		return
	}
}