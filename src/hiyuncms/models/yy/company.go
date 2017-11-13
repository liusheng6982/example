package yy

import (
	"log"
	"hiyuncms/models"
)

type YyCompany struct {
	Id              int64      `xorm:"pk BIGINT autoincr"`
	CompanyName     string     `xorm:"varchar(50) notnull unique"`

	CompanyType             int 		`xorm:"int"`
	CompanyProvince         string      `xorm:"varchar(20)"`
	CompanyCity		        string 	    `xorm:"varchar(20)"`
	CompanyAddress          string      `xorm:"varchar(120)"`
	CompanyBusinessLicense  string      `xorm:"varchar(120)"`
	CompanyImage 			string      `xorm:"varchar(100)"`
}


func init()  {
	err := models.DbMaster.Sync2( YyCompany{})
	log.Println( "init table yy_company", models.GetErrorInfo(err))
}

func CompanyReg(company *YyCompany, user *YyUser) error{

	session := models.DbMaster.NewSession()

	user.CompanyId =  company.Id

	defer session.Close()
	// add Begin() before any action
	err := session.Begin()
	if err != nil {
		return err
	}

	_, err1 := session.Insert( company )
	if err1 != nil {
		session.Rollback()
		return err1
	}
	_, err2 := session.Insert( user )
	if err2 != nil{
		session.Rollback()
		return err2
	}

	session.Commit()
	return nil
}


