package yy

import (
	"log"
	"hiyuncms/models"
)

func GetInviteTenderBySupplyId(size int, supplyId int64 )  []*YyPorject{
	hospitals := GetCompanyIdsBySupplyId( supplyId )
	hospitalIds := make( []int64, 0)
	for _,hospital:= range hospitals{
		hospitalIds = append(hospitalIds, hospital.HospitalId)
	}
	inviteTenderList := make([]*YyPorject, 0)
	err := models.DbSlave.Table(YyPorject{}).Limit(size,0).
		Where("project_type=1").Where("company_id in ？or auth=?", hospitalIds, "0").Find(&inviteTenderList)
	if err != nil {
		log.Printf("获取YyInviteTender数据:%s", models.GetErrorInfo(err))
	}
	return inviteTenderList
}

/**
获取这个公司下的所有项目
 */
func GetInviteTenderByCompanyId(size int, companyId int64 )[]*YyPorject{
	inviteTenderList := make([]*YyPorject, 0)
	err := models.DbSlave.Table(YyPorject{}).Limit(size,0).
		Where("project_type=1").Where("company_id = ?", companyId).Find(&inviteTenderList)
	if err != nil {
		log.Printf("获取YyInviteTender数据:%s", models.GetErrorInfo(err))
	}
	return inviteTenderList
}


func GetPurchaseBySupplyId(size int, supplyId int64 )  []*YyPorject{
	hospitals := GetCompanyIdsBySupplyId( supplyId )
	hospitalIds := make( []int64, 0)
	for _,hospital:= range hospitals{
		hospitalIds = append(hospitalIds, hospital.HospitalId)
	}
	inviteTenderList := make([]*YyPorject, 0)
	err := models.DbSlave.Table(YyPorject{}).Limit(size,0).
		Where("project_type=1").Where("company_id in ？or auth=?", hospitalIds, "0").Find(&inviteTenderList)
	if err != nil {
		log.Printf("获取YyInviteTender数据:%s", models.GetErrorInfo(err))
	}
	return inviteTenderList
}

/**
获取这个公司下的所有项目
 */
func GetPurchaseByCompanyId(size int, companyId int64 )[]*YyPorject{
	inviteTenderList := make([]*YyPorject, 0)
	err := models.DbSlave.Table(YyPorject{}).Limit(size,0).
		Where("project_type=1").Where("company_id = ?", companyId).Find(&inviteTenderList)
	if err != nil {
		log.Printf("获取YyInviteTender数据:%s", models.GetErrorInfo(err))
	}
	return inviteTenderList
}


