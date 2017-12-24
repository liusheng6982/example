package frontend

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"hiyuncms/models/yy"
	"hiyuncms/controllers/backend"
	"log"
	"time"
	"fmt"
	"os"
	"io"
	"github.com/gin-gonic/contrib/sessions"
	"encoding/json"
	"github.com/satori/go.uuid"
	"hiyuncms/config"
)

func UserLoginShow(c * gin.Context)  {
	c.HTML(http.StatusOK, "login.html", gin.H{
		"path":"",
		"sessionInfo":GetSessionInfo(c),
	})
}

func UserLogin(c * gin.Context)  {
	vcode := c.PostForm("vcode")
	session := sessions.Default(c)
	sessionCode := session.Get( FRONT_CAPTCHA_SESSION )
	if vcode != sessionCode {
		c.HTML(http.StatusOK, "login.html",gin.H{
			"msg":"验证码错误！",
			"path":"",
			"sessionInfo":GetSessionInfo(c),
		})
		return
	}
	userName := c.PostForm("UserPhone")
	passwd := c.PostForm("UserPassword")
	admin := yy.GetUserByPhone(userName)


	passwdMd5 := backend.Md5str(passwd)


	if admin.UserPassword == passwdMd5 {

		company := yy.GetById( admin.CompanyId )
		token := fmt.Sprintf("%s-%s",uuid.NewV4(), uuid.NewV4() )
		bus := UserSession{
			UserId:admin.Id,
			UserPhone:admin.UserPhone,
			UserName:admin.UserName,
			AccessToken:token,
			Success:true,
			CompanyId:company.Id,
			CompanyName:company.CompanyName,
			VipLevel:company.VipLevel,
		}
		session := sessions.Default(c)
		jsonBytes,_ := json.Marshal(bus)
		session.Set(FRONT_USER_SESSION,  string(jsonBytes) )
		session.Save()

		cookie := &http.Cookie{
			Name:     "accessToken",
			Value:    token,
			Domain:	  config.GetValue("cookie.domain"),
			Path:     "/",
			MaxAge:   config.GetInt("hiyuncms.server.frontend.session.timeout"),
			HttpOnly: true,
		}
		http.SetCookie(c.Writer, cookie)

		SetToken(token, &bus)

		c.Redirect(http.StatusFound, "/")
	} else{
		c.HTML(http.StatusOK, "login.html",gin.H{
			"msg":"用户名不存在或密码错误！",
			"path":"",
			"sessionInfo":GetSessionInfo(c),
			"bodyCss": "login-layout",
		})
	}
}

func RegistryShow(c * gin.Context)  {
	c.HTML(http.StatusOK, "registry.html", gin.H{
		"path":"",
		"sessionInfo":GetSessionInfo(c),
		"regsuccess":false,
	})
}

func Registry(c * gin.Context)  {
	isSuccess := true
	msg := ""
	user := yy.YyUser{}
	company := yy.YyCompany{}

	c.Bind( &user )
	c.Bind( &company )

	companyType := c.PostForm("company_type")
	vipType := c.PostForm("vip_type")

	file, header, err := c.Request.FormFile("license-img")
	if err != nil {
		c.String(http.StatusBadRequest, "Bad request")
		return
	}

	now := time.Now()
	filename := fmt.Sprintf("%d_%s",now.UnixNano(), header.Filename)

	fmt.Println(file, err, filename)

	dateFormart := "2006/01"

	filePath := fmt.Sprintf("license/image/%s/%s", now.Format(dateFormart), now.Format("02"))
	fmt.Println( filePath )
	os.MkdirAll(fmt.Sprintf("webroot/%s",filePath), os.ModePerm)
	out, err := os.Create(fmt.Sprintf("webroot/%s/%s", filePath,filename ) )
	if err != nil {
		isSuccess = false
		msg = "上传文件失败"
	}

	defer out.Close()
	_, err = io.Copy(out, file)
	if err != nil {
		isSuccess = false
		msg = "上传文件失败"
	}


	company.CompanyType = companyType
	company.CompanyVip = vipType
	company.CompanyBusinessLicense = fmt.Sprintf("webroot/%s/%s", filePath,filename )


	user.UserPassword = backend.Md5str(user.UserPassword)

	err, msg = yy.CompanyReg(&company, &user)
	if err != nil {
		log.Printf("%s\n", err.Error() )
		isSuccess = false
	}
	c.JSON(http.StatusOK, gin.H{
		"path":"",
		"regsuccess":isSuccess,
		"sessionInfo":GetSessionInfo(c),
		"msg": msg,
	})
}

func Logout(c *gin.Context)  {
	ClearSessionInfo(c)
	c.Redirect(http.StatusFound, "/")
}

func Verify(c *gin.Context)  {
	type  Verify struct{
		AppId string
		AccessToken string
	}

	verify := Verify{}
	c.Bind( &verify )
	if verify.AppId != config.GetValue("SSO") {
		c.JSON(http.StatusOK, gin.H{
			"success":false,
			"msg":"AppId不正确，没有权限访问！",
		})
	} else {
		sessionUser := GetToken( verify.AccessToken )
		if sessionUser != nil && sessionUser.UserPhone != "" {
			c.JSON(http.StatusOK, sessionUser)
		} else{
			c.JSON(http.StatusOK, gin.H{
				"success": false,
				"msg":     "token已失效！",
			})
		}
	}
}