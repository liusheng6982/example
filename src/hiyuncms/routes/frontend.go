package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"html/template"
	"hiyuncms/models/cms"
	"hiyuncms/models"
	"hiyuncms/controllers/frontend"
	"strconv"
	"github.com/gin-gonic/contrib/sessions"
	"hiyuncms/config"
	"hiyuncms/models/yy"
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
		"html"  : html,
		"addNum": addNum,
		"loadTopInviteTender" : yy.GetTopInviteTender,
		"loadTopPurchase":yy.GetTopPurchase,
		"loadTopRecommend":yy.GetTopRecommendInviteTender,
		"loadTopWinBid":yy.GetTopWinBidInviteTender,
		"loadAllProject":loadAllProject,
	})
	engine.LoadHTMLGlob("webroot/templates/frontend/**/*")
	engine.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})
	engine.StaticFS("static", http.Dir("webroot/static"))
	return engine
}

func regFrontRoute()  {
	columns := cms.GetAllColumnsToShow()
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
	FrontendRoute.GET ("/userlogin",frontend.UserLoginShow)
	FrontendRoute.POST("/userlogin",frontend.UserLogin)
	FrontendRoute.GET ("/logout", frontend.Logout)
	FrontendRoute.GET ("/captcha", frontend.Captcha)        //验证码
	FrontendRoute.GET ("/verifyToken", frontend.Verify)        //验证码

	FrontendRoute.GET ("/registry",frontend.RegistryShow)
	FrontendRoute.POST("/registry",frontend.Registry)
	
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

func loadAllProject(pageSize int, pageNo int ) * models.PageResponse {
	page := models.PageRequest{Rows:pageSize, Page:pageNo}
	response := yy.GetAllInviteTenderByPage(&page)
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

