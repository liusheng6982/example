package yy

import (
	"hiyuncms/config"
)

var PayProductInfo map[int64]*PaymentInfo

type PaymentInfo struct{
	Id int64
	PayInfo string
	PayAmount int64
}

func init()  {
	PayProductInfo = make(map[int64]*PaymentInfo)
	PayProductInfo[1] = &PaymentInfo{Id:1, PayInfo:"普通会员年费",PayAmount:int64(config.GetInt("vip.fee.normal"))}
	PayProductInfo[2] = &PaymentInfo{Id:2, PayInfo:"优选会员年费",PayAmount:int64(config.GetInt("vip.fee.excellent"))}
	PayProductInfo[3] = &PaymentInfo{Id:3, PayInfo:"VIP会员年费",PayAmount:int64(config.GetInt("vip.fee.top"))}
	PayProductInfo[4] = &PaymentInfo{Id:4, PayInfo:"现有供应商普通会员年费",PayAmount:int64(config.GetInt("vip.supply.fee.normal"))}
	PayProductInfo[5] = &PaymentInfo{Id:5, PayInfo:"现有供应商优选会员年费",PayAmount:int64(config.GetInt("vip.supply.fee.excellent"))}
	PayProductInfo[6] = &PaymentInfo{Id:6, PayInfo:"现有供应商VIP会员年费",PayAmount:int64(config.GetInt("vip.supply.fee.top"))}
}

func GetPayInfo(vipLevel int64)  *PaymentInfo{
	paymentInfo := PayProductInfo[vipLevel]
	return paymentInfo
}

func GetAllPayInfo(vipLevel int64)  map[int64]*PaymentInfo{
	return PayProductInfo
}
