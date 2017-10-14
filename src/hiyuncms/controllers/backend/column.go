package backend

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"hiyuncms/models"
	"hiyuncms/models/cms"
	"log"
	"strconv"
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
	responsePage := cms.GetAllColumnsByPage(&page)
	c.JSON(http.StatusOK, responsePage)
}

func ColumnEdit(c * gin.Context){
	column := cms.Column{}
	c.Bind( &column )
	oper, _ := c.GetPostForm("oper")
	if "edit" == oper {
		id, _:= c.GetPostForm("id")
		column.Id, _= strconv.ParseInt(id, 10, 64)
		_, err := models.DbMaster.ID(column.Id).Update(&column)
		if err != nil {
			log.Printf("更新Cloun报错:%s\n",models.GetErrorInfo(err))
		}
	}else if"add" == oper {
		_, err := models.DbMaster.Insert( &column )
		if err != nil {
			log.Printf("新增Cloun报错:%s\n",models.GetErrorInfo(err))
		}
	}
}
