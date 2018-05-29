package frontend

import (
	"github.com/gin-gonic/gin"
	"math/rand"
	"hiyuncms/util"
	"fmt"
	"github.com/gin-gonic/contrib/sessions"
	"time"
	"reflect"
	"log"
	"encoding/json"
	"net/http"
)

const(

	FRONT_USER_SESSION    = "hiyuncms.front.User"
	FRONT_CAPTCHA_SESSION = "hiyumcms.front.captcha"
	FRONT_SMS = "hiyumcms.front.sms"
)
type UserSession struct {
	UserId            int64  		`json:"userId"`
	UserName 	      string 		`json:"userName"`
	UserPhone         string 		`json:"userPhone"`
	CompanyId 		  int64  		`json:"companyId"`
	CompanyName       string     	`json:"companyName"`
	AccessToken 	  string		`json:"accessToken"`
	Success 		  bool          `json:"success"`
	VipLevel          int64           `json:"vipLevel"`
	VipExpired        int           `json:"vipExpired"`
	CompanyType       string 			`json:"companyType"`
}

func GetSessionInfo(c * gin.Context)  *UserSession{
	session := sessions.Default(c)
	sessionStr := session.Get(FRONT_USER_SESSION)
	log.Printf("sessionifo=%v\n", reflect.TypeOf(sessionStr))
	sessionInfo := &UserSession{}
	if sessionStr != nil  {
		json.Unmarshal([]byte(sessionStr.(string)), sessionInfo)
	}
	return sessionInfo
}

func ClearSessionInfo(c * gin.Context){
	session := sessions.Default(c)
	session.Clear()
	session.Save()
}

func SendSMS(c * gin.Context){
	d := make([]byte, 6)
	mobile := c.PostForm("UserPone")
	ss := ""
	rd := rand.New(rand.NewSource(time.Now().UnixNano()))
	for v := range d {
		d[v] = byte(rd.Intn(10))
		ss = fmt.Sprintf("%s%d", ss, d[v])
	}
	log.Printf("----------mima:%s", ss)
	session := sessions.Default(c)
	sessionSmsKey := fmt.Sprintf("%s%s",FRONT_SMS,mobile)
	session.Delete(sessionSmsKey)
	session.Set(sessionSmsKey, ss)
	session.Save()
	err := util.SendSms(ss, mobile)
	if err == nil {
		c.JSON(http.StatusOK, gin.H{
			"success": true,
		})
	}else{
		c.JSON(http.StatusOK, gin.H{
			"success": false,
		})
	}
}

func Captcha(c * gin.Context) {
	d := make([]byte, 4)

	ss := ""
	rd := rand.New(rand.NewSource(time.Now().UnixNano()))
	for v := range d {
		d[v] = byte(rd.Intn(10))
		ss = fmt.Sprintf("%s%d", ss, d[v])
	}
	c.Header("Content-Type", "image/png")

	fmt.Printf("ssssssss=%s\n", ss)
	session := sessions.Default(c)
	session.Delete(FRONT_CAPTCHA_SESSION)
	session.Set(FRONT_CAPTCHA_SESSION, ss)
	session.Save()
	util.NewImage(fmt.Sprintf("%d",rand.Int()),d, 80, 40).WriteTo(c.Writer)
}