package system

import (
	"log"
	"hiyuncms/models"
)

type Role struct {
	Id int64  `xorm:"pk BIGINT autoincr"`
	RoleName string `xorm:"varchar(25) notnull unique"`
}

func init()  {
	err := models.DbMaster.Sync2( Role{})
	log.Println( "init table role ", models.GetErrorInfo(err))
}