package main


//import "net/http"
//import "github.com/gin-contrib/multitemplate"
import (
	"github.com/gin-gonic/gin"
	"net/http"
	"controllers"
	//"reflect"
	"fmt"
	//"models"
	"service"
	"os/signal"
	"syscall"
	"os"
	"github.com/gin-contrib/multitemplate"
)

var port int
func init()  {
	port = 8080
}

func main() {
		//runServer()
		//runServer2()
		runServer3()
		RegService()
	}

func runTest()  {
		/*
		tt := models.User{Id:1, UserName:"Barak Obama"}
		var reflectType reflect.Type = reflect.TypeOf(tt)
		var ixField reflect.StructField
		for i := 0; i < 3; i++ {
			ixField = reflectType.Field(i)
			fmt.Printf("%s\n", ixField.Tag)
		}*/
	}

func createMyRender() multitemplate.Render {
	r := multitemplate.New()
	r.AddFromFiles("index", "template/base.html", "template/index.html")
	r.AddFromFiles("article", "template/base.html", "template/index.html", "template/article.html")
	return r
}

func runServer3() {
	router := gin.Default()
	router.HTMLRender = createMyRender()
	router.GET("/", func(c *gin.Context) {
		c.HTML(200, "index", gin.H{
			"title": "Html5 Template Engine",
		})
	})
	router.GET("/article", func(c *gin.Context) {
		c.HTML(200, "article", gin.H{
			"title": "Html5 Article Engine",
		})
	})
	go router.Run(":8083")
}

func runServer2()  {
	e := gin.New()
	e.Use(gin.Recovery())
	e.Use( gin.Logger() )
	e.LoadHTMLGlob("template/**/**/*")
	//e.LoadHTMLGlob("template/**/*")
	//e.LoadHTMLGlob("template/*")

	e.GET("/b", func(c *gin.Context) {
		c.HTML(http.StatusOK, "b.tmpl", gin.H{
			"title": "Main website",
		})
	})

	go http.ListenAndServe(fmt.Sprintf(":%d", 8081), e)
}


func runServer(){
	gin.SetMode(gin.DebugMode)
	router := gin.Default()
	router.GET("/index", controllers.Index)
	router.GET("/users", controllers.GetUsers)

	//router.POST("/orgs", controllers.OrgIndex)
	router.LoadHTMLGlob("template/**/**/*")
	//router.LoadHTMLGlob("template/**/*")
	//router.LoadHTMLGlob("template/*")

	//router.LoadHTMLFiles("templates/template1.html", "templates/template2.html")
	router.GET("/a", func(c *gin.Context) {
		c.HTML(http.StatusOK, "a.tmpl", gin.H{
			"title": "Main website",
		})
	})

	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})


	//router.LoadHTMLGlob(filepath.Join(os.Getenv("GOPATH"), "src/github.com/ameykpatil/gospike/templates/*"))

	router.StaticFS("static", http.Dir("/Users/liusheng/GoglandProjects/martini/src/static"))

	//print( "1111",err.Error() )

	//go router.Run(fmt.Sprintf(":%d", port), "/Users/liusheng/server.crt","/Users/liusheng/server.key")

	//go http.ListenAndServeTLS(":8080", "/Users/liusheng/server.crt", "/Users/liusheng/server.key", router)
	go http.ListenAndServe(fmt.Sprintf(":%d", port), router)


}

func RegService(){
	service.Port = port

	//`暂时不要祖册`
	//service.RegService()



	// Go signal notification works by sending `os.Signal`
	// values on a channel. We'll create a channel to
	// receive these notifications (we'll also make one to
	// notify us when the program can exit).
	sigs := make(chan os.Signal, 1)

	// `signal.Notify` registers the given channel to
	// receive notifications of the specified signals.
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	// This goroutine executes a blocking receive for
	// signals. When it gets one it'll print it out
	// and then notify the program that it can finish.
	<-sigs
	fmt.Println("停止！！！！！！！！！！！！！！！！！！！")
	//service.UnRegService()

}