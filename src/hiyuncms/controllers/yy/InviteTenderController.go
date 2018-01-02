package yy

import (
	"github.com/gin-gonic/gin"
	"hiyuncms/controllers"
	"net/http"
	"hiyuncms/models"
	"hiyuncms/models/yy"
	"hiyuncms/controllers/frontend"
	"strconv"
	"log"
)

func InviteTenderListShow(c *gin.Context){
	c.HTML(http.StatusOK, "invitetenderlist.html", gin.H{
		"bodyCss":"no-skin",
		"mainMenu" :"招标项目管理",
		"sessionInfo":controllers.GetSessionUser(c),
	})
}

func InviteTenderDetail(c *gin.Context){
	projectId, _ := c.GetQuery("id")
	log.Printf("projectId=%s\n", projectId)
	projectIdInt64, _ := strconv.ParseInt(projectId, 10, 64)
	projectDetail := yy.GetInviteTenderById( projectIdInt64 )
	sessionInfo := frontend.GetSessionInfo(c)
	if sessionInfo.UserName == "" {
		c.Redirect(http.StatusFound, "/userlogin")
		return
	}
	if sessionInfo.VipLevel == 0  {
		c.Redirect(http.StatusFound, "/novip")
		return
	}
	if sessionInfo.VipExpired == 1{
		c.Redirect(http.StatusFound, "/vipexpired")
		return
	}
	log.Printf("%+v", projectDetail)
	c.HTML(http.StatusOK, "projectdetail.html", gin.H{
		"projectDetail" : projectDetail,
		"path":"",
		"sessionInfo":frontend.GetSessionInfo(c),
	})
}

func InviteTenderList(c *gin.Context){
	page := models.PageRequest{}
	c.Bind( &page )
	responsePage := yy.GetAllInviteTenderByPage(&page)
	c.JSON(http.StatusOK, responsePage)
}

func InviteTenderEdit(c * gin.Context){
	tender := yy.YyPorject{}
	c.Bind( &tender)
	tender.ProjectType = 1
	openTenderTime := c.PostForm("InviteOpenTenderTime")
	tender.InviteOpenTenderTime.UnmarshalText( []byte(openTenderTime))
	submitTenderEndTime := c.PostForm("InviteSubmitTenderEndTime")
	tender.InviteSubmitTenderEndTime.UnmarshalText([]byte(submitTenderEndTime))
	oper, _ := c.GetPostForm("oper")
	if "edit" == oper {
		id, _:= c.GetPostForm("id")
		tender.Id, _= strconv.ParseInt(id, 10, 64)
		_, err := models.DbMaster.ID(tender.Id).Update(&tender)
		if err != nil {
			log.Printf("更新Purchase报错:%s\n",models.GetErrorInfo(err))
		}
	}else if "add" == oper {
		tender.CreateTime = models.Date{}
		_, err := models.DbMaster.Insert( &tender)
		if err != nil {
			log.Printf("新增Purchase报错:%s\n",models.GetErrorInfo(err))
			c.String(http.StatusInternalServerError, "%s", "fail")
			return
		}
		c.String(http.StatusOK, "%s", "success")
	} else if "del" == oper{
		id, _:= c.GetPostForm("id")
		tender.Id, _= strconv.ParseInt(id, 10, 64)
		_, err := models.DbMaster.Delete(&tender)
		if err != nil {
			log.Printf("删除Org报错:%s\n",models.GetErrorInfo(err))
			c.String(http.StatusInternalServerError, "%s", "fail")
			return
		}
		c.String(http.StatusOK, "%s", "success")
	}
}