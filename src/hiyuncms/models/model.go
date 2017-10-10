package models

import (
	"time"
	"log"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"github.com/go-xorm/core"
	"hiyuncms/config"
)

//var enginex * xorm.DbMaster
var DbMaster * xorm.Engine

/*
db.master.driver=mysql
db.master.dbname=hiyuncms
db.master.user=root
db.master.password=root
db.master.host=localhost:3306
db.master.prefix=hiyuncms_
db.master.encoding=utf8
 */

func init()  {
	driver := config.GetValue("db.master.driver")
	dbname := config.GetValue("db.master.dbname")
	user   := config.GetValue("db.master.user")
	password := config.GetValue("db.master.password")
	host := config.GetValue("db.master.host")
	encode := config.GetValue("db.master.encoding")
	prefix := config.GetValue("db.master.prefix")
	params := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=true", user, password, host, dbname,encode)
	var err error
	DbMaster, err = xorm.NewEngine(driver, params)
	log.Println( "init Database DbMaster ", GetErrorInfo(err))

	DbMaster.SetMaxIdleConns( 50 )
	DbMaster.SetMaxOpenConns( 200 )
	DbMaster.ShowSQL(true)
	tbMapper := core.NewPrefixMapper(core.SnakeMapper{}, prefix)
	DbMaster.SetTableMapper(tbMapper)
}


type PageRequest struct {
	Rows       int `form:"rows"`
	Page       int `form:"page"`
	Sidx 	   string `form:"sidx"`
	Sord	   string `form:"sord"`
	Filters    string `form:"filters"`
}

type PageResponse struct {
	Page int `json:"page"`
	Records  int64 `json:"records"`
	Total int64 `json:"total"`
	Rows  interface{} `json:"rows"`
}

func InitPageResponse(page * PageRequest, list interface{}, records int64 ) *PageResponse {
	pageResponse := PageResponse{}
	pageResponse.Rows = &list
	pageResponse.Page = page.Page
	pageResponse.Records = records
	if page.Rows == 0 {
		page.Rows = 10
	}
	total := records / int64(page.Rows)
	log.Printf("records=%d,  一共%d页\n",records,  total)
	if records % int64(page.Rows) != 0 {
		total +=1
	}
	pageResponse.Total = total

	return &pageResponse
}


func GetErrorInfo(err error)  string{
	if err == nil {
		return "success"
	}else {
		return  err.Error()
	}
}


type Time time.Time
type Date time.Time

const (
	timeFormart = "2006-01-02 15:04:05"
	dateFormart = "2006-01-02"
)



func (t *Time) UnmarshalJSON(data []byte) (err error) {
	now, err := time.ParseInLocation(`"`+timeFormart+`"`, string(data), time.Local)
	*t = Time(now)
	return
}

func (t Time) MarshalJSON() ([]byte, error) {
	b := make([]byte, 0, len(timeFormart)+2)
	b = append(b, '"')
	b = time.Time(t).AppendFormat(b, timeFormart)
	b = append(b, '"')
	return b, nil
}

func (t Time) String() string {
	return time.Time(t).Format(timeFormart)
}

func (t *Date) UnmarshalJSON(data []byte) (err error) {
	now, err := time.ParseInLocation(`"`+dateFormart+`"`, string(data), time.Local)
	*t = Date(now)
	return
}

func (t Date) MarshalJSON() ([]byte, error) {
	b := make([]byte, 0, len(dateFormart)+2)
	b = append(b, '"')
	b = time.Time(t).AppendFormat(b, dateFormart)
	b = append(b, '"')
	return b, nil
}

func (t Date) String() string {
	return time.Time(t).Format(dateFormart)
}

/*
type Session struct {
	xorm.Session
}


func (s * Session) MyLimit(page PageRequest) * Session {
	s.Statement.Limit(page.Rows, page.Rows * page.Page)
	return  s
}

type  DbMaster struct{xorm.DbMaster}

func (e *DbMaster) Where(query interface{}, args ...interface{}) *Session {
	session := e.NewSession()
	session.IsAutoClose = true
	session.Where(query, args...)
	newSession := &Session{*session}
	return newSession
}
*/





