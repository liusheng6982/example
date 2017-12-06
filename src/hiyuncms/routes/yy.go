package routes

import (
	"hiyuncms/controllers/yy"
)

func init(){
	BackendRoute.GET ("/purchaselist", yy.PurchaseListShow)					   //首页
}
