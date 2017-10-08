package backend

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"hiyuncms/models"
	"hiyuncms/models/cms"
	"log"
)

func ColumnList(c *gin.Context){
	c.HTML(http.StatusOK, "columnlist.html", gin.H{
		"bodyCss":"no-skin",
		"mainMenu" :"栏目管理",
	})
}

func ColumnDataList(c *gin.Context){
	page := models.PageRequest{}
	log.Print("before:%v\n", page)
	c.Bind( &page )
	log.Print("bind:%v\n", page)
	responsePage := cms.GetAllColumns(&page)
	log.Print("after:%v\n", responsePage)
	c.JSON(http.StatusOK, responsePage)
}
