package models

type UserRole struct {
	Id     int64        `xorm:"pk autoincr" json:"id"`
	OrgId  int64        `xorm:"bigint"      json:"orgId"`
	UserId int64        `xorm:"bigint"      josn:"userId"`
}



func init()  {
	err := engine.Sync(new(UserRole))
	println( "init UserRole struct ：", getErrorInfo(err))
}

func AddUserRole (ou UserRole)  {
	engine.Insert( ou )
}

func GetUserByRoleId(roleId int64, page Page)([]User,error) {
	userRoles := make( []UserRole, 0)
	err := engine.Where("role_id = ?", roleId).Limit((page.PageNum + 1)*page.PageSize, page.PageSize * page.PageNum).Find(&userRoles)
	if err != nil {
		users := make([]User, 0)
		for index, userRole := range userRoles {
			engine.Id(userRole.UserId).Get(&users[index])
		}
		return users,nil
	}
	return nil,err
}

func GetRoleByUserId(roleId int64, page Page)([]Role,error)   {
	userRoles := make( []UserRole, 0)
	err := engine.Where("user_id = ?", roleId).Limit((page.PageNum + 1)*page.PageSize, page.PageSize * page.PageNum).Find(&userRoles)
	if err != nil {
		roles := make([]Role, 0)
		for index, userRole := range userRoles {
			engine.Id(userRole.UserId).Get(&roles[index])
		}
		return roles,nil
	}
	return nil,err
}