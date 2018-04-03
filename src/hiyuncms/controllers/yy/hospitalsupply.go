package yy

import (
	"github.com/gin-gonic/gin"
	"hiyuncms/controllers"
	"net/http"
	"hiyuncms/controllers/backend/json"
	"hiyuncms/models/yy"
	"strconv"
	"hiyuncms/models"
	"fmt"
	"strings"
)

func HospitalSupplyShow(c *gin.Context)  {
	c.HTML(http.StatusOK, "hospitalsupply.html", gin.H{
		"mainMenu":"医院供应商关系",
		"bodyCss" : "no-skin",
		"user":controllers.GetSessionUser(c),
	})
}

func HospitalTree(c *gin.Context)  {

	hospitals := yy.GetHospital( )
	treeNodes := make([]*json.TreeNode, len(hospitals))
	for k,v := range hospitals {
		//tempOrgs := system.GetSubOrgByPraentId( v.Id )
		hasChildren := false
		icon := "ace-icon ace-icon fa fa-folder-o blue"
		//if tempOrgs != nil && len(tempOrgs) > 0 {
		//	hasChildren = true
		//	icon = "ace-icon ace-icon fa fa-folder blue"
		//}
		node := json.TreeNode{
			Id:       v.Id,
			Name:     v.CompanyName,
			Icon:     icon,
			Children: hasChildren,
			Type:     "1",
		}
		treeNodes[k] = &node
	}
	c.JSON(http.StatusOK, treeNodes)
}

func SupplyList(c *gin.Context)  {
	page := models.PageRequest{}
	c.Bind( &page )
	parentIdStr := c.PostForm("parentId")
	parentId,_:= strconv.ParseInt(parentIdStr, 10, 64)
	c.JSON( http.StatusOK, yy.GetSuppliesByHospitalId(parentId,&page) )
}

/**
设置用户角色面板
 */
func SupplyListShow (c * gin.Context){
	parentIdStr := c.Query("parentId")
	hospitalId, _ :=  strconv.ParseInt(parentIdStr, 10, 64)

	supplies := yy.GetAllSupplies()
	options := ""
	for _,supply := range supplies{
		if yy.IsSelectSupplyByHospitalId(hospitalId, supply.Id) {
			options = fmt.Sprintf("%s<option value='%d' selected='selected'> %s </option>", options, supply.Id,supply.CompanyName)
		} else {
			options = fmt.Sprintf("%s<option value='%d'>%s</option>", options, supply.Id,supply.CompanyName)
		}
	}
	c.JSON(http.StatusOK, options)
}

func HospitalSupplySave(c * gin.Context)  {

	hospitalIdStr := c.Query("hospitalId")
	hospitalId, _ :=  strconv.ParseInt(hospitalIdStr, 10, 64)
	supplyIdsStr := c.Query("supplies")
	supplyIds := strings.Split(supplyIdsStr,",")
	supplyIdsInt := make([]int64, len(supplyIds))
	for k,v := range supplyIds {
		supplyIdsInt[k],_ = strconv.ParseInt(v, 10,64)
	}
	yy.HospitalSupplySave(hospitalId, supplyIdsInt)
	c.JSON(http.StatusOK, gin.H{
		"result":"success",
	})
}





