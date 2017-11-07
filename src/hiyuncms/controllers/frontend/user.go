package frontend

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func UserLoginShow(c * gin.Context)  {
	c.HTML(http.StatusOK, "login.html", gin.H{
		"path":"",
	})
}

func UserLogin(c * gin.Context)  {
	c.HTML(http.StatusOK, "login.html", gin.H{
		"path":"",
	})
}

func RegistryShow(c * gin.Context)  {
	c.HTML(http.StatusOK, "registry.html", gin.H{
		"path":"",
	})
}

func Registry(c * gin.Context)  {
	c.HTML(http.StatusOK, "registry.html", gin.H{
		"path":"",
	})
}