package yy

import (
	"log"
	"hiyuncms/models"
)

type YyPurchase struct {
	Id             int64       `xorm:"pk BIGINT autoincr"`
	PurchaseName   string      `xorm:"varchar(100) notnull"`
	PurchaseNo     string      `xorm:"varchar(50)"`
	PurchaseType   string      `xorm:"varchar(20)"` //合格供应商，定向，公开
	ProjectContent string      `xorm:"varchar(2000)"`
	ProjectImage   string      `xorm:"varchar(200)"`
	ExpiredDate    models.Date `xorm:"DateTime"`
	CompanyId      int64       `xorm:"BIGINT"`
	CompanyName    string      `xorm:"varchar(50)"`

	QuotePriceEndTime   models.Date `xorm:"DateTime"`
	QuotePriceTaxRate   int         `xorm:"int"`
	QuotePriceLogistics int         `xorm:"int"`

	RequireImage string      `xorm:"varchar(200)"`
	Remark       string      `xorm:"varchar(2000)"`
	DeliveryTime models.Date `xorm:"DateTime"`
	Meterial     string 	 `xorm:"varchar(20)"`
	Detail       string      `xorm:"text"`
}

type YyInviteTender  struct {
	Id             int64       `xorm:"pk BIGINT autoincr"`
	ProjectName    string      `xorm:"varchar(100) notnull"`
	ProjectNo      string      `xorm:"varchar(50)"`
	Type           string      `xorm:"varchar(20)"` //合格供应商，定向，公开
	Contact	   	   string      ``
	ContactPhone   string 		``
	SubmitTenderEndTime models.Date
	OpenTenderTime     models.Date
	TenderDocument		string  ``


	ProjectContent string      `xorm:"varchar(2000)"`
	ProjectImage   string      `xorm:"varchar(200)"`
	ExpiredDate    models.Date `xorm:"DateTime"`
	CompanyId      int64       `xorm:"BIGINT"`
	CompanyName    string      `xorm:"varchar(50)"`

	QuotePriceEndTime   models.Date `xorm:"DateTime"`
	QuotePriceTaxRate   int         `xorm:"int"`
	QuotePriceLogistics int         `xorm:"int"`

	RequireImage string      `xorm:"varchar(200)"`
	Remark       string      `xorm:"varchar(2000)"`
	DeliveryTime models.Date `xorm:"DateTime"`
	Meterial     string 	 `xorm:"varchar(20)"`
	Detail       string      `xorm:"text"`
}

func init()  {
	err := models.DbMaster.Sync2( YyPurchase{} )
	log.Println( "init table yy_purchase ", models.GetErrorInfo(err))

	err = models.DbMaster.Sync2( YyInviteTender{} )
	log.Println( "init table yy_invite_tender ", models.GetErrorInfo(err))
}
