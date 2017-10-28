package system

import (
	"log"
	"hiyuncms/models"
)

type RoleResource struct {
	Id              int64     `xorm:"pk autoincr"`
	RoleId		    int64     `xorm:"bigint"`
	ResourceId      int64     `xorm:"bigint"`
}

func init()  {
	err := models.DbMaster.Sync2( RoleResource{})
	log.Println( "init table RoleResource ", models.GetErrorInfo(err))
}

func IsSelectResourceByRoleId(roleId, resourceId int64) bool {
	releResource := RoleResource{RoleId:roleId,ResourceId:resourceId}
	has, err := models.DbSlave.Table(RoleResource{}).Get( &releResource)
	if err != nil {
		log.Printf("角色与资源是关联查询报错:%s\n", models.GetErrorInfo(err))
	}
	return  has
}

func RoleResourceSave(roleId int64, resourceIds [] int64){
	roleResource := RoleResource{RoleId: roleId}
	models.DbMaster.Delete(&roleResource)
	roleResourcers := make([]*RoleResource,len( resourceIds ))
	for k,v := range resourceIds{
		roleResourcers[k] = &RoleResource{RoleId: roleId, ResourceId:v}
	}
	models.DbMaster.Insert( roleResourcers )
}

