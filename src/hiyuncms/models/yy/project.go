package yy

import (
	"log"
	"hiyuncms/models"
)

type YyPurchase struct {
	Id             int64       			`xorm:"pk BIGINT autoincr" json:"id"`
	PurchaseName   string      			`xorm:"varchar(100) notnull"`
	PurchaseNo     string      			`xorm:"varchar(50)"`
	PurchaseType   string      			`xorm:"varchar(20)"` //合格供应商，定向，公开
	ProjectContent      string      `xorm:"varchar(2000)"`
	ProjectImage        string      `xorm:"varchar(200)"`
	ExpiredDate         models.Date `xorm:"DateTime"`
	CompanyId           int64       `xorm:"BIGINT"`
	CompanyName         string      `xorm:"varchar(50)"`
	QuotePriceEndTime   models.Date `xorm:"DateTime"`
	QuotePriceTaxRate   int         `xorm:"int"`
	QuotePriceLogistics int         `xorm:"int"`
	RequireImage        string      `xorm:"varchar(200)"`
	Remark              string      `xorm:"varchar(2000)"`
	DeliveryTime        models.Date `xorm:"DateTime"`
	Material            string      `xorm:"varchar(20)"`
	Detail              string      `xorm:"text"`
	Published           int         `xorm:"int"`
	ImpFlag             int         `xorm:"int"`
	ImpId               string      `xorm:"varchar(100)"`
	CreateTime          models.Date `xorm:"DateTime"`
}

type YyInviteTender  struct {
	Id             int64       			`xorm:"pk BIGINT autoincr" json:"id"`
	ProjectName    string      			`xorm:"varchar(100) notnull"`
	ProjectNo      string      			`xorm:"varchar(50)"`
	Type           string      			`xorm:"varchar(20)"` //合格供应商，定向，公开
	Contact	   	   string      			`xorm:"varchar(30)"`
	ContactPhone   string 	   			`xorm:"varchar(30)"`
	SubmitTenderEndTime models.Date 	`xorm:"DateTime"`
	OpenTenderTime      models.Date  	`xorm:"DateTime"`
	TenderDocument		string      	`xorm:"varchar(130)"`
	Published           int         	`xorm:"int"`
	ImpFlag				int 			`xorm:"int"`
	CreateTime				models.Date `xorm:"DateTime"`
	ImpId            string             `xorm:"varchar(100)"`
	Recommend        int             	`xorm:"int"`
	WinBidFlag       int				`xorm:"int"`
	WinBidCompany    string             `xorm:"varchar(100)"`

}

func init()  {
	err := models.DbMaster.Sync2( YyPurchase{} )
	log.Println( "init table yy_purchase ", models.GetErrorInfo(err))

	err = models.DbMaster.Sync2( YyInviteTender{} )
	log.Println( "init table yy_invite_tender ", models.GetErrorInfo(err))
}

func GetTopInviteTender(size int) []*YyInviteTender {
	inviteTenderList := make([]*YyInviteTender, 0)
	err := models.DbSlave.Table(YyInviteTender{}).Limit(size,0).Find(&inviteTenderList)
	if err != nil {
		log.Printf("获取YyInviteTender数据:%s", models.GetErrorInfo(err))
	}
	return inviteTenderList
}

func GetTopRecommendInviteTender(size int) []*YyInviteTender {
	inviteTenderList := make([]*YyInviteTender, 0)
	err := models.DbSlave.Table(YyInviteTender{}).Where("recommend = 1").Limit(size,0).Find(&inviteTenderList)
	if err != nil {
		log.Printf("获取YyInviteTender数据:%s", models.GetErrorInfo(err))
	}
	return inviteTenderList
}

func GetTopWinBidInviteTender(size int) []*YyInviteTender {
	inviteTenderList := make([]*YyInviteTender, 0)
	err := models.DbSlave.Table(YyInviteTender{}).Where("win_bid_flag = 1").Limit(size,0).Find(&inviteTenderList)
	if err != nil {
		log.Printf("获取YyInviteTender数据:%s", models.GetErrorInfo(err))
	}
	return inviteTenderList
}

func GetInviteTenderById(id int64) * YyInviteTender  {
	log.Printf("id=%d\n", id)
	inviteTender := YyInviteTender{}
	_, err := models.DbSlave.ID(id).Get(&inviteTender)
	if err != nil {
		log.Printf("获取YyInviteTender数据:%s", models.GetErrorInfo(err))
	}
	return &inviteTender
}

func GetAllInviteTenderByPage(page *models.PageRequest) * models.PageResponse  {
	inviteTenderList := make([]*YyInviteTender, 0)
	err := models.DbSlave.Table(YyInviteTender{}).Limit(page.Rows, (page.Page - 1)* page.Rows).Find(&inviteTenderList)
	if err != nil {
		log.Printf("获取YyInviteTender数据:%s", models.GetErrorInfo(err))
	}
	records,_:= models.DbSlave.Table(YyInviteTender{}).Count(YyInviteTender{})
	pageResponse := models.InitPageResponse(page, inviteTenderList, records)
	return pageResponse
}



func GetTopPurchase(size int) []*YyPurchase {
	yyPurchaseList := make([]*YyPurchase, 0)
	err := models.DbSlave.Table(YyPurchase{}).Limit(size,0).Find(&yyPurchaseList)
	if err != nil {
		log.Printf("获取YyInviteTender数据:%s", models.GetErrorInfo(err))
	}
	return yyPurchaseList
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
