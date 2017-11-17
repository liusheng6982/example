package yy

import (
	"log"
	"hiyuncms/models"
)

type YyProject struct {
	Id            int64       `xorm:"pk BIGINT autoincr"`
	ProjectName   string      `xorm:"varchar(100) notnull"`
	ProjectNo	  string 	  `xorm:"varchar(50)"`
	ProjectType   string 	  `xorm:"varhcar(20)"`

	ExpiredDate   models.Date `xorm:"DateTime"`
	CompanyId	  int64       `xorm:"BIGINT"`
}

func init()  {
	err := models.DbMaster.Sync2( YyProject{})
	log.Println( "init table yy_project ", models.GetErrorInfo(err))
}
