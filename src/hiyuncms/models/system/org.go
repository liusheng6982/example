package system

import (
	"log"
	"hiyuncms/models"
	"fmt"
)

type Org struct {
	Id 			int64   `xorm:"pk BIGINT autoincr" json:"id"`
	OrgName 	string  `xorm:"varchar(40) notnull"`
	OrgCode		string  `xorm:"varchar(40)"`
	ParentId 	int64 	`xorm:"BIGINT"`
	OrderNo		int		`xorm:"INT"`
}

func init()  {
	err := models.DbMaster.Sync2( Org{})
	log.Println( "init table Org ", models.GetErrorInfo(err))
}

func GetSubOrgByPraentId(parentId int64)  []*Org {
	orgs := make( []*Org, 0)
	models.DbSlave.Where(fmt.Sprintf("parent_id =%d", parentId)).Find( &orgs )
	return orgs
}

func GetSubOrgByPage(parentId int64, page *models.PageRequest) *models.PageResponse {
	orgs := make( []*Org, 0)

	err := models.DbSlave.Table(Org{}).Where(fmt.Sprintf("parent_id =%d", parentId)).Limit(page.Rows, (page.Page - 1)* page.Rows).Find(&orgs)
	if err != nil {
		log.Printf("查询组织分页报错%s", err.Error())
	}
	records ,_ := models.DbSlave.Table(Org{}).Where(fmt.Sprintf("parent_id =%d", parentId)).Count(Org{})
	pageResponse := models.InitPageResponse(page, &orgs, records)
	return pageResponse
}