package backend

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"hiyuncms/models"
	"log"
	"strconv"
	"hiyuncms/models/system"
)

func  RoleList(c *gin.Context){
	c.HTML(http.StatusOK, "rolelist.html", gin.H{
		"bodyCss":"no-skin",
		"mainMenu" :"角色管理",
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
			log.Printf("更新Cloun报错:%s\n",models.GetErrorInfo(err))
		}
	}else if"add" == oper {
		_, err := models.DbMaster.Insert( &role )
		if err != nil {
			log.Printf("新增Cloun报错:%s\n",models.GetErrorInfo(err))
		}
	}
}
