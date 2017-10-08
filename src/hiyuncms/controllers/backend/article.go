package backend

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"hiyuncms/models/cms"
	"hiyuncms/models"
)

func ArticleShow( c *gin.Context )  {
	columns := cms.GetAll()
	c.HTML(http.StatusOK, "article.html", gin.H{
		"mainMenu":"新增文档",
		"bodyCss" : "no-skin",
		"columns" : columns,
	})
}


func ArticleListShow(c * gin.Context){
	c.HTML(http.StatusOK, "articlelist.html", gin.H{
		"mainMenu":"文档列表",
		"bodyCss" : "no-skin",
	})
}

func ArticleListData(c *gin.Context){
	page := models.PageRequest{}
	c.Bind( &page )
	responsePage := cms.GetAllArticles(&page)
	c.JSON(http.StatusOK, responsePage)
}


func ArticleSave(c *gin.Context)  {
	article := cms.Article{}
	c.Bind(&article)
	columns := c.PostFormArray("Columns[]")
	cms.SaveArticle(&article, columns)
	c.JSON(http.StatusOK, gin.H{
		"flag":"SUCCESS",
	})
}
