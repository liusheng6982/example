package yy

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"hiyuncms/controllers"
	"hiyuncms/models"
	"hiyuncms/models/yy"
)

/**
用户列表显示
 */
func VipUserShow( c *gin.Context )  {
	c.HTML(http.StatusOK, "vip-user-list.html", gin.H{
		"mainMenu":"会员用户列表",
		"bodyCss" : "no-skin",
		"user":controllers.GetSessionUser(c),
	})
}

/**
用户数据
 */
func VipUserList(c *gin.Context){
	page := models.PageRequest{}
	c.Bind( &page )
	c.JSON( http.StatusOK, yy.GetVipUsers(&page) )
}

