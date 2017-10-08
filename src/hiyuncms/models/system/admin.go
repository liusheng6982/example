package system

import (
	"log"
	"hiyuncms/models"
)

type Admin struct {
	Id int64  `xorm:"pk BIGINT autoincr"`
	LoginName string `xorm:"varchar(25) notnull unique"`
	LoginPassword string `xorm:"varchar(64) null"`
}

func GetUserByUserName(userName string ) Admin{
	if  userName == "" {
		return Admin{}
	}
	admin := Admin{ LoginName:userName }
	models.DbMaster.Get(&admin)
	return admin
}

func init()  {
	err := models.DbMaster.Sync2( Admin{})
	log.Println( "init table admin ", models.GetErrorInfo(err))
}