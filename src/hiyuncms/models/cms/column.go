package cms

import (
	"hiyuncms/models"
	"log"
)

type Column struct {
	Id   		int64 	`xorm:"pk BIGINT autoincr`
	Name 		string 	`xorm:"varchar(100)"`
	Url  		string 	`xorm:"varchar(200)"`
	ParentId 	int64 	`xorm:"BIGINT"`
	ShowFlag	int		`xorm:"INT"`
	OrderNum 	int 	`xorm:"int"`
	TemplatePath string `xorm:varchar(100)`
}

func  GetAllColumnsToSelect() *[]*Column{
	columnList := make([]*Column, 0)
	err := models.DbMaster.Table(Column{}).Find(&columnList)
	if err != nil {
		log.Printf("获取Column数据:%s", models.GetErrorInfo(err))
	}
	return &columnList
}

/**
获取所有栏目
 */
func GetAllColumnsByPage(page *models.PageRequest) *models.PageResponse{
	columnList := make([]*Column, 0)
	err := models.DbMaster.Table(Column{}).Limit(page.Rows, (page.Page - 1)* page.Rows).Find(&columnList)
	if err != nil {
		log.Printf("获取Column数据:%s", models.GetErrorInfo(err))
	}
	pageResponse := models.PageResponse{}
	pageResponse.Rows = columnList
	pageResponse.Page = page.Page
	pageResponse.Records ,_= models.DbMaster.Table(Column{}).Count(Column{})
	return &pageResponse
}

func  GetAllColumnsToShow() *[]*Column{
	columnList := make([]*Column, 0)
	err := models.DbMaster.Table(Column{}).Where("Show_Flag = 1").Find(&columnList)
	if err != nil {
		log.Printf("获取Column数据:%s", models.GetErrorInfo(err))
	}
	return &columnList
}


func init()  {
	err := models.DbMaster.Sync2( Column{})
	log.Println( "init table Column ", models.GetErrorInfo(err))
}
