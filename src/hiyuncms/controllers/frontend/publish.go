package frontend

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"hiyuncms/models/cms"
	"strconv"
)

func ArticlesShow( c *gin.Context )  {
	pageNoStr := c.Query("pageNo")
	if pageNoStr == "" {
		pageNoStr = "1"
	}
	pageSizeStr := c.Query("pageSize")
	if pageSizeStr == "" {
		pageSizeStr = "10"
	}
	path := c.Request.URL.Path
	column :=  cms.GetColumnByPath( path )  //查询出模板路径

	pageNo,_:= strconv.Atoi(pageNoStr)
	pageSize,_ := strconv.Atoi(pageSizeStr)
	c.HTML(http.StatusOK, column.TemplatePath, gin.H{
		"path":path,
		"pageNo":pageNo,
		"pageSize":pageSize ,
	})
}

func ArticleShow( c *gin.Context )  {
	articleId := c.Query("articleId")
	c.HTML(http.StatusOK, "articleshow.html", gin.H{
		"articleId":articleId,
		"path":"",
	})
}
