package yy

var PayProductInfo map[int64]*PaymentInfo

type PaymentInfo struct{
	Id int64
	PayInfo string
	PayAmount int64
}

func init()  {
	PayProductInfo = make(map[int64]*PaymentInfo)
	PayProductInfo[1] = &PaymentInfo{Id:1, PayInfo:"普通会员年费",PayAmount:120000}
	PayProductInfo[2] = &PaymentInfo{Id:2, PayInfo:"优选会员年费",PayAmount:180000}
	PayProductInfo[3] = &PaymentInfo{Id:3, PayInfo:"VIP会员年费",PayAmount:600000}
	PayProductInfo[4] = &PaymentInfo{Id:4, PayInfo:"现有供应商普通会员年费",PayAmount:60000}
	PayProductInfo[5] = &PaymentInfo{Id:5, PayInfo:"现有供应商优选会员年费",PayAmount:120000}
	PayProductInfo[6] = &PaymentInfo{Id:6, PayInfo:"现有供应商VIP会员年费",PayAmount:480000}
}

func GetPayInfo(vipLevel int64)  *PaymentInfo{
	paymentInfo := PayProductInfo[vipLevel]
	return paymentInfo
}
