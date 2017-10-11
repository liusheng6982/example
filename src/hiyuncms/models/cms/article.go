package cms

import (
	"log"
	"hiyuncms/models"
	"fmt"
	"strconv"
)

func init()  {
	err := models.DbMaster.Sync2( Article{})
	log.Println( "init table Article ", models.GetErrorInfo(err))
}



type Article struct {
	Id             int64      `xorm:"pk autoincr"`
	Title          string     `xorm:"varchar(80)"`
	Content        string     `xorm:"text"`
	Copyfrom       string     `xorm:"varchar(100)"`
	Keywords       string     `xorm:"varchar(255)"`
	Description    string     `xorm:"varchar(255)"`
	Status         int64      `xorm:"int(1)"`
	Createtime     models.Time     `xorm:"DateTime"`
	Updatetime     models.Time     `xorm:"DateTime"`

	ColumnNames    string 	  `xorm:"varchar(1000)"`
}

/*
新增文档
 */
func SaveArticle(article *Article, columnIds []string)  {

	after := func(bean interface{}){
		tempArticle := bean.(*Article)
		articleId := tempArticle.Id
		models.DbMaster.Delete(ColumnArticle{ArticleId:articleId})
		columnNames := ""
		for _, columnId := range columnIds {
			ca := ColumnArticle{}
			ca.ArticleId = articleId
			ca.ColumnId ,_ = strconv.ParseInt( columnId, 0, 64 )
			models.DbMaster.Insert( ca )
			column := Column{}
			models.DbMaster.Id(ca.ColumnId).Get(&column)
			columnNames = fmt.Sprintf("%s,%s", columnNames, column.Name)
		}
		article.ColumnNames = columnNames
		_,err := models.DbMaster.Id(article.Id).Update( article )
		if err != nil  {
			log.Printf("保存Article的栏目名:%s", models.GetErrorInfo(err))
		}
	}
	_,err := models.DbMaster.After(after).Insert( article )
	if err != nil {
		log.Printf("保存Article数据:%s", models.GetErrorInfo(err))
	}
}

/**
获取所有文章
 */
func  GetAllArticles(page *models.PageRequest) *models.PageResponse{
	articleList := make([]Article, 0)
	err := models.DbMaster.Table(Article{}).Limit(page.Rows, (page.Page - 1)* page.Rows).Find(&articleList)
	if err != nil {
		log.Printf("获取Article数据:%s", models.GetErrorInfo(err))
	}
	records ,_ := models.DbMaster.Table(Article{}).Count(Article{})
	pageResponse := models.InitPageResponse(page, &articleList, records)
	return pageResponse

}

/**
通过Column的URL获取Article
 */
func GetArticlesByPath(page *models.PageRequest, path string) * models.PageResponse{
	articles_ := make([]Article, 0)
	log.Printf("%v", page)
	err := models.DbMaster.Table(Article{}).Alias("a").
		Limit(page.Rows, (page.Page - 1) * page.Rows).
		Join("INNER", []string{"hiyuncms_column_article","ca"}, "a.id=ca.article_id").
		Join("INNER", []string{"hiyuncms_column" ,"c"},"c.id=ca.column_id and c.url='"+ path +"'").
		Where(" a.status=1").
		Find(&articles_)
	if err != nil {
		log.Printf("通过Column的URL获取Article数据:%s", models.GetErrorInfo(err))
	}
	records,_ :=  models.DbMaster.Table(Article{}).Alias("a").
		Join("INNER", []string{"hiyuncms_column_article","ca"}, "a.id=ca.article_id").
		Join("INNER", []string{"hiyuncms_column" ,"c"},"c.id=ca.column_id and c.url='"+ path +"'").
		Where(" a.status=1").
		Count(Article{})

	pageResponse := models.InitPageResponse(page, &articles_, records)
	return  pageResponse

}


