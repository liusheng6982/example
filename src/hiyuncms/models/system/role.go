package system

import (
	"log"
	"hiyuncms/models"
)

type Role struct {
	Id int64  `xorm:"pk BIGINT autoincr" json:"id"`
	RoleName string `xorm:"varchar(40) notnull"`
	RoleCode string `xorm:"varchar(25) notnull unique"`
}

func init()  {
	err := models.DbMaster.Sync2( Role{})
	log.Println( "init table role ", models.GetErrorInfo(err))
}

/**
获得角色分页
 */
func GetAllRolesByPage(page *models.PageRequest) *models.PageResponse{
	roleList := make([]*Role, 0)
	err := models.DbSlave.Table(Role{}).Limit(page.Rows, (page.Page - 1)* page.Rows).Find(&roleList)
	if err != nil {
		log.Printf("获取Column数据:%s", models.GetErrorInfo(err))
	}
	records,_ := models.DbMaster.Table(Role{}).Count(Role{})
	pageResponse := models.InitPageResponse(page, roleList, records)

	return pageResponse
}

/**
获得角色数据
 */
func GetAllRoles() []*Role{
	roleList := make([]*Role, 0)
	err := models.DbSlave.Table(Role{}).Find(&roleList)
	if err != nil {
		log.Printf("获取Role数据:%s", models.GetErrorInfo(err))
	}
	return roleList
}