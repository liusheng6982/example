
package yy

import (
	"log"
	"hiyuncms/models"
)

type YyCompanyRelation struct{
	Id              		int64      `xorm:"pk BIGINT autoincr"`
	HospitalId				int64  	   `xorm:"bigint"`
	SupplyId				int64      `xorm:"bigint"`
}

func init()  {
	err := models.DbMaster.Sync2( YyCompanyRelation{} )
	log.Println( "init table yy_company_relation", models.GetErrorInfo(err))
}

func GetCompanyIdsBySupplyId( supplyId int64 )([]*YyCompanyRelation){
	result := make( [] *YyCompanyRelation, 0)
	models.DbSlave.Table(YyCompanyRelation{}).Where("supply_id =?", supplyId).Find( &result )
	return result
}