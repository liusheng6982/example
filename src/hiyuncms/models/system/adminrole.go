package system

import (
	"log"
	"hiyuncms/models"
)

type AdminRole struct {
	Admin `xorm:"extends"`
	Role  `xorm:"extends"`
}

func init()  {
	err := models.DbMaster.Sync2( AdminRole{})
	log.Println( "init table AdminRole ", models.GetErrorInfo(err))
}