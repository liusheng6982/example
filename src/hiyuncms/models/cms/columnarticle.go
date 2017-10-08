package cms

import (
	"log"
	"hiyuncms/models"
)

func init()  {
	err := models.DbMaster.Sync2( ColumnArticle{})
	log.Println( "init table ColumnArticle ", models.GetErrorInfo(err))
}


type ColumnArticle struct {
	Id              int64     `xorm:"pk autoincr"`
	ArticleId		int64     `xorm:"bigint"`
	ColumnId        int64     `xorm:"bigint"`
}

type ColumnArticleJoin struct {
	ColumnArticle   `xorm:"extends"`
	Article	   		`xorm:"extends"`
	Column	   		`xorm:"extends"`
}
