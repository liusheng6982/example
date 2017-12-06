package yy

import (
	"github.com/gin-gonic/gin"
	"hiyuncms/controllers"
	"net/http"
)

func PurchaseListShow(c *gin.Context){
	c.HTML(http.StatusOK, "purchaselist.html", gin.H{
		"bodyCss":"no-skin",
		"mainMenu" :"采购项目管理",
		"user":controllers.GetSessionUser(c),
	})
}