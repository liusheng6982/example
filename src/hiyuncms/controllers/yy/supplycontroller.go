package yy

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"hiyuncms/controllers"
)

/**
显示编辑页面
 */
func SupplyShow( c *gin.Context )  {
	c.HTML(http.StatusOK, "supply-list.html", gin.H{
		"mainMenu":"供应商列表",
		"bodyCss" : "no-skin",
		"user":controllers.GetSessionUser(c),
	})
}