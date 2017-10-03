package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"hiyuncms/util"
)

var FrontendRoute *gin.Engine

func init()  {
	FrontendRoute = initRouteFrontend()
}

func initRouteFrontend()   *gin.Engine{
	e := gin.New()
	e.Use(gin.Recovery())
	e.Use( gin.Logger() )
	e.LoadHTMLGlob("webroot/templates/frontend/**/*")
	e.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})
	e.StaticFS("static", http.Dir(util.GetCurrPath() + "webroot/static"))
	return e
}

