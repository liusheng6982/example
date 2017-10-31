package yy

import (
	"log"
	"hiyuncms/models"
)

type User struct {
	Id            int64  `xorm:"pk BIGINT autoincr" json:"id"`
	Name 	      string `xorm:"varchar(25)"`
	Phone         string `xomr:"varchar(20)"`
	LoginName     string `xorm:"varchar(25) notnull unique"`
	LoginPassword string `xorm:"varchar(64) null" json:"-"`
}

func (u * User) TableName() string {
	return "yy_user"
}

func init()  {
	err := models.DbMaster.Sync2( Project{})
	log.Println( "init table yy_user ", models.GetErrorInfo(err))
}

