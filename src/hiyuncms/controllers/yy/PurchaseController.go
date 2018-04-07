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
		"sessionInfo":controllers.GetSessionUser(c),
	})
}

func PurchaseList(c *gin.Context){
	page := models.PageRequest{}
	c.Bind( &page )
	responsePage := yy.GetAllYyPurchaseByPage(&page)
	c.JSON(http.StatusOK, responsePage)
}

func PurchaseEdit(c * gin.Context){
	purchase := yy.YyPorject{}
	purchase.ProjectType = 2
	bindErr := c.Bind( &purchase)
	temp := yy.GetById( purchase.CompanyId )
	purchase.CompanyName =  temp.CompanyName
	quotePriceEndDate := c.PostForm("PurchaseQuotePriceEndTime")
	convertErr := purchase.PurchaseQuotePriceEndTime.UnmarshalText( []byte(quotePriceEndDate))
	if convertErr != nil {
		log.Printf("quotePriceEndDate 绑定数据出错:%s\n",models.GetErrorInfo(convertErr))
	}
	purchaseDeliveryTime := c.PostForm("PurchaseDeliveryTime")
	convert2Err := purchase.PurchaseDeliveryTime.UnmarshalText( []byte(purchaseDeliveryTime))
	if convert2Err != nil {
		log.Printf("quotePriceEndDate 绑定数据出错:%s\n",models.GetErrorInfo(convert2Err))
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

func PushPurchaseProject( c * gin.Context ){
	type YyPorjectTemp struct {
		ProjectName    						string			`form:"projectName"`
		ProjectNo      						string			`form:"projectNo"`    //项目编号
		ProjectContent 						string			`form:"projectContent"`  //内容
		CompanyId         					int64    	    `form:"companyId"`         //公司ID
		CompanyName       					string   	    `form:"companyName"`    //公司名称
		ContactPhone   						string 	   		`form:"contactPhone"`    //联系人电话
		Contact	   	   						string          `form:"contact"`    	//联系人
		PurchaseType   						string			`form:"varchar(20)"`   //合格供应商，定向，公开
		PurchaseExpiredDate                 models.Date 	`form:"DateTime"`      //采购有效期
		PurchaseQuotePriceEndTime   		models.Date		`form:"DateTime"`      //报价截止时间
		PurchaseDeliveryTime                models.Date		`form:"DateTime"`      //交货时间
		BusinessCategory					string			`form:"varchar(20)"`   //建设、理疗器械、后勤物资、行政物资
	}
	purchase := yy.YyPorject{}
	purchase.ProjectType = 2
	purchase.ImpFlag = 1
	purchaseTemp := YyPorjectTemp{}
	bindErr := c.Bind( &purchaseTemp)
	if bindErr != nil {
		log.Printf("bind form出错:%s\n",models.GetErrorInfo(bindErr))
	}
	models.CopyStruct(purchaseTemp,purchase)
	_, err := models.DbMaster.Insert( &purchase)
	if err != nil {
		log.Printf("bind form出错:%s\n",models.GetErrorInfo(bindErr))
		c.JSON(http.StatusOK, gin.H{
			"success":"false",
			"msg":"推送失败",
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"success":"true",
		"msg":"推送成功",
	})
}