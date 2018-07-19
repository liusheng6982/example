package cms

import (
	"log"
	"hiyuncms/models"
	"github.com/opesun/goquery"
	"fmt"
	"strconv"
	"time"
	"unicode/utf8"
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

	FirstImage	   string
	FirstContent   string
}

/*
新增文档
 */
func SaveArticle(article *Article, columnIds []string)  {

	article.Createtime = models.Time(time.Now())

	after := func(bean interface{}){
		/*
		tempArticle := bean.(*Article)
		articleId := tempArticle.Id
		models.DbMaster.Delete(&ColumnArticle{ArticleId:articleId})
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
		*/
	}
	if article.Id == 0   {
		_, err := models.DbMaster.After(after).Insert(article)
		if err != nil {
			log.Printf("保存Article数据:%s", models.GetErrorInfo(err))
		}
	}else {
		_, err := models.DbMaster.After(after).ID(article.Id).Update(article)
		if err != nil {
			log.Printf("更新Article数据:%s", models.GetErrorInfo(err))
		}
	}


	articleId := article.Id
	models.DbMaster.Delete(&ColumnArticle{ArticleId:articleId})
	columnNames := ""
	for index, columnId := range columnIds {
		ca := ColumnArticle{}
		ca.ArticleId = articleId
		ca.ColumnId ,_ = strconv.ParseInt( columnId, 0, 64 )
		models.DbMaster.Insert( &ca )
		column := Column{}
		models.DbMaster.Id(ca.ColumnId).Get(&column)
		if index == 0 {
			columnNames = fmt.Sprintf("%s", column.Name)
		}else {
			columnNames = fmt.Sprintf("%s,%s", columnNames, column.Name)
		}
	}
	article.ColumnNames = columnNames
	_,err := models.DbMaster.Id(article.Id).Update( article )
	if err != nil  {
		log.Printf("保存Article的栏目名:%s", models.GetErrorInfo(err))
	}

}

/**
获得单一文章
 */
func GetArticle(articleId int64) *Article{
	article := Article{Id: articleId}
	models.DbSlave.ID( articleId ).Get( &article )
	return &article
}

/**
获取所有文章
 */
func  GetAllArticles(page *models.PageRequest) *models.PageResponse{
	articleList := make([]Article, 0)
	err := models.DbSlave.Table(Article{}).OrderBy("Createtime desc").Limit(page.Rows, (page.Page - 1)* page.Rows).Find(&articleList)
	if err != nil {
		log.Printf("获取Article数据:%s", models.GetErrorInfo(err))
	}
	records ,_ := models.DbSlave.Table(Article{}).Count(Article{})
	pageResponse := models.InitPageResponse(page, &articleList, records)
	return pageResponse

}

/**
通过Column的URL获取Article
 */
func GetArticlesByPath(page *models.PageRequest, path string) * models.PageResponse{
	articles_ := make([]*Article, 0)
	//log.Printf("%v", page)
	err := models.DbSlave.Table(Article{}).Alias("a").
		Limit(page.Rows, (page.Page - 1) * page.Rows).
		Join("INNER", []string{"hiyuncms_column_article","ca"}, "a.id=ca.article_id").
		Join("INNER", []string{"hiyuncms_column" ,"c"},"c.id=ca.column_id and c.url='"+ path +"'").
		Where(" a.status=1").
		Find(&articles_)
	if err != nil {
		log.Printf("通过Column的URL获取Article数据:%s", models.GetErrorInfo(err))
	}
	records,_ :=  models.DbSlave.Table(Article{}).Alias("a").
		Join("INNER", []string{"hiyuncms_column_article","ca"}, "a.id=ca.article_id").
		Join("INNER", []string{"hiyuncms_column" ,"c"},"c.id=ca.column_id and c.url='"+ path +"'").
		Where(" a.status=1").
		Count(Article{})

	initFirstContent( articles_ )

	pageResponse := models.InitPageResponse(page, &articles_, records)
	return  pageResponse

}

func initFirstContent(articles_ []*Article)  {
	for _, artic := range articles_ {
		nodes, _ :=goquery.ParseString( artic.Content )

		artic.FirstImage= nodes.Find("img").First().Attr("src")

		//nodes.Find("p").First().
		str := nodes.Find("p").Text()
		if utf8.RuneCountInString(str) > 150 {
			artic.FirstContent = fmt.Sprintf("%s%s",string([]rune(str)[0:150]), "...")
		}else{
			artic.FirstContent = str
		}
	}
}

func GetArticlesByPathTop(path string, begin, end int ) []*Article{
	articles_ := make([]*Article, 0)
	//log.Printf("%v", page)
	err := models.DbSlave.Table(Article{}).Alias("a").
		Limit(end-begin, begin).
		Join("INNER", []string{"hiyuncms_column_article","ca"}, "a.id=ca.article_id").
		Join("INNER", []string{"hiyuncms_column" ,"c"},"c.id=ca.column_id and c.url='"+ path +"'").
		Where(" a.status=1").
		Find(&articles_)
	if err != nil {
		log.Printf("通过Column的URL获取Article数据:%s", models.GetErrorInfo(err))
	}
	initFirstContent( articles_ )
	return  articles_

}

/**
删除
 */
func DeleteArticle(articleId int64)  {
	article := Article{Id:articleId}
	models.DbMaster.Delete( &article )
	models.DbMaster.Delete(&ColumnArticle{ArticleId:articleId})
}

/**
发布
 */
func PublishArticle(articleId int64)  {
	article := Article{Id:articleId, Status:1}
	_, err := models.DbMaster.Id( articleId ).Update(&article)
	if err != nil {
		log.Printf("发布Article出错:%s", models.GetErrorInfo(err))
	}
}

/**
撤销发布
 */
func PublishCancelArticle(articleId int64)  {
	article := Article{Id:articleId,Status:2}
	_, err := models.DbMaster.ID(articleId).Update(&article)
	if err != nil {
		log.Printf("撤回Article出错%d:%s", articleId, models.GetErrorInfo(err))
	}
}





