package yy

import (
	"log"
	"hiyuncms/models"
)

type Project struct {
	Id            int64       `xorm:"pk BIGINT autoincr"`
	Name          string      `xorm:"varchar(50) notnull"`
    Content       string      `xorm:"varhcar(2000)"`
	StartDate     models.Time `xorm:"DateTime"`
	ExpiredDate   models.Date `xorm:"DateTime"`
	CompanyId	  int64       `xorm:"BIGINT"`
}

func (p * Project) TableName() string {
	return "yy_project"
}

func init()  {
	err := models.DbMaster.Sync2( Project{})
	log.Println( "init table yy_project ", models.GetErrorInfo(err))
}