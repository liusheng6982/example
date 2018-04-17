
package yy

import (
	"log"
	"hiyuncms/models"
	"fmt"
	"io/ioutil"
	"net/http"
	"hiyuncms/config"
	"strings"
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

func IsSelectSupplyByHospitalId(hospitalId, supplyId int64) bool {
	userRole := YyCompanyRelation{HospitalId:hospitalId,	SupplyId:supplyId}
	has, err := models.DbSlave.Table(YyCompanyRelation{}).Get( &userRole)
	if err != nil {
		log.Printf("医院与供应商是关联查询报错:%s\n", models.GetErrorInfo(err))
	}
	return  has
}

func IsSupply(supplyId int64) bool{
	log.Printf("是否是已有供应商:%d",supplyId)
	userRole := YyCompanyRelation{SupplyId:supplyId}
	has, err := models.DbSlave.Table(YyCompanyRelation{}).Get( &userRole)
	if err != nil {
		log.Printf("是否是已有供应商:%s\n", models.GetErrorInfo(err))
	}
	return  has
}

func GetCompanyIdsBySupplyId( supplyId int64 )([]*YyCompanyRelation){
	result := make( [] *YyCompanyRelation, 0)
	models.DbSlave.Table(YyCompanyRelation{}).Where("supply_id =?", supplyId).Find( &result )
	return result
}

func GetSuppliesByHospitalId( hospitalId int64, page *models.PageRequest )*models.PageResponse{
	result := make( [] *YyCompany, 0)
	log.Printf("hospital_id=%d", hospitalId)
	models.DbSlave.Table(YyCompany{}).Alias("com").
		Join("INNER", []string{"hiyuncms_Yy_Company_Relation", "re"}, "com.id = re.supply_id").
			Limit(page.Rows, (page.Page - 1)* page.Rows).
			Where("re.hospital_id=  ?", hospitalId).Find(&result)
	records ,_ := models.DbSlave.Table(YyCompany{}).Alias("com").
		Join("INNER", []string{"hiyuncms_Yy_Company_Relation", "re"}, "com.id = re.supply_id").
		Where("re.hospital_id=  ?", hospitalId).Count()
	pageResponse := models.InitPageResponse(page, &result, records)

	return pageResponse
}

func HospitalSupplySave(hospitalId int64, supplies [] int64){
	log.Printf("supplies===:%v", supplies)
	companyRelation := make([]*YyCompanyRelation, 0)
	models.DbSlave.Table(YyCompanyRelation{}).Where("hospital_id = ?", hospitalId ).Find(&companyRelation)
	log.Printf("info  -2-2-2-2-2-=================：%+v", companyRelation)
	info := make(map[int64]int64)
	log.Printf("info  -1-1-1-1-1-=================：%v", info)
	for _, v := range companyRelation{
		info[v.SupplyId] = v.SupplyId
	}

	log.Printf("info  000000=================：%v", info)

	for _, v := range supplies {
		delete(info, v)
		if v == 0 {
			continue
		}
		isExist := false
		//models.DbSlave.Table(YyCompanyRelation{}).Where("hospital_id = ?", hospitalId ).And("supply_id = ?", v).Exist(&isExist)
		save := YyCompanyRelation{HospitalId:hospitalId, SupplyId:v}
		isExist,_ = models.DbSlave.Table(YyCompanyRelation{}).Get(&save)

		log.Printf("是否存在：%v", isExist)

		if !isExist{


			ha,err  := models.DbMaster.Insert( &save )
			if err != nil {
				log.Printf("保存医院与供应商关系出错%s", err.Error())
			}
			go HospitalSupplySync(config.GetValue("sync.supply.guoxin.url"), hospitalId, v, "1")
			//go HospitalSupplySync(config.GetValue("sync.supply.chuanyiwang.url"), hospitalId, v, "1")
			log.Printf("save=============:%d", ha )
		}
	}

	log.Printf("info  11111111=================：%v", info)

	for _,v := range info{
		if v == 0 {
			continue
		}
		go HospitalSupplySync(config.GetValue("sync.supply.guoxin.url"), hospitalId, v, "0")
		//go HospitalSupplySync(config.GetValue("sync.supply.chuanyiwang.url"), hospitalId, v, "0")
		delData := YyCompanyRelation{HospitalId:hospitalId, SupplyId:v}
		models.DbMaster.Delete(&delData)
	}
}

func HospitalSupplySync(url string, hospitalId,supplyId int64, flag string ){ //医院供应商关系同步
	var getUrl string
	if strings.Contains(url, "?"){
		getUrl = fmt.Sprintf("%s&hospitalId=%d&supplyId=%d&flag=%s",
			url, hospitalId, supplyId, flag)
	} else{
		getUrl = fmt.Sprintf("%s?hospitalId=%d&supplyId=%d&flag=%s",
			url, hospitalId, supplyId, flag)
	}

	log.Printf("同步的url:%s", getUrl )

	res, err := http.Get(getUrl)
	log.Printf("!!!!!!!!!!!!!!!!!!%s", err)
	if err == nil {
		body, err1 := ioutil.ReadAll(res.Body)
		log.Printf("%s\n", body)
		if err1 != nil {
			log.Printf("-----err1=%s\n", err1)
		}
	} else {
		log.Printf("------err=%s\n", err)
	}
}

