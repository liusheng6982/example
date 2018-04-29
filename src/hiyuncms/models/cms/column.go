package cms

import (
	"hiyuncms/models"
	"log"
	"fmt"
)

type Column struct {
	Id   		int64 	`xorm:"pk BIGINT autoincr" json:"id"`
	Name 		string 	`xorm:"varchar(100)"`
	Url  		string 	`xorm:"varchar(200)"`
	ParentId 	int64 	`xorm:"BIGINT"`
	ShowFlag	int		`xorm:"INT"`
	OrderNum 	int 	`xorm:"int"`
	TemplatePath string `xorm:"varchar(100)"`
}

/**
编辑文章关联的数据
 */
func  GetAllColumnsToSelect() *[]*Column{
	columnList := make([]*Column, 0)
	err := models.DbSlave.Table(Column{}).Find(&columnList)
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
	err := models.DbSlave.Table(Column{}).Limit(page.Rows, (page.Page - 1)* page.Rows).Find(&columnList)
	if err != nil {
		log.Printf("获取Column数据:%s", models.GetErrorInfo(err))
	}
	pageResponse := models.PageResponse{}
	pageResponse.Rows = columnList
	pageResponse.Page = page.Page
	pageResponse.Records ,_= models.DbMaster.Table(Column{}).Count(Column{})
	return &pageResponse
}

/**
发布时，用到的栏目（显示栏目位）
 */
func  GetAllColumnsToShow() *[]*Column{
	columnList := make([]*Column, 0)
	err := models.DbSlave.Table(Column{}).Where("Show_Flag = 1").Where("parent_id=0").OrderBy("order_num asc").Find(&columnList)
	if err != nil {
		log.Printf("获取Column数据:%s", models.GetErrorInfo(err))
	}
	return &columnList
}


/**
发布时，用到的子栏目（显示子栏目位）
 */
func  GetSubColumnsToShow(parentPath string) *[]*Column{
	parentColumn := GetColumnByPath(parentPath)
	columnList := make([]*Column, 0)
	err := models.DbSlave.Table(Column{}).
		Where("Show_Flag = 1").
		Where(fmt.Sprintf("parent_id=%d", parentColumn.Id)).OrderBy("order_num asc").Find(&columnList)
	if err != nil {
		log.Printf("获取Column数据:%s", models.GetErrorInfo(err))
	}
	return &columnList
}

/**
启动时用到的，router路径
 */
 func  GetAllColumnsToRoute() *[]*Column{
	columnList := make([]*Column, 0)
	err := models.DbSlave.Table(Column{}).Where("Show_Flag = 1").OrderBy("order_num asc").Find(&columnList)
	if err != nil {
		log.Printf("获取Column数据:%s", models.GetErrorInfo(err))
	}
	return &columnList
}

/**
根据路径，获取栏目对象
 */
func GetColumnByPath(path string) *Column {
	column := Column{}
	models.DbSlave.Table(Column{}).Where( fmt.Sprintf("Url = '%s'", path )).Get( &column )
	return &column
}


func init()  {
	err := models.DbMaster.Sync2( Column{})
	log.Println( "init table Column ", models.GetErrorInfo(err))
}
