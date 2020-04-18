package infrastructure

import (
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/tanimutomo/clean-architecture-api-go/interfaces/database"
)

type DBHandler struct {
	Conn *gorm.DB
}

func NewDBHandler() database.DBHandler {
	DBMS := os.Getenv("CAAG_DBMS")
	USER := os.Getenv("CAAG_USER")
	PASS := os.Getenv("CAAG_PASS")
	DBNAME := os.Getenv("CAAG_DBNAME")
	CONNECT := USER + ":" + PASS + "@/" + DBNAME + "?charset=utf8&parseTime=true&loc=Local"
	conn, err := gorm.Open(DBMS, CONNECT)
	if err != nil {
		// TODO
	}
	dbHandler := new(DBHandler)
	dbHandler.Conn = conn
	return dbHandler
}

func (handler *DBHandler) First(out interface{}, where ...interface{}) *gorm.DB {
	return handler.Conn.First(out, where...)
}

func (handler *DBHandler) Find(out interface{}, where ...interface{}) *gorm.DB {
	return handler.Conn.Find(out, where...)
}

func (handler *DBHandler) Exec(query string, values ...interface{}) *gorm.DB {
	return handler.Conn.Exec(query, values...)
}

func (handler *DBHandler) Raw(query string, values ...interface{}) *gorm.DB {
	return handler.Conn.Raw(query, values)
}

func (handler *DBHandler) Create(values interface{}) *gorm.DB {
	return handler.Conn.Create(values)
}

func (handler *DBHandler) Save(value interface{}) *gorm.DB {
	return handler.Conn.Delete(value)
}

func (handler *DBHandler) Delete(value interface{}) *gorm.DB {
	return handler.Conn.Delete(value)
}

func (handler *DBHandler) Where(query interface{}, args ...interface{}) *gorm.DB {
	return handler.Conn.Where(query, args...)
}
