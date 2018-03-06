package payment

import(
 	"github.com/ascoders/alipay"
	"github.com/gin-gonic/gin"
)


var alipayClient alipay.Client

func init()  {
	alipayClient = alipay.Client{
		Partner   : "", // 合作者ID
		Key       : "", // 合作者私钥
		ReturnUrl : "", // 同步返回地址
		NotifyUrl : "", // 网站异步返回地址
		Email     : "", // 网站卖家邮箱地址
	}
}

func GetUnifPay() string {
	form := alipayClient.Form(alipay.Options{
		OrderId:  "123",	// 唯一订单号
		Fee:      99.8,		// 价格
		NickName: "翱翔大空",	// 用户昵称，支付页面显示用
		Subject:  "充值100",	// 支付描述，支付页面显示用
	})
	return form
}

func AliPay(c *gin.Context)  {
	//c.Writer.Header().
	//c.Writer.Write("text/html;charset=utf-8")


}
