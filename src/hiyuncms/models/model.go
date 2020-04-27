package models

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
	_ "github.com/mattn/go-sqlite3"
	"hiyuncms/config"
	"log"
	"time"
)

//var enginex * xorm.DbMaster
var DbMaster *xorm.Engine
var DbSlave *xorm.Engine

/*
db.master.driver=mysql
db.master.dbname=hiyuncms
db.master.user=root
db.master.password=root
db.master.host=localhost:3306
db.master.prefix=hiyuncms_
db.master.encoding=utf8
*/

func init() {
	log.Println( config.GetValue("db.type") )
	prefix := config.GetValue("db.master.prefix")
	if config.GetValue("db.type") == "mysql" {
		driver := config.GetValue("db.master.driver")
		dbname := config.GetValue("db.master.dbname")
		user := config.GetValue("db.master.user")
		password := config.GetValue("db.master.password")
		host := config.GetValue("db.master.host")
		encode := config.GetValue("db.master.encoding")

		params := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=true&allowNativePasswords=true", user, password, host, dbname, encode)
		var err error
		DbMaster, err = xorm.NewEngine(driver, params)
		log.Println("init Database DbMaster ", params,GetErrorInfo(err))

		maxIdle := config.GetInt("db.master.max.idle")
		maxConn := config.GetInt("db.master.max.conn")
		DbMaster.SetMaxIdleConns(maxIdle)
		DbMaster.SetMaxOpenConns(maxConn)
		showSql := config.GetBool("db.master.show.sql")
		DbMaster.ShowSQL(showSql)
		tbMapper := core.NewPrefixMapper(core.SnakeMapper{}, prefix)
		DbMaster.SetTableMapper(tbMapper)
		//cacher := xorm.NewLRUCacher(xorm.NewMemoryStore(), 1000)
		//DbMaster.SetDefaultCacher(cacher)

		initSlave()
	} else if config.GetValue("db.type") == "sqllite" {
		println("into sqllite")
		var err error
		DbMaster, err = xorm.NewEngine("sqlite3", "./hiyum.db")
		if err != nil {
			log.Println("init Database DbSlave ", GetErrorInfo(err))
		}
		DbSlave = DbMaster
		tbMapper := core.NewPrefixMapper(core.SnakeMapper{}, prefix)
		DbMaster.SetTableMapper(tbMapper)
		println("%+v", DbSlave)
		println("%+v", DbMaster)
	}
}

func initSlave() {
	driver := config.GetValue("db.slave.driver")
	dbname := config.GetValue("db.slave.dbname")
	user := config.GetValue("db.slave.user")
	password := config.GetValue("db.slave.password")
	host := config.GetValue("db.slave.host")
	encode := config.GetValue("db.slave.encoding")
	prefix := config.GetValue("db.slave.prefix")
	params := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=%s&parseTime=true", user, password, host, dbname, encode)
	var err error
	DbSlave, err = xorm.NewEngine(driver, params)
	log.Println("init Database DbSlave ", GetErrorInfo(err))

	maxIdle := config.GetInt("db.slave.max.idle")
	maxConn := config.GetInt("db.slave.max.conn")
	DbSlave.SetMaxIdleConns(maxIdle)
	DbSlave.SetMaxOpenConns(maxConn)
	showSql := config.GetBool("db.slave.show.sql")
	DbSlave.ShowSQL(showSql)
	tbMapper := core.NewPrefixMapper(core.SnakeMapper{}, prefix)
	DbSlave.SetTableMapper(tbMapper)

	//cacher := xorm.NewLRUCacher(xorm.NewMemoryStore(), 1000)
	//DbSlave.SetDefaultCacher(cacher)
}

type PageRequest struct {
	Rows    int    `form:"rows"`
	Page    int    `form:"page"`
	Sidx    string `form:"sidx"`
	Sord    string `form:"sord"`
	Filters string `form:"filters"`
}

type PageResponse struct {
	Page     int         `json:"page"`
	Records  int64       `json:"records"`
	Total    int         `json:"total"`
	Rows     interface{} `json:"rows"`
	PageSize int
	Path     string
}

func InitPageResponse(page *PageRequest, list interface{}, records int64) *PageResponse {
	pageResponse := PageResponse{}
	pageResponse.Rows = &list
	pageResponse.Page = page.Page
	pageResponse.PageSize = page.Rows
	pageResponse.Records = records
	if page.Rows == 0 {
		page.Rows = 10
	}
	total := int(records) / page.Rows
	if records%int64(page.Rows) != 0 {
		total += 1
	}
	//log.Printf("records=%d,  一共%d页\n",records,  total)
	pageResponse.Total = total
	return &pageResponse
}

func GetErrorInfo(err error) string {
	if err == nil {
		return "success"
	} else {
		return err.Error()
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
