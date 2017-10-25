package system

import (
	"log"
	"hiyuncms/models"
	"fmt"
)

type Resource struct {
	Id           int64  `xorm:"pk BIGINT autoincr" json:"id"`
	ResourceName string `xorm:"varchar(40) notnull"`
	ResourceCode string `xorm:"varchar(120) notnull unique"`
	ResourceUrl  string `xorm:"varchar(120)"`
	ParentId 	 int64 	`xorm:"BIGINT"`                         //
	OrderNo      int    `xorm:"INT"`     						//排序号
}

func init()  {
	err := models.DbMaster.Sync2( Resource{})
	log.Println( "init table Resource ", models.GetErrorInfo(err))
}


func GetResourceByPraentId(parentId int64)  []*Resource {
	resources := make( []*Resource, 0)
	models.DbSlave.Where(fmt.Sprintf("parent_id =%d", parentId)).Find( &resources)
	return resources
}



/**
根据parenId获得resources
 */
func GetResourceByPage(page *models.PageRequest, parentId int64) * models.PageResponse{
	resources := make([]*Resource, 0)
	//log.Printf("%v", page)
	err := models.DbSlave.
		Limit(page.Rows, (page.Page - 1) * page.Rows).
		Where(fmt.Sprintf("Parent_Id=%d",parentId)).
		Find(&resources)
	if err != nil {
		log.Printf("获取Parent数据:%s", models.GetErrorInfo(err))
	}
	records,_ :=  models.DbSlave.Where(fmt.Sprintf("Parent_Id=%d",parentId)).
		Count(Resource{})
	pageResponse := models.InitPageResponse(page, &resources, records)
	return  pageResponse
}
