package backend

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"hiyuncms/models"
	"log"
	"strconv"
	"hiyuncms/models/system"
	"hiyuncms/controllers/backend/json"
	"strings"
	"hiyuncms/controllers"
)

func  RoleList(c *gin.Context){
	c.HTML(http.StatusOK, "rolelist.html", gin.H{
		"bodyCss":"no-skin",
		"mainMenu" :"角色管理",
		"user":controllers.GetSessionUser(c),

	})
}

func RoleDataList(c *gin.Context){
	page := models.PageRequest{}
	c.Bind( &page )
	responsePage := system.GetAllRolesByPage(&page)
	c.JSON(http.StatusOK, responsePage)
}


func RoleEdit(c * gin.Context){
	role := system.Role{}
	c.Bind( &role )
	oper, _ := c.GetPostForm("oper")
	if "edit" == oper {
		id, _:= c.GetPostForm("id")
		role.Id, _= strconv.ParseInt(id, 10, 64)
		_, err := models.DbMaster.ID(role.Id).Update(&role)
		if err != nil {
			log.Printf("更新Role报错:%s\n",models.GetErrorInfo(err))
		}
	}else if"add" == oper {
		_, err := models.DbMaster.Insert( &role )
		if err != nil {
			log.Printf("新增Role报错:%s\n",models.GetErrorInfo(err))
		}
	} else if "del" == oper {
		id, _:= c.GetPostForm("id")
		role.Id, _= strconv.ParseInt(id, 10, 64)
		_, err := models.DbMaster.Delete( role )
		if err != nil {
			log.Printf("删除Role报错:%s\n",models.GetErrorInfo(err))
		}

	}
}

func RoleResource(c * gin.Context){
	id, _:= c.GetPostForm("roleId")
	roleId, _ := strconv.ParseInt(id, 10, 64)
	resources := system.GetResourceByRole( roleId )
	c.JSON(http.StatusOK, resources)
}

/**
角色资源保存
 */
func RoleResourceSave(c * gin.Context){
	roleIdStr := c.Query("roleId")
	roleId, _ :=  strconv.ParseInt(roleIdStr, 10, 64)
	resourceIdsStr := c.Query("resourceIds")
	resourceIds := strings.Split(resourceIdsStr,",")
	resourceIdsInt := make([]int64, len(resourceIds))
	for k,v := range  resourceIds{
		resourceIdsInt[k],_ = strconv.ParseInt(v, 10,64)
	}
	system.RoleResourceSave(roleId, resourceIdsInt)
	c.JSON(http.StatusOK, gin.H{
		"result":"success",
	})
}

/**
资源树
 */
func RoleResourceTree(c * gin.Context){
	parentIdStr := c.Query("parentId")
	roleIdStr := c.Query("roleId")

	parentId,_:= strconv.ParseInt(parentIdStr, 10, 64)
	roleId, _ :=  strconv.ParseInt(roleIdStr, 10, 64)

	orgs := system.GetResourceByPraentId( parentId )
	treeNodes := make([]*json.TreeNode, len(orgs))
	for k,v := range orgs{
		hasSelected := system.IsSelectResourceByRoleId(roleId, v.Id)
		tempOrgs := system.GetResourceByPraentId( v.Id )
		hasChildren := false

		icon := "ace-icon ace-icon fa fa-folder-o blue"
		if tempOrgs != nil && len(tempOrgs) > 0 {
			hasChildren = true
			icon = "ace-icon ace-icon fa fa-folder blue"
		}
		state := map[string]interface{}{"selected":hasSelected}

		node := json.TreeNode{
			Id:       v.Id,
			Name:     v.ResourceName,
			Icon:     icon,
			Children: hasChildren,
			Type:     "1",
			State:    state,
		}
		treeNodes[k] = &node
	}
	c.JSON(http.StatusOK, treeNodes)
}
