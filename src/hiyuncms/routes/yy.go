package routes

import (
	"hiyuncms/controllers/yy"
	"hiyuncms/controllers/frontend"
)

func init(){
	BackendRoute.GET( "/getHospital", yy.GetHospital)
	BackendRoute.GET ("/purchaselist", yy.PurchaseListShow)					   //采购项目列表显示
	BackendRoute.POST("/purchaselist",yy.PurchaseList)                           //采购项目列表数据
	BackendRoute.POST("/purchaseedit",yy.PurchaseEdit) 						   //采购项目操作（增删改）

	//InviteTender
	BackendRoute.GET ("/invitetenderlist", yy.InviteTenderListShow)		       //招标项目列表显示
	BackendRoute.POST("/invitetenderlist",yy.InviteTenderList)                   //招标项目列表数据
	BackendRoute.POST("/invitetenderedit",yy.InviteTenderEdit) 				   //招标项目操作（增删改）

	//医院与供应商关系表
	BackendRoute.GET ("/hospitalsupply", yy.HospitalSupplyShow)
	BackendRoute.GET ("/hospitaltree", yy.HospitalTree)
	BackendRoute.POST("/supplylist", yy.SupplyList)
	BackendRoute.GET ("/hospitalsupplydlg", yy.SupplyListShow)
	BackendRoute.GET ("/hospitalsupplysave",yy.HospitalSupplySave)

	FrontendRoute.GET ("/projectdetail",  yy.InviteTenderDetail)
	FrontendRoute.GET ( "/company_projectdetail",  yy.Company_ProjectDetail )    //医院

	FrontendRoute.GET ("/novip",  frontend.NoVip)
	FrontendRoute.GET ("/vipexpired", frontend.VipExpired)

	FrontendRoute.POST("/sendSms", frontend.SendSMS )
	FrontendRoute.POST("/pushPurchasePorject", yy.PushPurchaseProject)
	FrontendRoute.POST( "/notifyProjectNo", yy.PushInviteTenderProject)

	FrontendRoute.POST("/RegistryVerify", frontend.RegistryVerify)

	FrontendRoute.GET ("/aliprepay",    yy.AliPrePay)
	FrontendRoute.POST("/alipaynotify", yy.AliPayNotify)

	FrontendRoute.GET ("/companyindex",    frontend.CompanyIndexShow)

	FrontendRoute.GET ("/hospitalindex", frontend.HospitalIndexShow)
	FrontendRoute.GET ("/hospitallogin", frontend.HospitalLoginShow)
	FrontendRoute.POST("/hospitallogin", frontend.HospitalUserLogin)
	FrontendRoute.GET ("/hospitallogout", frontend.HospitalLogout)
}

