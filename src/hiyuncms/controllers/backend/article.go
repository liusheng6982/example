package backend

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"hiyuncms/models/cms"
	"hiyuncms/models"
)

/**
显示编辑页面
 */
func ArticleShow( c *gin.Context )  {
	columns := cms.GetAllColumnsToSelect()
	c.HTML(http.StatusOK, "article.html", gin.H{
		"mainMenu":"新增文档",
		"bodyCss" : "no-skin",
		"columns" : columns,
	})
}

/**
文章列表显示页面
 */
func ArticleListShow(c * gin.Context){
	c.HTML(http.StatusOK, "articlelist.html", gin.H{
		"mainMenu":"文档列表",
		"bodyCss" : "no-skin",
	})
}

/**
文章列表
 */
func ArticleListData(c *gin.Context){
	page := models.PageRequest{}
	c.Bind( &page )
	responsePage := cms.GetAllArticles(&page)
	c.JSON(http.StatusOK, responsePage)
}

/**
保存文章
 */
func ArticleSave(c *gin.Context)  {
	article := cms.Article{}
	c.Bind(&article)
	columns := c.PostFormArray("Columns[]")
	cms.SaveArticle(&article, columns)
	c.JSON(http.StatusOK, gin.H{
		"flag":"SUCCESS",
	})
}
