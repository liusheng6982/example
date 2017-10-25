package backend

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"hiyuncms/models/system"
	"hiyuncms/models"
	"hiyuncms/controllers/backend/json"
	"log"
)

func ResourceList(c *gin.Context){
	c.HTML(http.StatusOK, "resourcelist.html", gin.H{
		"bodyCss":"no-skin",
		"mainMenu" :"资源管理",
	})
}

func ResourceListData(c *gin.Context){
	page := models.PageRequest{}
	c.Bind( &page )
	parentIdStr,_:= c.GetPostForm("parentId")
	parentId,_:= strconv.ParseInt(parentIdStr, 10, 64)
	responsePage := system.GetResourceByPage(&page, parentId)
	c.JSON(http.StatusOK, responsePage)
}

/*
组织操作
 */
func ResourceEdit(c * gin.Context){
	org := system.Resource{}
	c.Bind( &org)
	oper, _ := c.GetPostForm("oper")
	if "edit" == oper {
		id, _:= c.GetPostForm("id")
		org.Id, _= strconv.ParseInt(id, 10, 64)
		_, err := models.DbMaster.ID(org.Id).Update(&org)
		if err != nil {
			log.Printf("更新Resource报错:%s\n",models.GetErrorInfo(err))
		}
	}else if "add" == oper {
		_, err := models.DbMaster.Insert( &org)
		if err != nil {
			log.Printf("新增Resource报错:%s\n",models.GetErrorInfo(err))
			c.String(http.StatusInternalServerError, "%s", "fail")
			return
		}
		c.String(http.StatusOK, "%s", "success")
	} else if "del" == oper{
		id, _:= c.GetPostForm("id")
		org.Id, _= strconv.ParseInt(id, 10, 64)
		_, err := models.DbMaster.Delete(&org)
		if err != nil {
			log.Printf("删除Resource报错:%s\n",models.GetErrorInfo(err))
			c.String(http.StatusInternalServerError, "%s", "fail")
			return
		}
		c.String(http.StatusOK, "%s", "success")
	}
}


/**
资源树树的显示
 */
func GetResource(c * gin.Context){
	parentIdStr := c.Query("parentId")
	parentId,_:= strconv.ParseInt(parentIdStr, 10, 64)
	orgs := system.GetResourceByPraentId( parentId )
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
			Name:     v.ResourceName,
			Icon:     icon,
			Children: hasChildren,
			Type:     "1",
		}
		treeNodes[k] = &node
	}
	c.JSON(http.StatusOK, treeNodes)
}