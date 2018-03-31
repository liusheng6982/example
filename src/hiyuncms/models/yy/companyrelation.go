
package yy

import (
	"log"
	"hiyuncms/models"
)

type YyCompanyRelation struct{
	Id              		int64      `xorm:"pk BIGINT autoincr"`
	HospitailId				int64  		`xorm:"bigint"`
	SupllyId				int64      `xorm:"bigint"`
}

func init()  {
	err := models.DbMaster.Sync2( YyCompanyRelation{} )
	log.Println( "init table yy_company_relation", models.GetErrorInfo(err))
}