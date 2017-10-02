package models

type RoleResource struct {
	Id         int64        `xorm:"pk autoincr" json:"id"`
	RoleId     int64        `xorm:"bigint"      json:"orgId"`
	ResourceId int64        `xorm:"bigint"      josn:"userId"`
}



func init()  {
	err := engine.Sync2(new(RoleResource))
	println( "init RoleResource struct ：", getErrorInfo(err))
}

func AddRoleResource (ou RoleResource)  {
	engine.Insert( ou )
}

func GetResourceByRoleId(roleId int64, page Page)([]Resource,error) {
	roleResources := make( []RoleResource, 0);
	err := engine.Where("role_id = ?", roleId).Limit((page.PageNum + 1)*page.PageSize, page.PageSize * page.PageNum).Find(&roleResources)
	if err != nil {
		resources := make([]Resource, 0)
		for index, userRole := range roleResources {
			engine.Id(userRole.ResourceId).Get(&resources[index])
		}
		return resources,nil
	}
	return nil,err
}

func GetRoleByResourceId(roleId int64, page Page)([]Role,error)   {
	roleResources := make( []RoleResource, 0)
	err := engine.Where("resource_id = ?", roleId).Limit((page.PageNum + 1)*page.PageSize, page.PageSize * page.PageNum).Find(&roleResources)
	if err != nil {
		roles := make([]Role, 0)
		for index, roleResource := range roleResources {
			engine.Id(roleResource.RoleId).Get(&roles[index])
		}
		return roles,nil
	}
	return nil,err
}