package backend

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/contrib/sessions"
	"hiyuncms/models/system"
	"hiyuncms/controllers"
	"net/http"
	"log"
	"encoding/json"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"hiyuncms/models"
	"strconv"
	"strings"
)

func md5str(passwd string) string  {
	m5 := md5.New()
	m5.Write([]byte(passwd))
	m5.Write([]byte(string("hihi")))
	st := m5.Sum(nil)
	passwdMd5 := fmt.Sprintf("%s", hex.EncodeToString(st))
	return passwdMd5
}

/**
后台用户登录
 */
func  UserLogin(c * gin.Context)  {
	vcode := c.PostForm("vcode")
	session := sessions.Default(c)
	sessionCode := session.Get( controllers.BACK_CAPTCHA_SESSION )
	log.Printf("%s----------%v", vcode, sessionCode)
	if vcode != sessionCode {
		c.HTML(http.StatusOK, "login.html",gin.H{
			"msg":"验证码错误！",
			"bodyCss": "login-layout",
		})
		return
	}
	userName := c.PostForm("Username")
	passwd := c.PostForm("Password")
	admin := system.GetUserByUserName(userName)


	passwdMd5 := md5str(passwd)

	/*
	m5 = md5.New()
	m5.Write([]byte(admin.LoginPassword))
	m5.Write([]byte(string("hihi")))
	st = m5.Sum(nil)
	loginPasswdMd5 := fmt.Sprintf("%s", hex.EncodeToString(st))
	*/

	if admin.LoginPassword == passwdMd5 {
	//if admin.LoginPassword == passwd {
		log.Printf("登录成功！\n")
		bus := controllers.BackendUserSession{Name:admin.LoginName, Id:admin.Id}
		session := sessions.Default(c)
		jsonBytes,_ := json.Marshal(bus)
		session.Set(controllers.BACK_USER_SESSION,  string(jsonBytes) )
		session.Save()

		c.Redirect(http.StatusFound, "/index")
	} else{
		c.HTML(http.StatusOK, "login.html",gin.H{
			"msg":"用户名不存在或密码错误！",
			"bodyCss": "login-layout",
		})
	}
}

func ChangePassword(c *gin.Context)  {
	user := controllers.GetSessionUser(c)
	log.Printf("========%v", user)
	log.Printf("========%s", user.LoginName)
	admin := system.GetUserById( user.Id )

	log.Printf("========%s", admin.LoginPassword)

	currentPassword := c.PostForm("currentPassword")
	newPassword := c.PostForm("newPassword")

	log.Printf("after========%s",md5str (currentPassword ))
	if admin.LoginPassword != md5str (currentPassword ){
		c.JSON(http.StatusOK, "当前密码错误")
	}else{
		admin.LoginPassword = md5str( newPassword )
		_, err := models.DbMaster.Id( admin.Id ).Update( admin )
		if err == nil {
			c.JSON(http.StatusOK, true)
		} else{
			c.JSON(http.StatusOK, fmt.Sprintf("请与管理员联系！%s",err.Error()))
		}
	}


}

func UserList(c *gin.Context){
	c.HTML(http.StatusOK, "userlist.html", gin.H{
		"bodyCss":"no-skin",
		"mainMenu" :"用户管理",
		"user":controllers.GetSessionUser(c),
	})
}

func UserListData(c *gin.Context){
	page := models.PageRequest{}
	c.Bind( &page )
	orgIdStr,_:= c.GetPostForm("orgId")
	orgId,_:= strconv.ParseInt(orgIdStr, 10, 64)
	responsePage := system.GetUsersByOrg(&page, orgId)
	c.JSON(http.StatusOK, responsePage)
}

/*
用户操作
 */
func UserEdit(c * gin.Context){
	user := system.User{}
	c.Bind( &user)
	oper, _ := c.GetPostForm("oper")
	if "edit" == oper {
		id, _:= c.GetPostForm("id")
		user.Id, _= strconv.ParseInt(id, 10, 64)
		_, err := models.DbMaster.ID(user.Id).Update(&user)
		if err != nil {
			log.Printf("更新Org报错:%s\n",models.GetErrorInfo(err))
		}
	}else if "add" == oper {
		passwordMd5 := md5str("000000")
		user.LoginPassword = passwordMd5
		orgIdStr,_ := c.GetPostForm("orgId")
		orgId,_:= strconv.ParseInt(orgIdStr, 10, 64)
		system.SaveUser( &user, orgId, 0)

		c.String(http.StatusOK, "%s", "success")
	} else if "del" == oper{
		id, _:= c.GetPostForm("id")
		userId, _ := strconv.ParseInt(id, 10, 64)
		system.DelUser(userId)
		c.String(http.StatusOK, "%s", "success")
	}
}

/**
设置用户角色面板
 */
func UserRoles (c * gin.Context){
	userIdStr := c.Query("userId")
	userId, _ :=  strconv.ParseInt(userIdStr, 10, 64)

	roles := system.GetAllRoles()
	options := ""
	for _,role := range roles{
		if system.IsSelectRoleByUserId(userId, role.Id) {
			options = fmt.Sprintf("%s<option value='%d' selected='selected'> %s </option>", options, role.Id,role.RoleName)
		} else {
			options = fmt.Sprintf("%s<option value='%d'>%s</option>", options, role.Id,role.RoleName)
		}
	}
	c.JSON(http.StatusOK, options)
}

func UserRolesSave(c * gin.Context)  {

	userIdStr := c.Query("userId")
	userId, _ :=  strconv.ParseInt(userIdStr, 10, 64)
	roleIdsStr := c.Query("roleIds")
	roleIds := strings.Split(roleIdsStr,",")
	roleIdsInt := make([]int64, len(roleIds))
	for k,v := range  roleIds{
		roleIdsInt[k],_ = strconv.ParseInt(v, 10,64)
	}
	system.UserRoleSave(userId, roleIdsInt)
	c.JSON(http.StatusOK, gin.H{
		"result":"success",
	})
}
