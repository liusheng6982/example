package system

import (
	"log"
	"hiyuncms/models"
)

type Role struct {
	Id int64  `xorm:"pk BIGINT autoincr"`
	RoleName string `xorm:"varchar(40) notnull"`
	RoleCode string `xorm:"varchar(25) notnull unique"`
}

func init()  {
	err := models.DbMaster.Sync2( Role{})
	log.Println( "init table role ", models.GetErrorInfo(err))
}


func GetAllRolesByPage(page *models.PageRequest) *models.PageResponse{
	columnList := make([]*Role, 0)
	err := models.DbSlave.Table(Role{}).Limit(page.Rows, (page.Page - 1)* page.Rows).Find(&columnList)
	if err != nil {
		log.Printf("获取Column数据:%s", models.GetErrorInfo(err))
	}
	records,_ := models.DbMaster.Table(Role{}).Count(Role{})
	pageResponse := models.InitPageResponse(page, columnList, records)

	return pageResponse
}