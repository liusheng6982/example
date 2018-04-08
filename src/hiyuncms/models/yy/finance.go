package yy

import (
	"hiyuncms/models"
	"log"
)

type YyPayment struct {
	Id 				int64 				`xorm:"pk BIGINT autoincr"`
    CompanyId 		int64 				`xorm:"bigint"`
	UserId          int64   			`xorm:"bigint"`
	OrderNo			string				`xorm:"varchar(100)"`
	OrderInfo       string 				`xorm:"varchar(128)"`
	VipLevel 		int64				`xorm:"int"`
	PayAmount		int64				`xorm:"bigint"`
	PayTime			models.Time			`xorm:"datetime"`
	PayStatus		int					`xorm:"int"`
	TradeNo         string              `xorm:"varchar(128)"`
}

func init()  {
	err := models.DbMaster.Sync2( YyPayment{})
	log.Println( "init table yy_payment", models.GetErrorInfo(err))
}

func SavePayment(payment * YyPayment){
	_,err := models.DbMaster.Insert( payment )
	if err != nil {
		log.Printf("保存支付信息出错！:%s", err.Error())
	}
}

func GetPaymentByOrderNo(orderNo string)*YyPayment{
	yyPayment := YyPayment{}
	_, err := models.DbSlave.Table(YyPayment{}).Where("Order_No=?", orderNo).Get(&yyPayment)
	if err != nil {
		log.Printf("根据%s获取YyPayment出错:%s",orderNo, err.Error())
	}
	return &yyPayment
}

func UpdatePamyment(payment * YyPayment)  {
	models.DbMaster.Id(payment.Id).Update(payment)
}
