package yy

import (
	"hiyuncms/models/yy"
	"hiyuncms/models"
	"time"
	"github.com/satori/go.uuid"
)

func PaymentPrePay(vipLevel, companyId, userId int64)*yy.YyPayment{
	paymentInfo := yy.GetPayInfo(vipLevel)
	payment := yy.YyPayment{}
	payment.VipLevel = vipLevel
	payment.PayStatus = 0
	payment.CompanyId = companyId
	payment.UserId = userId
	payment.OrderInfo = paymentInfo.PayInfo
	payment.OrderNo = uuid.NewV4().String()
	payment.PayTime = models.Time(time.Now())
	payment.PayAmount = paymentInfo.PayAmount

	yy.SavePayment( &payment )

	return &payment

}

func PaymentSuccess(orderNo,tradeNo string){
	yyPayment := yy.GetPaymentByOrderNo(orderNo)
	if yyPayment.PayStatus == 1{
		return
	}
	yyPayment.PayStatus = 1
	yyPayment.TradeNo = tradeNo

	yyCompany := yy.GetById( yyPayment.CompanyId )
	yyCompany.VipLevel = yyPayment.VipLevel
	tempTime := time.Time(  yyCompany.VipExpired )
	if tempTime.Before( time.Now() ){
		yyCompany.VipExpired = models.Date(time.Now().AddDate(1, 0, 0))
	}else{
		yyCompany.VipExpired = models.Date(tempTime.AddDate(1, 0, 0))
	}
	yy.UpdateCompany( yyCompany )
	yy.UpdatePamyment( yyPayment )
}
