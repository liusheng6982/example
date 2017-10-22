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
	models.DbSlave.Get(&admin)
	return admin
}

func init()  {
	err := models.DbMaster.Sync2( Admin{})
	log.Println( "init table admin ", models.GetErrorInfo(err))
	adminUser := Admin{Id:1, LoginName:"admin", LoginPassword:"8211c2dc6aa7cf474144ab9bfa73893e"}
	_,err = models.DbMaster.Insert( &adminUser )
	if err != nil {
		log.Println("管理员账号已存在： ", models.GetErrorInfo(err))
	}
}