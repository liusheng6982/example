package routes

import (
	"hiyuncms/controllers/yy"
	"hiyuncms/controllers/frontend"
)

func init(){
	BackendRoute.GET ("/purchaselist", yy.PurchaseListShow)					   //采购项目列表显示
	BackendRoute.POST("/purchaselist",yy.PurchaseList)                           //采购项目列表数据
	BackendRoute.POST("/purchaseedit",yy.PurchaseEdit) 						   //采购项目操作（增删改）

	//InviteTender
	BackendRoute.GET ("/invitetenderlist", yy.InviteTenderListShow)		       //招标项目列表显示
	BackendRoute.POST("/invitetenderlist",yy.InviteTenderList)                   //招标项目列表数据
	BackendRoute.POST("/invitetenderedit",yy.InviteTenderEdit) 				   //招标项目操作（增删改）
	
	FrontendRoute.GET ("/projectdetail",  yy.InviteTenderDetail)

	FrontendRoute.GET ("/novip",  frontend.NoVip)
	FrontendRoute.GET ("/vipexpired", frontend.VipExpired)
}

