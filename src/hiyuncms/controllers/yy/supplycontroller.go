package yy

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"hiyuncms/controllers"
	"hiyuncms/models"
	"hiyuncms/models/yy"
)

/**
显示编辑页面
 */
func CompanyShow( c *gin.Context )  {
	c.HTML(http.StatusOK, "supply-list.html", gin.H{
		"mainMenu":"公司列表",
		"bodyCss" : "no-skin",
		"user":controllers.GetSessionUser(c),
	})
}

/**
公司列表
 */
func CompanyList(c *gin.Context){
	page := models.PageRequest{}
	c.Bind( &page )
	c.JSON( http.StatusOK, yy.GetSupplies(&page) )
}