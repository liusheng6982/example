package system

import (
	"log"
	"hiyuncms/models"
)

type OrgUser struct {
	Id              int64     `xorm:"pk autoincr"`
	OrgId		    int64     `xorm:"bigint"`
	UserId          int64     `xorm:"bigint"`
	OrderNo		    int		  `xorm:"INT"`
}

func init()  {
	err := models.DbMaster.Sync2( OrgUser{})
	log.Println( "init table ColumnArticle ", models.GetErrorInfo(err))
}