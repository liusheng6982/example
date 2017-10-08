package routes

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"hiyuncms/controllers/backend"
	"github.com/gin-gonic/contrib/sessions"
	"time"
	"strings"
)

var BackendRoute *gin.Engine

const SessionName = "hiyun_backend_session"

func init()  {
	BackendRoute = initRouteBackend()
	regRoute()
}

func MiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		sessionUser := session.Get("hiyuncms.back.user")

		if sessionUser == nil {
			if c.Request.URL.Path == "/login" || c.Request.URL.Path == "/userlogin" || strings.Contains(c.Request.URL.Path,"/static/") {
				c.Next()
			} else{
				c.Redirect(http.StatusFound, "/login")
			}
		}
		c.Next()
	}
}

func initRouteBackend() *gin.Engine {
	//gin.SetMode(gin.ReleaseMode)
	engine := gin.New()

	engine.Use(gin.Recovery())
	engine.Use( gin.Logger() )
	store := sessions.NewCookieStore([]byte("hiyuncms.secret"))
	store.Options(sessions.Options{
		MaxAge: int(30 * time.Minute), //30min
		Path:   "/",
	})
	engine.Use( sessions.Sessions(SessionName, store) )
	engine.Use( MiddleWare() )
	engine.LoadHTMLGlob("webroot/templates/backend/**/*")
	engine.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})
	engine.StaticFS("static", http.Dir("webroot/static"))
	return engine
}

func regRoute()  {
	BackendRoute.GET("/index", backend.Index)
	BackendRoute.GET("/login", backend.Login ) //打开login页面
	BackendRoute.POST("/login",backend.UserLogin)//提交登录

	BackendRoute.GET("/columnlist",backend.ColumnList)      //栏目列表
	BackendRoute.POST("/columnlist",backend.ColumnDataList) //栏目列表数据
	BackendRoute.GET("/article", backend.ArticleShow)       //新增文档时显示
	BackendRoute.POST("/article", backend.ArticleSave)      //新增文档
	BackendRoute.GET("/articlelist", backend.ArticleListShow)   //新增列表页
	BackendRoute.POST("/articlelist", backend.ArticleListData)  //新增表数据
	BackendRoute.GET("/UEditorAction", backend.UEdit)
	BackendRoute.POST("/UEditorAction", backend.UEditAction)
}

