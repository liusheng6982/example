package backend

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/contrib/sessions"
	"net/http"
	"log"
	"hiyuncms/controllers"
)
func Index( c *gin.Context )  {
	session := sessions.Default(c)
	sessionUser := session.Get(controllers.BACK_USER_SESSION)

	log.Printf("121212122211=%v\n", sessionUser)
	c.HTML(http.StatusOK, "index.html", gin.H{
		"mainMenu":"首页",
		"bodyCss": "no-skin",
		})

}

func Login(c *gin.Context){
	c.HTML(http.StatusOK, "login.html", gin.H{
		"bodyCss":"login-layout",

	})
}
