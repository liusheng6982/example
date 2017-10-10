package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"html/template"
	"hiyuncms/models/cms"
	"hiyuncms/models"
	"hiyuncms/controllers/frontend"
)

var FrontendRoute *gin.Engine

func init()  {
	FrontendRoute = initRouteFrontend()
	regFrontRoute()
}

func initRouteFrontend()   *gin.Engine{
	engine := gin.New()
	engine.Use(gin.Recovery())
	engine.Use( gin.Logger() )
	engine.SetFuncMap(template.FuncMap{
		"loadArticles": loadArticles,
		"html" : html,
	})
	engine.LoadHTMLGlob("webroot/templates/frontend/**/*")
	engine.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})
	engine.StaticFS("static", http.Dir("webroot/static"))
	return engine
}

func regFrontRoute()  {
	FrontendRoute.GET("/publish", frontend.ArticleShow)
}

func html (x string) interface{} {
	return template.HTML(x)
}

func loadArticles(path string, pageSize int, pageNo int ) * models.PageResponse {
	page := models.PageRequest{Rows:pageSize, Page:pageNo}
	return cms.GetArticlesByPath(&page, path)
}


