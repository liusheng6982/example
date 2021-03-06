package backend

import (
	"github.com/gin-gonic/gin"
	"hiyuncms/models/system"
	"strconv"
	"hiyuncms/controllers/backend/json"
	"net/http"
	"hiyuncms/models"
	"log"
	"hiyuncms/controllers"
)

func OrgList(c *gin.Context){
	c.HTML(http.StatusOK, "orglist.html", gin.H{
		"bodyCss":"no-skin",
		"mainMenu" :"组织管理",
		"user":controllers.GetSessionUser(c),
	})
}

/**
列表数据
 */
func OrgListData(c *gin.Context){
	page := models.PageRequest{}
	c.Bind( &page )
	parentIdStr := c.PostForm("parentId")
	parentId,_:= strconv.ParseInt(parentIdStr, 10, 64)
	c.JSON( http.StatusOK, system.GetSubOrgByPage(parentId,&page) )
}

/*
组织操作
 */
func OrgEdit(c * gin.Context){
	org := system.Org{}
	c.Bind( &org)
	oper, _ := c.GetPostForm("oper")
	if "edit" == oper {
		id, _:= c.GetPostForm("id")
		org.Id, _= strconv.ParseInt(id, 10, 64)
		_, err := models.DbMaster.ID(org.Id).Update(&org)
		if err != nil {
			log.Printf("更新Org报错:%s\n",models.GetErrorInfo(err))
		}
	}else if "add" == oper {
		_, err := models.DbMaster.Insert( &org)
		if err != nil {
			log.Printf("新增Org报错:%s\n",models.GetErrorInfo(err))
			c.String(http.StatusInternalServerError, "%s", "fail")
			return
		}
		c.String(http.StatusOK, "%s", "success")
	} else if "del" == oper{
		id, _:= c.GetPostForm("id")
		org.Id, _= strconv.ParseInt(id, 10, 64)
		_, err := models.DbMaster.Delete(&org)
		if err != nil {
			log.Printf("删除Org报错:%s\n",models.GetErrorInfo(err))
			c.String(http.StatusInternalServerError, "%s", "fail")
			return
		}
		c.String(http.StatusOK, "%s", "success")
	}
}

/**
树的显示
 */
func GetSubOrg(c * gin.Context){
	parentIdStr := c.Query("parentId")
	parentId,_:= strconv.ParseInt(parentIdStr, 10, 64)
	orgs := system.GetSubOrgByPraentId( parentId )
	treeNodes := make([]*json.TreeNode, len(orgs))
	for k,v := range orgs{
		tempOrgs := system.GetSubOrgByPraentId( v.Id )
		hasChildren := false
		icon := "ace-icon ace-icon fa fa-folder-o blue"
		if tempOrgs != nil && len(tempOrgs) > 0 {
			hasChildren = true
			icon = "ace-icon ace-icon fa fa-folder blue"
		}
		node := json.TreeNode{
			Id:       v.Id,
			Name:     v.OrgName,
			Icon:     icon,
			Children: hasChildren,
			Type:     "1",
		}
		treeNodes[k] = &node
	}
	c.JSON(http.StatusOK, treeNodes)
}


