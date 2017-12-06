package routes

import (
	"hiyuncms/controllers/yy"
	"hiyuncms/controllers/backend"
)

func init(){
	BackendRoute.GET ("/purchaselist", yy.PurchaseListShow)					   //采购项目列表显示
	BackendRoute.POST("/purchaselist",yy.PurchaseList)                            //采购项目列表数据
	BackendRoute.POST("/purchaseedit",backend.OrgEdit) 						   //采购项目操作（增删改）
}
