package system

import (
	"log"
	"hiyuncms/models"
	"fmt"
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

func GetRolesByUserId( userId int64) []*Role{
	roles := make([]*Role, 0)
	err := models.DbSlave.Table(Role{}).Alias("r").
		Select("r.*").
		Join("INNER", []string{"hiyuncms_user_role","rr"}, fmt.Sprintf("rr.role_id=r.id and rr.user_id=%d",userId)).
		Find(&roles)
	if err != nil {
		log.Printf("根据user_id获取role数据:%s", models.GetErrorInfo(err))
	}
	return roles
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

