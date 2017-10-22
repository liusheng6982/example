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
)

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

	m5 := md5.New()
	m5.Write([]byte(passwd))
	m5.Write([]byte(string("hihi")))
	st := m5.Sum(nil)
	passwdMd5 := fmt.Sprintf("%s", hex.EncodeToString(st))
	log.Printf("111111111111111111111111111111111111=%s\n", passwdMd5)

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
