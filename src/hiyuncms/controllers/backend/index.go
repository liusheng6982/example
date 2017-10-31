package backend

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"hiyuncms/controllers"

)
func Index( c *gin.Context )  {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"mainMenu":"首页",
		"bodyCss": "no-skin",
		"user":controllers.GetSessionUser(c),
	})

}

func Login(c *gin.Context){
	c.HTML(http.StatusOK, "login.html", gin.H{
		"bodyCss":"login-layout",
	})
}

func Logout(c *gin.Context)  {
	controllers.ClearSessionUser(c)
	c.Redirect(http.StatusFound, "/login")
}
