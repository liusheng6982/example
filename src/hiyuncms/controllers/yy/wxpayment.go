package yy

import (
	"time"
	"strconv"
	"fmt"
	"os"
	"encoding/base64"
	"hiyuncms/payment/Wxpay"
	"github.com/gin-gonic/gin"
	"odeke-em/qr"
)

func  WxPrePay(c *gin.Context) {

   orderNumber := c.PostForm ("id") //获取订单号
	payAmount := c.PostForm("price")       //获取价格
	params := make(map[string]interface{})
	params["body"] = "****company-" + orderNumber //显示标题
	params["out_trade_no"] = orderNumber
	params["total_fee"] = payAmount
	params["product_id"] = orderNumber
	params["attach"] = "abc" //自定义参数

	var modwx Wxpay.UnifyOrderReq
	res := modwx.CreateOrder(c.ClientIP(), params)

	//拿到数据之后，需要生成二维码。
	Img(res.Code_url)

}

func WxPayNotify(c *gin.Context) {
	var notifyReq Wxpay.WXPayNotifyReq
	res := notifyReq.WxpayCallback(c.Request, c.Writer)
	//beego.Debug("res",res)
	if res != nil {
		//这里可以组织res的数据 处理自己的业务逻辑：
		sendData := make(map[string]interface{})
		sendData["id"] = res["out_trade_no"]
		sendData["trade_no"] = res["transaction_id"]
		paid_time, _ := time.Parse("20060102150405", res["time_end"].(string))
		paid_timestr := paid_time.Format("2006-01-02 15:04:05")
		sendData["paid_time"] = paid_timestr
		sendData["payment_type"] = "wxpay"
		intfee := res["cash_fee"].(int)
		floatfee := float64(intfee)
		cashfee := floatfee / 100
		sendData["payment_amount"] = strconv.FormatFloat(cashfee, 'f', 2, 32)

		//api(sendData)...自己的逻辑处理
		//

	}
}

func Img(url string) string {
	code, err := qr.Encode(url, qr.H)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	imgByte := code.PNG()
	str := base64.StdEncoding.EncodeToString(imgByte)

	return str
}