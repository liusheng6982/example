package models

type OrgUser struct {
	Id     int64        `xorm:"pk autoincr" json:"id"`
	OrgId  int64        `xorm:"bigint"      json:"orgId"`
	UserId int64        `xorm:"bigint"      josn:"userId"`
}



func init()  {
	err := engine.Sync2(new(OrgUser))
	println( "init OrgUser struct ：", getErrorInfo(err))
}

func AddOrgUser (ou OrgUser)  {
	engine.Insert( ou )
}

func GetUserByOrgId(OrgId int64, page Page)([]User,error) {
	orgUsers := make( []OrgUser, 0);
	err := engine.Where("org_id = ?", OrgId).Limit((page.PageNum + 1)*page.PageSize, page.PageSize * page.PageNum).Find(&orgUsers)
	if err != nil {
		users := make([]User, 0)
		for index,orgUser := range orgUsers{
			engine.Id(orgUser.UserId).Get(&users[index])
		}
		return users,nil
	}
	return nil,err
}