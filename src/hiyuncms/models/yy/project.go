package yy

import (
	"log"
	"hiyuncms/models"
)

type YyProject struct {
	Id            int64       `xorm:"pk BIGINT autoincr"`
	CompanyName          string      `xorm:"varchar(50) notnull"`
    Content       string      `xorm:"varhcar(2000)"`
	StartDate     models.Time `xorm:"DateTime"`
	ExpiredDate   models.Date `xorm:"DateTime"`
	CompanyId	  int64       `xorm:"BIGINT"`
}

func init()  {
	err := models.DbMaster.Sync2( YyProject{})
	log.Println( "init table yy_project ", models.GetErrorInfo(err))
}
