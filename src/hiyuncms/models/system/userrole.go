package system

import (
	"log"
	"hiyuncms/models"
)

type UserRole struct {
	Id              int64     `xorm:"pk autoincr"`
	UserId		    int64     `xorm:"bigint"`
	RoleId          int64     `xorm:"bigint"`
}

func init()  {
	err := models.DbMaster.Sync2( UserRole{})
	log.Println( "init table UserRole ", models.GetErrorInfo(err))
}

func IsSelectRoleByUserId(userId, roleId int64) bool {
	userRole := UserRole{RoleId:roleId,UserId:userId}
	has, err := models.DbSlave.Table(UserRole{}).Get( &userRole)
	if err != nil {
		log.Printf("用户与角色是关联查询报错:%s\n", models.GetErrorInfo(err))
	}
	return  has
}

func UserRoleSave(userId int64, roleIds [] int64){
	userRole := UserRole{UserId: userId}
	models.DbMaster.Delete(&userRole)
	userRoles := make([]*UserRole,len(roleIds))
	for k,v := range roleIds {
		userRoles[k] = &UserRole{UserId: userId, RoleId:v}
	}
	models.DbMaster.Insert(userRoles)
}

