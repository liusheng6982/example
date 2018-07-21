package yy

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"hiyuncms/controllers"
	"hiyuncms/models"
	"hiyuncms/models/yy"
	"github.com/360EntSecGroup-Skylar/excelize"
	"fmt"
	"log"
	"hiyuncms/controllers/backend"
)

/**
显示编辑页面
 */
func CompanyShow( c *gin.Context )  {
	c.HTML(http.StatusOK, "supply-list.html", gin.H{
		"mainMenu":"公司列表",
		"bodyCss" : "no-skin",
		"user":controllers.GetSessionUser(c),
	})
}

/**
公司列表
 */
func CompanyList(c *gin.Context){
	page := models.PageRequest{}
	c.Bind( &page )
	c.JSON( http.StatusOK, yy.GetSupplies(&page) )
}

/**
导入供应商
 */
func SupplyImport(c * gin.Context){
	file, _, err := c.Request.FormFile("excel-file")
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"impsuccess":false,
			"msg":"无效的请求！",
		})
	}
	xlsx, err := excelize.OpenReader(file)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"impsuccess":false,
			"msg":"无效的Excel文件！",
		})
	}
	rows := xlsx.GetRows("Sheet1")
	for k, row := range rows {
		if k == 0 {
			continue
		}
		company := yy.YyCompany{}
		has, err := models.DbSlave.Where("Company_Name=?", row[0]).Get(&company)
		if err != nil {
			log.Printf("导入时，查询供应商报错:%s", row[0])
		}
		if !has {
			company.CompanyName =  row[0]
			company.CompanyType = "2"
			models.DbMaster.Insert( &company )
		}

		user := yy.YyUser{}
		has, err = models.DbSlave.Where("User_Phone=?", row[2]).Get(&user)
		if err != nil {
			log.Printf("导入时，查询用户报错：%s", row[2])
		}
		if has {
			log.Printf("用户已存在：%",  row[2])
		}else{
			user.UserName = row[ 1 ]
			user.UserPhone = row[ 2 ]
			user.UserPassword = backend.Md5str( row[ 3] )
			user.CompanyId = company.Id

			models.DbMaster.Insert( &user )
		}
		fmt.Printf("%s,%s,%s", row[0], row[1], row[2])
	}
	c.JSON(http.StatusOK, gin.H{
		"impsuccess":true,
	})
}