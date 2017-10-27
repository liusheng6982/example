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