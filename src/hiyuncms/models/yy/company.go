package yy

import (
	"log"
	"hiyuncms/models"
)

type Company struct {
	Id              int64       `xorm:"pk BIGINT autoincr"`
	Name            string      `xorm:"varchar(50) notnull"`

	BusinessType    int 		`xorm:"int"`
	Province        string      `xorm:"varchar(20)"`
	City		    string 	    `xorm:"varchar(20)"`
	Address         string      `xorm:"varchar(120)"`
	BusinessLicense string      `xorm:"varchar(120)"`
}

func (c * Company) TableName() string {
	return "yy_company"
}

func init()  {
	err := models.DbMaster.Sync2( Project{})
	log.Println( "init table yy_company", models.GetErrorInfo(err))
}


