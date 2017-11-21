package yy

import "hiyuncms/models"

type  YyTradeInfo struct{
	CompanyId     string      `xorm:"varchar(20)"`

	TradeDate     models.Date  `xorm:"datetime"`
	TradeAmount   float64     `xorm:"double"`

}
