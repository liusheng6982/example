package yy

import (
	"log"
	"hiyuncms/models"
)

type YyUser struct {
	Id            	  int64  `xorm:"pk BIGINT autoincr" json:"id"`
	UserName 	      string `xorm:"varchar(25)"`
	UserPhone             string `xomr:"varchar(20)"`
	//UserLoginName         string `xorm:"varchar(25) notnull unique"`
	UserPassword     string `xorm:"varchar(64) null" json:"-"`
	CompanyId 		  int64  `xorm:"BIGINT"`
}


func init()  {
	err := models.DbMaster.Sync2( YyUser{})
	log.Println( "init table yy_user ", models.GetErrorInfo(err))
}



