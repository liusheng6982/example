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
	"net/url"
	"io/ioutil"
	"hiyuncms/models"
)

/**
用户登录界面
 */
func UserLoginShow(c * gin.Context)  {
	c.HTML(http.StatusOK, "login.html", gin.H{
		"path":"",
		"sessionInfo":GetSessionInfo(c),
	})
}


func CompanyIndexShow(c * gin.Context)  {
	c.HTML(http.StatusOK, "companyindex.html", gin.H{
		"path":"",
		"sessionInfo":GetSessionInfo(c),
	})
}

/**
用户登录动作
 */
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
		now := time.Now()

		expired := 1
		if now.Before( time.Time(company.VipExpired) ) || now.Equal( time.Time(company.VipExpired) ) {
			log.Printf("asdfasdfasdfasdfasdfasdf\n")
			expired = 0
		}
		log.Printf("asdfasdfasdfasdfasdfasdf expired=%d\n", expired)
		bus := UserSession{
			UserId:admin.Id,
			UserPhone:admin.UserPhone,
			UserName:admin.UserName,
			AccessToken:token,
			Success:true,
			CompanyId:company.Id,
			CompanyName:company.CompanyName,
			VipExpired:expired,
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

		if company.CompanyType == "1"{
			c.Redirect(http.StatusFound, "/companyindex")
		}else{
			c.Redirect(http.StatusFound, "/")
		}


	} else{
		c.HTML(http.StatusOK, "login.html",gin.H{
			"msg":"用户名不存在或密码错误！",
			"path":"",
			"sessionInfo":GetSessionInfo(c),
			"bodyCss": "login-layout",
		})
	}
}

/**
用户注册界面显示
 */
func RegistryShow(c * gin.Context)  {
	c.HTML(http.StatusOK, "registry.html", gin.H{
		"path":"",
		"sessionInfo":GetSessionInfo(c),
		"regsuccess":false,
	})
}

/**
用户注册界面显示
 */
func RegistryVerify(c * gin.Context)  {
	isSuccess := true
	msg := ""
	user := yy.YyUser{}
	company := yy.YyCompany{}

	verifyCode, _ := c.GetPostForm("VerifyCode")

	fmt.Printf("http verify code = %s\n", verifyCode)

	c.Bind( &user )
	c.Bind( &company )

	if company.CompanyName == ""{
		isSuccess = false
		msg = "公司名称不能为空！"

		c.JSON(http.StatusOK,gin.H{
			"isSuccess":isSuccess,
			"msg":msg,
		})
		return
	}

	if user.UserPhone == ""{
		isSuccess = false
		msg = "手机号不能为空！"

		c.JSON(http.StatusOK,gin.H{
			"isSuccess":isSuccess,
			"msg":msg,
		})
		return
	}


	if verifyCode == ""{
		isSuccess = false
		msg = "手机短信验证码不能为空！"

		c.JSON(http.StatusOK,gin.H{
			"isSuccess":isSuccess,
			"msg":msg,
		})
		return
	}

	count1, err := models.DbSlave.Count(&user)
	if err != nil {
		isSuccess = false
		msg = err.Error()
	}
	if count1 > 0 {
		isSuccess = false
		msg = "手机号码注册过！"

		c.JSON(http.StatusOK,gin.H{
			"isSuccess":isSuccess,
			"msg":msg,
		})
		return
	}


	count, err := models.DbSlave.Count(&company)
	if err != nil {
		isSuccess = false
		msg = err.Error()
	}
	if count > 0 {
		isSuccess = false
		msg = "公司名称已存在！"

		c.JSON(http.StatusOK,gin.H{
			"isSuccess":isSuccess,
			"msg":msg,
		})
		return
	}
	session := sessions.Default(c)
	sessionSmsKey := fmt.Sprintf("%s%s",FRONT_SMS, user.UserPhone)
	tempVerifyCode := session.Get( sessionSmsKey )
	fmt.Printf("session verify code = %s", tempVerifyCode)
	if tempVerifyCode != verifyCode {
		isSuccess = false
		msg ="验证码不正确！"
	}
	c.JSON(http.StatusOK,gin.H{
		"isSuccess":isSuccess,
		"msg":msg,
	})
}


/**
用户注册
 */
func Registry(c * gin.Context)  {
	isSuccess := true
	msg := ""
	user := yy.YyUser{}
	company := yy.YyCompany{}

	c.Bind( &user )
	c.Bind( &company )

	verifyCode, _ := c.GetPostForm("VerifyCode")

	session := sessions.Default(c)
	sessionSmsKey := fmt.Sprintf("%s%s",FRONT_SMS, user.UserPhone)
	tempVerifyCode := session.Get( sessionSmsKey )
	fmt.Printf("session verify code = %s", tempVerifyCode)
	if tempVerifyCode != verifyCode {
		isSuccess = false
		msg ="验证码不正确！"
		c.JSON(http.StatusOK, gin.H{
			"path":"",
			"regsuccess":isSuccess,
			"sessionInfo":GetSessionInfo(c),
			"msg": msg,
		})
		return
	}

	companyType := c.PostForm("company_type")
	log.Printf("1231231231231231231231231231231231231231231%s",companyType)
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

	chuanyiwangSync := func(){ //用户同步
		url := config.GetValue("sync.user.chuanyiwang.url")
		getUrl := fmt.Sprintf("%s?userId=%d&userPhone=%s&userName=%s&companyId=%d&companyName=%s&companyType=%s",
			url, user.Id, user.UserPhone, user.UserName, company.Id, company.CompanyName, company.CompanyType)


		res, err := http.Get(getUrl)
		log.Printf("!!!!!!!!!!!!!!!!!!%s", err)
		if err == nil {
			body, err1 := ioutil.ReadAll(res.Body)
			log.Printf("%s\n", body)
			if err1 != nil {
				log.Printf("err1=%s\n", err1)
			}
		} else {
			log.Printf("err=%s\n", err)
		}
	}
	go chuanyiwangSync()


	guoxinSync := func() {
		{ //组织结构同步
			data := make(url.Values)
			data["type"] = []string{"YYG"}
			orgRoleName := ""
			if companyType == "1" {
				orgRoleName = "11"
			}
			if companyType == "2"{
				orgRoleName = "12"
			}
			companyInfo := fmt.Sprintf("{\"orgId\":\"%d\", \"orgName\":\"%s\", \"orgRoleName\":\"%s\"}", company.Id, company.CompanyName, orgRoleName)
			data["data"] = []string{companyInfo}

			res, err := http.PostForm(config.GetValue("sync.org.guoxin.url"), data)
			log.Printf("!!!!!!!!!!!!!!!!!!%s", err)
			if err == nil {
				body, err1 := ioutil.ReadAll(res.Body)
				log.Printf("%s\n", body)
				if err1 != nil {
					log.Printf("err1=%s\n", err1)
				}
			} else {
				log.Printf("err=%s\n", err)
			}
		}

		{ //用户同步
			data := make(url.Values)
			data["type"] = []string{"YYG"}
			data["userId"] = []string{fmt.Sprintf("%d", user.Id)}
			data["userPhone"] = []string{user.UserPhone}
			data["username"] = []string{user.UserName}
			data["companyId"] = []string{fmt.Sprintf("%d", company.Id)}
			data["companyName"] = []string{company.CompanyName}
			data["companyType"] = []string{company.CompanyType}
			data["password"] = []string{user.UserPassword}

			res, err := http.PostForm(config.GetValue("sync.user.guoxin.url"), data)
			log.Printf("!!!!!!!!!!!!!!!!!!%s", err)
			if err == nil {
				body, err1 := ioutil.ReadAll(res.Body)
				log.Printf("%s\n", body)
				if err1 != nil {
					log.Printf("err1=%s\n", err1)
				}
			} else {
				log.Printf("err=%s\n", err)
			}
		}
	}
	go guoxinSync()

	c.JSON(http.StatusOK, gin.H{
		"path":"",
		"regsuccess":isSuccess,
		"sessionInfo":GetSessionInfo(c),
		"msg": msg,
	})
}

/**
用户登出
 */
func Logout(c *gin.Context)  {
	sessionInfo := GetSessionInfo(c)
	ClearSessionInfo(c)
	DelToken( sessionInfo.AccessToken )
	c.Redirect(http.StatusFound, "/")
}

/**
用户没有成为VIP
 */
func NoVip(c *gin.Context){
	c.HTML(http.StatusOK, "novip.html", gin.H{
		"path":"",
		"sessionInfo":GetSessionInfo(c),
		"regsuccess":false,
	})
}

/**
VIP过期
 */
func VipExpired(c *gin.Context)  {
	c.HTML(http.StatusOK, "vipexpired.html", gin.H{
		"path":"",
		"sessionInfo":GetSessionInfo(c),
		"regsuccess":false,
	})
}

/**
单点登录token验证
 */
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
