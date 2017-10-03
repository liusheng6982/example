package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

var BackendRoute *gin.Engine

func init()  {
	BackendRoute = initRouteBackend()
	initRoute()
}

func initRouteBackend() *gin.Engine {
	e := gin.New()
	e.Use(gin.Recovery())
	e.Use( gin.Logger() )
	e.LoadHTMLGlob("webroot/templates/backend/**/*")
	e.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})
	e.StaticFS("static", http.Dir("webroot/static"))
	return e
}

func initRoute()  {

}

