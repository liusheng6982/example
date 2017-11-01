package system

import (
	"log"
	"hiyuncms/models"
	"fmt"
)

type User struct {
	Id            int64  `xorm:"pk BIGINT autoincr" json:"id"`
	Name 	      string `xorm:"varchar(25)"`
	Phone         string `xomr:"varchar(20)"`
	LoginName     string `xorm:"varchar(25) notnull unique"`
	LoginPassword string `xorm:"varchar(64) null" json:"-"`
}

func GetUserById(userId int64 ) *User  {
	admin := User{ Id:userId }
	models.DbSlave.Get(&admin)
	return &admin
}

func GetUserByUserName(userName string ) *User {
	if  userName == "" {
		return &User{}
	}
	admin := User{ LoginName:userName }
	models.DbSlave.Get(&admin)
	return &admin
}

func init()  {
	err := models.DbMaster.Sync2( User{})
	log.Println( "init table user ", models.GetErrorInfo(err))
	adminUser := User{Id:1, LoginName:"admin", LoginPassword:"8211c2dc6aa7cf474144ab9bfa73893e"}
	_,err = models.DbMaster.Insert( &adminUser )
	if err != nil {
		log.Println("管理员账号已存在： ", models.GetErrorInfo(err))
	}
}

/**
保存用户
 */
func SaveUser(user * User, orgId int64, orderNo int)  {
	_, err := models.DbMaster.Insert( user  )
	if err != nil {
		log.Printf("保存用户报错：%s\n", err.Error())
	}
	orgUser := OrgUser{
		OrgId:orgId,
		OrderNo:orderNo,
		UserId:user.Id,
	}
	models.DbMaster.Insert( orgUser  )
}

func DelUser(userId int64)  {
	article := User{Id:userId}
	models.DbMaster.Delete( &article )
	models.DbMaster.Delete(&OrgUser{UserId:userId})
}

/**
根据组织获得用户
 */
func GetUsersByOrg(page *models.PageRequest, orgId int64) * models.PageResponse{
	users := make([]*User, 0)
	//log.Printf("%v", page)
	err := models.DbSlave.Table(User{}).Alias("u").
		Select("u.*").
		Limit(page.Rows, (page.Page - 1) * page.Rows).
		Join("INNER", []string{"hiyuncms_org_user","ou"}, fmt.Sprintf("u.id=ou.user_id and ou.org_id=%d",orgId)).
		Find(&users)
	if err != nil {
		log.Printf("通过Column的URL获取Article数据:%s", models.GetErrorInfo(err))
	}
	records,_ :=  models.DbSlave.Table(User{}).Alias("u").
		Join("INNER", []string{"hiyuncms_org_user","ou"}, fmt.Sprintf("u.id=ou.user_id and ou.org_id=%d",orgId)).
		Count(User{})

	pageResponse := models.InitPageResponse(page, &users, records)
	return  pageResponse

}