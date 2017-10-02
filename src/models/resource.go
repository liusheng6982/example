package models


type Resource struct {
	Id   	  		int64   	 `xorm:"pk autoincr" json:"id"`
	ResourceName 	string  	 `xorm:"varchar(25)" json:"userName"`
	ResourceUrl     string       `xorm:"varchar(50)" json:"loginName"`
	CreateTime 		Time	     `xorm:"created"     json:"createTime"`
	UpdateTime		Time	     `xorm:"updated"     json:"updateTime"`
}


func init()  {
	err := engine.Sync2(new(Resource))
	println( "init Resource struct ：", getErrorInfo(err))
}

func  AddResource(resource * Resource) error {
	_, err := engine.Insert( resource )
	if err != nil {
		println(err.Error())
	}
	return  err
}

func GetAllResource(page * Page) ([]Resource,error) {
	resources := make([]Resource, 0)
	err := engine.Limit(page.PageSize * (page.PageNum + 1), page.PageSize * page.PageNum).Find(&resources)
	return resources,err
}

