package yy

import (
	"log"
	"hiyuncms/models"
	"fmt"
)

type YyPorject struct {
	Id             						int64 			`xorm:"pk BIGINT autoincr" json:"id"`
	ProjectName    						string			`xorm:"varchar(100) notnull"`
	ProjectNo      						string			`xorm:"varchar(50)"`    //项目编号
	ProjectType    						int				`xorm:"int"`            //1:招标 2：采购
	ProjectContent 						string			`xorm:"text"`  			//内容
	ProjectFile	        				string      	`xorm:"varchar(200)"`   //项目文件
	ProjectRemark						string          `xorm:"varchar(2000)"`  //备注
	ProjectAreaName 					string          `xorm:"varchar(100)"`  //备注
	CompanyId         					int64    	    `xorm:"BIGINT"`         //公司ID
	CompanyName       					string   	    `xorm:"varchar(50)"`    //公司名称
	ContactPhone   						string 	   		`xorm:"varchar(30)"`    //联系人电话
	Contact	   	   						string          `xorm:"varchar(30)"`    //联系人

	ImpFlag          					int         	`xorm:"int"`			//是否导入
	ImpId            					string      	`xorm:"varchar(100)"`   //导入ID
	ImpPlatform							string          `xorm:"varchar(30)"`    //导入的平台

	PurchaseType   						string			`xorm:"varchar(20)"`   //合格供应商，定向，公开
	PurchaseExpiredDate                 models.Date 	`xorm:"DateTime"`      //采购有效期
	PurchaseQuotePriceEndTime   		models.Date		`xorm:"DateTime"`      //报价截止时间
	PurchaseDeliveryTime                models.Date		`xorm:"DateTime"`      //交货时间

	InviteType          	 			string      	`xorm:"varchar(20)"`   //合格供应商，定向，公开
	InviteEnterStartTime				models.Date     `xorm:"DateTime"`      //报名开始时间
	InviteEnterEndTime					models.Date     `xorm:"DateTime"`      //报名结束时间
	InviteWinBidFlag       				int				`xorm:"int"`           //是否中标
	InviteWinBidCompany    				string     		`xorm:"varchar(100)"`  //中标公司
	InviteSubmitTenderEndTime 			models.Date 	`xorm:"DateTime"`      //投标截止时间
	InviteOpenTenderTime      			models.Date  	`xorm:"DateTime"`      //开标时间

	Recommended							int             `xorm:"int"`           //是否是推荐项目
	Published        					int         	`xorm:"int"`           //是否发布
	PublishedTime						models.Date  	`xorm:"DateTime"`      //发布时间
	CreateTime          				models.Date 	`xorm:"DateTime"`      //创建时间


	BusinessCategory					string			`xorm:"varchar(20)"`   //建设、理疗器械、后勤物资、行政物资

	Auth								int				`xorm:"int"`           //查看权限
}

func init()  {
	err := models.DbMaster.Sync2( YyPorject{} )
	log.Println( "init table yy_project ", models.GetErrorInfo(err))
}

func GetTopInviteTender(size int) []*YyPorject {
	inviteTenderList := make([]*YyPorject, 0)
	err := models.DbSlave.Table(YyPorject{}).Limit(size,0).Where("project_type=1").Find(&inviteTenderList)
	if err != nil {
		log.Printf("获取YyInviteTender数据:%s", models.GetErrorInfo(err))
	}
	return inviteTenderList
}

func GetInviteTenderByCompanyIdAndSupplyId(size int, companyId,supplyId int64 )[]*YyPorject{
	hospitals := GetCompanyIdsBySupplyId( supplyId )
	hospitalIds := make( []int64, 0)
	for _,hospital:= range hospitals{
		hospitalIds = append(hospitalIds, hospital.HospitalId)
	}
	inviteTenderList := make([]*YyPorject, 0)
	err := models.DbSlave.Table(YyPorject{}).Limit(size,0).
		Where("project_type=1").Where("company_id = ?", companyId).Find(&inviteTenderList)
	if err != nil {
		log.Printf("获取YyInviteTender数据:%s", models.GetErrorInfo(err))
	}
	return inviteTenderList
}


func GetInviteTenderBySupplyId(size int, supplyId int64 )[]*YyPorject{
	hospitals := GetCompanyIdsBySupplyId( supplyId )
	hospitalIds := make( []int64, 0)
	for _,hospital:= range hospitals{
		hospitalIds = append(hospitalIds, hospital.HospitalId)
	}
	inviteTenderList := make([]*YyPorject, 0)
	err := models.DbSlave.Table(YyPorject{}).Limit(size,0).
		Where("project_type=1").Where("company_id in ? or auth= ?", hospitalIds, 0).Find(&inviteTenderList)
	if err != nil {
		log.Printf("获取YyInviteTender数据:%s", models.GetErrorInfo(err))
	}
	return inviteTenderList
}

func GetTopInviteTenderByCompanyId(size int, companyId int64) []*YyPorject {
	inviteTenderList := make([]*YyPorject, 0)
	err := models.DbSlave.Table(YyPorject{}).Limit(size,0).
		Where("project_type=1").Where(fmt.Sprintf("company_id=%d", companyId)).Find(&inviteTenderList)
	if err != nil {
		log.Printf("获取YyInviteTender数据:%s", models.GetErrorInfo(err))
	}
	return inviteTenderList
}

func GetTopRecommendInviteTender(size int) []*YyPorject {
	inviteTenderList := make([]*YyPorject, 0)
	err := models.DbSlave.Table(YyPorject{}).Where("Recommended = 1").Where("project_type=1").Limit(size,0).Find(&inviteTenderList)
	if err != nil {
		log.Printf("获取YyInviteTender数据:%s", models.GetErrorInfo(err))
	}
	return inviteTenderList
}

func GetTopWinBidInviteTender(size int) []*YyPorject {
	inviteTenderList := make([]*YyPorject, 0)
	err := models.DbSlave.Table(YyPorject{}).Where("invite_win_bid_flag = 1").Where("project_type=1").Limit(size,0).Find(&inviteTenderList)
	if err != nil {
		log.Printf("获取YyInviteTender数据:%s", models.GetErrorInfo(err))
	}
	return inviteTenderList
}

func GetInviteTenderById(id int64) * YyPorject  {
	log.Printf("id=%d\n", id)
	inviteTender := YyPorject{}
	_, err := models.DbSlave.ID(id).Get(&inviteTender)
	if err != nil {
		log.Printf("获取YyInviteTender数据:%s", models.GetErrorInfo(err))
	}
	return &inviteTender
}

func GetAllInviteTenderByPage(page *models.PageRequest) * models.PageResponse  {
	inviteTenderList := make([]*YyPorject, 0)
	err := models.DbSlave.Table(YyPorject{}).Where("project_type=1").Limit(page.Rows, (page.Page - 1)* page.Rows).Find(&inviteTenderList)
	if err != nil {
		log.Printf("获取YyInviteTender数据:%s", models.GetErrorInfo(err))
	}
	records,_:= models.DbSlave.Table(YyPorject{}).Where("project_type=1").Count(YyPorject{})
	pageResponse := models.InitPageResponse(page, inviteTenderList, records)
	return pageResponse
}



func GetTopPurchase(size int) []*YyPorject {
	yyPurchaseList := make([]*YyPorject, 0)
	err := models.DbSlave.Table(YyPorject{}).Where("project_type=2").Limit(size,0).Find(&yyPurchaseList)
	if err != nil {
		log.Printf("获取YyInviteTender数据:%s", models.GetErrorInfo(err))
	}
	return yyPurchaseList
}


func GetTopPurchaseByCompanyId(size int, companyId int64) []*YyPorject {
	yyPurchaseList := make([]*YyPorject, 0)
	err := models.DbSlave.Table(YyPorject{}).Where("project_type=2").Where(fmt.Sprintf("company_id=%d", companyId)).Limit(size,0).Find(&yyPurchaseList)
	if err != nil {
		log.Printf("获取YyInviteTender数据:%s", models.GetErrorInfo(err))
	}
	return yyPurchaseList
}

func GetAllYyPurchaseByPage(page *models.PageRequest) * models.PageResponse  {
	purchaseList := make([]*YyPorject, 0)
	err := models.DbSlave.Table(YyPorject{}).Where("project_type=2").Limit(page.Rows, (page.Page - 1)* page.Rows).Find(&purchaseList)
	if err != nil {
		log.Printf("获取YyPurchase数据:%s", models.GetErrorInfo(err))
	}
	records,_:= models.DbSlave.Table(YyPorject{}).Count(YyPorject{})
	pageResponse := models.InitPageResponse(page, purchaseList, records)
	return pageResponse
}
