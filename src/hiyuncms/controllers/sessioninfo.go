package controllers

import (
	"github.com/gin-gonic/gin"
	"math/rand"
	"hiyuncms/util"
	"fmt"
	"github.com/gin-gonic/contrib/sessions"
	"time"
)

const(
	BACK_USER_SESSION = "hiyuncms.back.user"
	BACK_CAPTCHA_SESSION = "hiyumcms.back.captcha"
)
type BackendUserSession struct {
	Name string `json:"name"`
	Id   int64  `json:"id"`
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
	session.Delete(BACK_CAPTCHA_SESSION)
	session.Set(BACK_CAPTCHA_SESSION, ss)
	session.Save()
	util.NewImage(fmt.Sprintf("%d",rand.Int()),d, 80, 40).WriteTo(c.Writer)
}