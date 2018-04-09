package frontend

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"hiyuncms/models/cms"
	"time"
	"hiyuncms/models/yy"
	"log"
	"github.com/gin-gonic/contrib/sessions"
	"encoding/json"
	"strconv"
)

func ArticlesShow( c *gin.Context )  {
	pageNoStr := c.Query("pageNo")
	if pageNoStr == "" {
		pageNoStr = "1"
	}
	pageSizeStr := c.Query("pageSize")
	if pageSizeStr == "" {
		pageSizeStr = "10"
	}
	path := c.Request.URL.Path
	column :=  cms.GetColumnByPath( path )  //查询出模板路径

	{
		userSessionInfo:= GetSessionInfo(c)
		if userSessionInfo != nil {
			company := yy.GetById(userSessionInfo.CompanyId)
			now := time.Now()
			expired := 1
			if now.Before(time.Time(company.VipExpired)) || now.Equal(time.Time(company.VipExpired)) {
				expired = 0
			}
			log.Printf("用户会员信息是否过期，expired： 1是过期 0是未过期 expired=%d\n", expired)
			userSessionInfo.VipExpired = expired
			session := sessions.Default(c)
			jsonBytes, _ := json.Marshal(userSessionInfo)
			session.Set(FRONT_USER_SESSION, string(jsonBytes))
			session.Save()
		}
	}

	pageNo,_:= strconv.Atoi(pageNoStr)
	pageSize,_ := strconv.Atoi(pageSizeStr)
	c.HTML(http.StatusOK, column.TemplatePath, gin.H{
		"path":path,
		"pageNo":pageNo,
		"pageSize":pageSize ,
		"sessionInfo":GetSessionInfo(c),
	})
}

func ArticleShow( c *gin.Context )  {
	articleId := c.Query("articleId")
	c.HTML(http.StatusOK, "articleshow.html", gin.H{
		"articleId":articleId,
		"path":"",
		"sessionInfo":GetSessionInfo(c),
	})
}
