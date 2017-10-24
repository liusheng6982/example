package backend

import (
	"github.com/gin-gonic/gin"
	"hiyuncms/models/system"
	"strconv"
	"hiyuncms/controllers/backend/json"
	"net/http"
	"hiyuncms/models"
)

func OrgList(c *gin.Context){
	c.HTML(http.StatusOK, "orglist.html", gin.H{
		"bodyCss":"no-skin",
		"mainMenu" :"组织管理",
	})
}

func OrgListData(c *gin.Context){
	page := models.PageRequest{}
	c.Bind( &page )
	parentIdStr := c.Query("parentId")
	parentId,_:= strconv.ParseInt(parentIdStr, 10, 64)
	c.JSON( http.StatusOK, system.GetSubOrgByPage(parentId,&page) )
}

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


