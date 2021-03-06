package routes

import (
	"strings"
	"github.com/gin-gonic/gin"
	"net/http"
	"html/template"
	"hiyuncms/models/cms"
	"hiyuncms/models"
	"hiyuncms/controllers/frontend"
	"strconv"
	"github.com/gin-gonic/contrib/sessions"
	"hiyuncms/config"
	"fmt"
)

var FrontendRoute *gin.Engine

func init()  {
	FrontendRoute = initRouteFrontend()
	regFrontRoute()
}

func initRouteFrontend()   *gin.Engine{
	if !config.GetBool("hiyuncms.application.debug") {
		gin.SetMode(gin.ReleaseMode)
	}
	engine := gin.New()
	engine.Use(gin.Recovery())
	engine.Use( gin.Logger() )
	store := sessions.NewCookieStore([]byte("jsessionid_front_session"))
	store.Options(sessions.Options{
		MaxAge: int(config.GetInt("hiyuncms.server.frontend.session.timeout")), //30min
		Path:   "/",
	})
	engine.Use( sessions.Sessions("jsessionid", store) )
	engine.SetFuncMap(template.FuncMap{
		"loadColumn":   loadColumn,
		"loadArticles": loadArticlesByPage,
		"loadArticlesTop": loadArticlesTop,
		"loadArticle" : loadArticle,
		"loadSubColumn": loadSubColumn,
		"html"  : html,
		"addNum": addNum,
		"hasPrefix": HasPrefix,
		"substring": Substring,
		"substring2": Substring2,
	})
	engine.LoadHTMLGlob("webroot/templates/frontend/**/*")
	engine.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})
	engine.StaticFS("static", http.Dir("webroot/static"))
	return engine
}

func regFrontRoute()  {
	columns := cms.GetAllColumnsToRoute()
	indexPage := ""
	for k,column := range *columns{
		if k == 0 {
			indexPage = column.Url
		}
		FrontendRoute.GET( column.Url, frontend.ArticlesShow)
	}

	if indexPage != ""  {
		FrontendRoute.GET("/", func(context *gin.Context) {
			context.Redirect(http.StatusFound, indexPage)
		})
	}
	//FrontendRoute.GET("/published/:route", frontend.ArticlesShow)
	FrontendRoute.GET ("/articleShow", frontend.ArticleShow)

	
}

func Substring2(str string, length int) string {
	if len([]rune(str)) > length {
		return fmt.Sprintf("%s...",string([]rune(str)[0:length]))
	}
	return str
}

func Substring(str, split string) string{
	index := strings.LastIndex(str, split)
	if index > 0 {
		return str[0:index]
	}
	return str
}

func HasPrefix(str, substr string) bool{
	return strings.HasPrefix(str, substr)
}

func addNum(x int, y int )int{
	return x + y
}

/**
转换成html
 */
func html (x string) interface{} {
	return template.HTML(x)
}

/**
获得某个栏目下的所有文档
 */
func loadArticlesByPage(path string, pageSize int, pageNo int ) * models.PageResponse {
	page := models.PageRequest{Rows:pageSize, Page:pageNo}
	response := cms.GetArticlesByPath(&page, path)
	response.Path = path
	return response
}



/**
获得某个栏目下的前几条记录
 */
func loadArticlesTop(path string, begin int, end int ) []*cms.Article {
	response := cms.GetArticlesByPathTop(path, begin, end)
	return response
}

/**
根据ID获得文档详细
 */
func loadArticle(articleId string) *cms.Article {
	articleId64,_ := strconv.ParseInt(articleId, 10, 64)
	article := cms.GetArticle(articleId64)
	return article
}

/**
获得所有需要显示的栏目
 */
func loadColumn() *[]*cms.Column{
	return cms.GetAllColumnsToShow()
}

func loadSubColumn(parentPath string)  *[]*cms.Column{
	return cms.GetSubColumnsToShow(parentPath)
}

