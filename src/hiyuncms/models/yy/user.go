package yy

import (
	"log"
	"hiyuncms/models"
)

type YyUser struct {
	Id            	  int64  `xorm:"pk BIGINT autoincr" json:"id"`
	UserName 	      string `xorm:"varchar(25)"`
	UserPhone         string `xomr:"varchar(20) unique"`
	//UserLoginName         string `xorm:"varchar(25) notnull unique"`
	UserPassword      string `xorm:"varchar(64) null" json:"-"`
	CompanyId 		  int64  `xorm:"BIGINT"`
	CompanyName		  string
}


func init()  {
	err := models.DbMaster.Sync2( YyUser{})
	log.Println( "init table yy_user ", models.GetErrorInfo(err))
}

func GetUserByPhone(phone string ) *YyUser {
	if  phone == "" {
		return &YyUser{}
	}
	user := YyUser{ UserPhone: phone}
	models.DbSlave.Get(&user)
	return &user
}


func GetVipUsers(page *models.PageRequest )*models.PageResponse{
	result := make( [] *YyUser, 0)
	models.DbSlave.Table(YyUser{}).Alias("u").
		//Join("LEFT OUTER", []string{"hiyuncms_yy_company", "c"}, "c.id = u.company_id").
		Limit(page.Rows, (page.Page - 1)* page.Rows).
		Find(&result)
	for _, v := range result{
		compnayInfo := YyCompany{}
		models.DbSlave.Id( v.CompanyId ).Get(&compnayInfo)
		v.CompanyName = compnayInfo.CompanyName
	}
	records ,_ := models.DbSlave.Table(YyUser{}).Alias("u").
		//Join("LEFT OUTER", []string{"hiyuncms_yy_company", "c"}, "c.id = u.company_id").
		Count()
	pageResponse := models.InitPageResponse(page, &result, records)
	return pageResponse
}




