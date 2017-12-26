package util

/*`
import (
	"github.com/ascoders/alipay"
	"strconv"
)


import (
	"strconv"
	"github.com/ascoders/alipay"

	"github.com/astaxie/beego"
	"hiyuncms/config"
)

type AlipayController struct {

}

func newClient() *alipay.Client {
	return &alipay.Client{
		Partner:   config.GetValue("alipartner"),                               // 合作者ID
		Key:       config.GetValue("alikey"),                                   // 合作者私钥
		ReturnUrl: "http://" + config.GetValue("domainurl") + "/alipay/return", // 同步返回地址
		NotifyUrl: "http://" + config.GetValue("domainurl") + "/alipay/notify", // 网站异步返回地址
		Email:     config.GetValue("aliemail"),                                 // 网站卖家邮箱地址
	}
}

func (this *AlipayController) Native() {
	orderNo := this.GetString("orderNo") //获取自己的订单号
	schemestr := this.Ctx.Input.Site()
	alipayClient := newClient()
	fee, _ := strconv.ParseFloat("100.5")//价格转换
	ots := alipay.Options{
		OrderId:            orderNo,
		Fee:                float32(fee),
		NickName:           "ricky",
		Subject:            "某某订单" + orderNo,
		Extra_common_param: schemestr, //加上自己需要用到的参数
	}
	form := alipayClient.Form(ots)
	res := map[string]interface{}{"form": form}
	this.Data["json"] = res
	this.ServeJSON()
}

func (this *AlipayController) Return() {
	alipayClient := newClient()
	result := alipayClient.Return(&this.Controller)
	//beego.Debug("notify", result)
	if result.Status == 1 { //付款成功，处理订单
		//处理订单
		if result.Extra_common_param != "" {
			url := typestr[1] + "/order/detail/" + result.OrderNo
			this.Ctx.Redirect(302, url)
		}
	} else {
		res := map[string]interface{}{"msg": "来源验证失败"}
		this.Data["json"] = res
		this.ServeJSON()
	}
}

func (this *AlipayController) Notify() {
	alipayClient := newClient()
	result := alipayClient.Notify(&this.Controller)

	timetest := this.GetString("gmt_payment")
	if result.Status == 1 { //付款成功，处理订单
		sendData := make(map[string]interface{})
		sendData["id"] = result.OrderNo
		sendData["trade_no"] = result.TradeNo
		sendData["paid_time"] = timetest
		sendData["payment_type"] = "alipay"
		sendData["payment_amount"] = result.TotalFee
		//这里处理自己的业务逻辑
		if result.Extra_common_param != "" {
			//your method 例如修改数据库中订单的状态为付款。。
		}

	}
}
`*/