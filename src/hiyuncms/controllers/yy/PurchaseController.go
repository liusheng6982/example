package yy

import (
	"github.com/gin-gonic/gin"
	"hiyuncms/controllers"
	"net/http"
	"hiyuncms/models"
	"hiyuncms/models/yy"
	"strconv"
	"log"
)

func PurchaseListShow(c *gin.Context){
	c.HTML(http.StatusOK, "purchaselist.html", gin.H{
		"bodyCss":"no-skin",
		"mainMenu" :"采购项目管理",
		"user":controllers.GetSessionUser(c),
	})
}

func PurchaseList(c *gin.Context){
	page := models.PageRequest{}
	c.Bind( &page )
	responsePage := yy.GetAllYyPurchaseByPage(&page)
	c.JSON(http.StatusOK, responsePage)
}

func PurchaseEdit(c * gin.Context){
	purchase := yy.YyPurchase{}
	bindErr := c.Bind( &purchase)
	quotePriceEndDate := c.PostForm("QuotePriceEndTime")
	convertErr := purchase.QuotePriceEndTime.UnmarshalText( []byte(quotePriceEndDate))
	if convertErr != nil {
		log.Printf("quotePriceEndDate 绑定数据出错:%s\n",models.GetErrorInfo(convertErr))
	}
	log.Printf("quotePriceEndDate=%s\n",quotePriceEndDate)
	if bindErr != nil {
		log.Printf("新增Purchase 绑定数据出错:%s\n",models.GetErrorInfo(bindErr))
	}
	log.Printf("%+v", purchase)
	oper, _ := c.GetPostForm("oper")
	if "edit" == oper {
		id, _:= c.GetPostForm("id")
		purchase.Id, _= strconv.ParseInt(id, 10, 64)
		_, err := models.DbMaster.ID(purchase.Id).Update(&purchase)
		if err != nil {
			log.Printf("更新Purchase报错:%s\n",models.GetErrorInfo(err))
		}
	}else if "add" == oper {
		purchase.CreateTime = models.Date{}
		_, err := models.DbMaster.Insert( &purchase)
		if err != nil {
			log.Printf("新增Purchase报错:%s\n",models.GetErrorInfo(err))
			c.String(http.StatusInternalServerError, "%s", "fail")
			return
		}
		c.String(http.StatusOK, "%s", "success")
	} else if "del" == oper{
		id, _:= c.GetPostForm("id")
		purchase.Id, _= strconv.ParseInt(id, 10, 64)
		_, err := models.DbMaster.Delete(&purchase)
		if err != nil {
			log.Printf("删除Org报错:%s\n",models.GetErrorInfo(err))
			c.String(http.StatusInternalServerError, "%s", "fail")
			return
		}
		c.String(http.StatusOK, "%s", "success")
	}
}