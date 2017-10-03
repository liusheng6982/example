package models

import (
	"time"
	"log"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"github.com/go-xorm/core"
)

var enginex * xorm.Engine
var engine * Engine
func init()  {

	log.Println( "init Database engine ")
	//driver, _ := c.String("database", "db.driver")
	//dbname, _ := c.String("database", "db.dbname")
	//user, _ := c.String("database", "db.user")
	//password, _ := c.String("database", "db.password")
	//host, _ := c.String("database", "db.host")
	//params := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=true", user, password, host, dbname)

	enginex, _ = xorm.NewEngine("mysql", "root:root@/ADC?charset=utf8")

	engine = &Engine{*enginex}

	engine.SetMaxIdleConns( 50 )
	engine.SetMaxOpenConns( 200 )
	engine.ShowSQL(true)
	tbMapper := core.NewPrefixMapper(core.SnakeMapper{}, "sys_")
	engine.SetTableMapper(tbMapper)
}

type Page struct {
	PageSize int `json:"pageSize"`
	PageNum  int `json:"pageNo"`
}


func getErrorInfo(err error)  string{
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


type Session struct {
	xorm.Session
}

func (this * Session) MyLimit(page Page) * Session {
	this.Statement.Limit(page.PageSize, page.PageSize * page.PageNum)
	return  this
}

type  Engine struct{xorm.Engine}

func (this *Engine) Where(query interface{}, args ...interface{}) *Session {
	session := this.NewSession()
	session.IsAutoClose = true
	session.Where(query, args...)
	newSession := &Session{*session}
	return newSession
}




