package backend

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"hiyuncms/models/cms"
	"hiyuncms/models"
	"fmt"
	"strconv"
	"hiyuncms/controllers"
)

/**
显示编辑页面
 */
func ArticleShow( c *gin.Context )  {
	columns := cms.GetAllColumnsToSelect()
	articleId := c.Query("id")
	article := cms.Article{}
	columnArticles := make( []* cms.ColumnArticle, 0)
	if articleId != "" {
		models.DbMaster.Id( articleId ).Get( &article )
		models.DbMaster.Table(cms.ColumnArticle{}).Where(fmt.Sprintf("Article_Id='%s'", articleId )).Find(&columnArticles)
	}
	c.HTML(http.StatusOK, "article.html", gin.H{
		"mainMenu":"新增文档",
		"bodyCss" : "no-skin",
		"columns" : columns,
		"article" : article,
		"columnArticles" : columnArticles,
		"user":controllers.GetSessionUser(c),
	})
}

/**
文章列表显示页面
 */
func ArticleListShow(c * gin.Context){
	c.HTML(http.StatusOK, "articlelist.html", gin.H{
		"mainMenu":"文档列表",
		"bodyCss" : "no-skin",
		"user":controllers.GetSessionUser(c),
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
	article := cms.Article{Status:2}
	c.Bind(&article)
	columns := c.PostFormArray("Columns[]")
	cms.SaveArticle(&article, columns)
	c.JSON(http.StatusOK, gin.H{
		"flag":"SUCCESS",
	})
}

func ArticleDel(c *gin.Context){
	articleIdstr, _ := c.GetPostForm("articleId")
	articleId, _ := strconv.ParseInt(articleIdstr, 10, 64)
	cms.DeleteArticle( articleId )
	c.JSON(http.StatusOK, gin.H{
		"flag":"SUCCESS",
	})
}

func ArticlePublish(c *gin.Context){
	articleIdstr, _ := c.GetPostForm("articleId")
	articleId, _ := strconv.ParseInt(articleIdstr, 10, 64)
	cms.PublishArticle( articleId )
	c.JSON(http.StatusOK, gin.H{
		"flag":"SUCCESS",
	})
}
