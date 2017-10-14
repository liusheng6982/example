package controllers

import (
	"github.com/gin-gonic/gin"
	"math/rand"
	"hiyuncms/util"
	"fmt"
	"github.com/gin-gonic/contrib/sessions"
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
	for v := range d {
		d[v] = byte(rand.Intn(10))
		ss = fmt.Sprintf("%s%d", ss, d[v])
	}
	c.Header("Content-Type", "image/png")

	session := sessions.Default(c)
	session.Delete(BACK_CAPTCHA_SESSION)
	session.Set(BACK_CAPTCHA_SESSION, ss)
	util.NewImage(fmt.Sprintf("%d",rand.Int()),d, 80, 40).WriteTo(c.Writer)
}