package backend

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"hiyuncms/models"
	"hiyuncms/models/cms"
)

func ColumnList(c *gin.Context){
	c.HTML(http.StatusOK, "columnlist.html", gin.H{
		"bodyCss":"no-skin",
		"mainMenu" :"栏目管理",
	})
}

func ColumnDataList(c *gin.Context){
	page := models.PageRequest{}
	c.Bind( &page )
	responsePage := cms.GetAllColumns(&page)
	c.JSON(http.StatusOK, responsePage)
}
