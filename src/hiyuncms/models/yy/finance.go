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
	VipLevel 		int 				`xorm:"int"`
	PayAmount		int64				`xorm:"bigint"`
	PayTime			models.Time			`xorm:"datetime"`
	PayStatus		int					`xorm:"int"`
	TradeNo         string              `xorm:"varchar(128)"`
}

func init()  {
	err := models.DbMaster.Sync2( YyCompany{})
	log.Println( "init table yy_payment", models.GetErrorInfo(err))
}

func SavePayment(payment * YyPayment){
	models.DbMaster.Insert( payment )
}
