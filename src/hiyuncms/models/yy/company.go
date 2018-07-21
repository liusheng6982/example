package yy

import (
	"log"
	"hiyuncms/models"
	"strings"
)

type YyCompany struct {
	Id              		int64      `xorm:"pk BIGINT autoincr" json:"id"`
	CompanyName     		string     `xorm:"varchar(50) notnull unique"`

	CompanyType             string 		`xorm:"varchar(50)"`  //1:医院 2：供应商
	CompanyProvince         string      `xorm:"varchar(30)"`
	CompanyCity		        string 	    `xorm:"varchar(30)"`
	CompanyAddress          string      `xorm:"varchar(120)"`
	CompanyBusinessLicense  string      `xorm:"varchar(120)"`
	CompanyImage 			string      `xorm:"varchar(100)"`
	CompanyVip				string      `xorm:"varchar(10)"`
	VipExpired 				models.Date `xorm:"DateTime"`
	VipLevel				int64 		`xorm:"INT"`
	LogoImage 				string      `xorm:"varchar(200)"`

}


func init()  {
	err := models.DbMaster.Sync2( YyCompany{})
	log.Println( "init table yy_company", models.GetErrorInfo(err))
}

func GetById(id int64)  *YyCompany {
	company := YyCompany{}
	models.DbSlave.Id(id).Get(&company)
	return &company
}

func GetHospital() []* YyCompany{
	companies := make([]* YyCompany, 0)
	models.DbSlave.Table(YyCompany{}).Where("Company_Type=1").Find(&companies)
	return companies
}

func GetAllSupplies() []* YyCompany{
	companies := make([]* YyCompany, 0)
	models.DbSlave.Table(YyCompany{}).Where("Company_Type=2").Find(&companies)
	return companies
}

func UpdateCompany(company * YyCompany){
	_, err:=models.DbMaster.Id( company.Id ).Update( company )
	if err != nil {
		log.Printf("更新公司信息报错：%s\n%+v", err.Error(),company )
	}
}

func CompanyReg(company *YyCompany, user *YyUser) (error,string){

	msg := "success"
	session := models.DbMaster.NewSession()
	defer session.Close()
	//add Begin() before any action
	err := session.Begin()
	if err != nil {
		return err, ""
	}

	_, err1 := session.Insert( company )
	if err1 != nil {
		session.Rollback()
		if strings.Contains(err1.Error(), "Error 1062"){
			msg ="企业名称不能重复！"
		} else{
			msg = err1.Error()
		}
		return err1,msg
	}

	user.CompanyId =  company.Id

	_, err2 := session.Insert( user )
	if err2 != nil{
		session.Rollback()
		if strings.Contains(err2.Error(), "Error 1062"){
			msg ="用户手机不能重复！"
		}else{
			msg = err2.Error()
		}
		return err2,msg
	}

	session.Commit()
	return nil, msg
}

func JudgeVIP(userCompanyId,projectCompanyId int64)(int64){
	is := IsSelectSupplyByHospitalId(userCompanyId, projectCompanyId)
	if is {
		return 1
	}else{
		return 0 //不是供应商
	}
}




