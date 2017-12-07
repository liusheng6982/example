package yy

import (
	"log"
	"hiyuncms/models"
)

type YyPurchase struct {
	Id             int64       `xorm:"pk BIGINT autoincr" json:"id"`
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
	Contact	   	   string      `xorm:"varhcar(30)"`
	ContactPhone   string 	   `xorm:"varhcar(30)"`
	SubmitTenderEndTime models.Date `xorm:"DateTime"`
	OpenTenderTime     models.Date  `xorm:"DateTime"`
	TenderDocument		string      `xorm:"varhcar(130)"`


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


func GetAllInviteTenderByPage(page *models.PageRequest) * models.PageResponse  {
	inviteTenderList := make([]*YyInviteTender, 0)
	err := models.DbSlave.Table(YyPurchase{}).Limit(page.Rows, (page.Page - 1)* page.Rows).Find(&inviteTenderList)
	if err != nil {
		log.Printf("获取YyInviteTender数据:%s", models.GetErrorInfo(err))
	}

	records,_:= models.DbSlave.Table(YyPurchase{}).Count(YyInviteTender{})
	pageResponse := models.InitPageResponse(page, inviteTenderList, records)
	return pageResponse
}

func GetAllYyPurchaseByPage(page *models.PageRequest) * models.PageResponse  {
	purchaseList := make([]*YyPurchase, 0)
	err := models.DbSlave.Table(YyPurchase{}).Limit(page.Rows, (page.Page - 1)* page.Rows).Find(&purchaseList)
	if err != nil {
		log.Printf("获取YyPurchase数据:%s", models.GetErrorInfo(err))
	}

	records,_:= models.DbSlave.Table(YyPurchase{}).Count(YyPurchase{})
	pageResponse := models.InitPageResponse(page, purchaseList, records)
	return pageResponse
}
