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


	log.Printf("form 提交的密码用户名,%s----%s\n", userName, passwd)
	admin := system.GetUserByUserName(userName)
	log.Printf( "%v\n", admin.LoginPassword )


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

func UserList(c *gin.Context){
	c.HTML(http.StatusOK, "userlist.html", gin.H{
		"bodyCss":"no-skin",
		"mainMenu" :"用户管理",
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
组织操作
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
		user.Id, _= strconv.ParseInt(id, 10, 64)
		_, err := models.DbMaster.Delete(&user)
		if err != nil {
			log.Printf("删除Org报错:%s\n",models.GetErrorInfo(err))
			c.String(http.StatusInternalServerError, "%s", "fail")
			return
		}
		c.String(http.StatusOK, "%s", "success")
	}
}
