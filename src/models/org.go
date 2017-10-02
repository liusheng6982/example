package models


type Org struct {
	Id   int64
	OrgName         string   `xorm:"varchar(50)" json:"orgName"`
	CreateTime 		Time	 `xorm:"created"       json:"createTime"`
	UpdateTime		Time	 `xorm:"updated"     json:"updateTime"`
	ParentId		int64    `xorm:"bigint"      json:"parentId"`
}



func init()  {
	err := engine.Sync(new(Org))
	println( "init Org struct ：", getErrorInfo(err))
}

func  AddOrg(userName string) error {
	user := Org{OrgName:userName}
	_, err := engine.Insert( user )
	if err != nil {
		println(err.Error())
	}
	return  err
}

