package mysql

import (
	_ "database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"log"
	"reflect"
)

type Mysql struct {
	db *sqlx.DB
}

func NewMysqlHelper(userName, password, host, dbName string) Mysql {
	url := userName + ":" + password + "@tcp(" + host + ")/" + dbName + "?charset=utf8mb4&parseTime=True"
	db, err := sqlx.Connect("mysql", url)
	if err != nil {
		log.Fatal(err)
	}
	db.SetMaxOpenConns(20) // 设置数据库连接池的最大连接数
	db.SetMaxIdleConns(10) // 设置数据库连接池的最大空闲连接数
	mysqlHelper := Mysql{
		db: db,
	}
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("连接数据库成功")
	}
	return mysqlHelper
}

// CreateTable 创建表
func (h *Mysql) CreateTable(tableName string, tableStruct any) {
	f := reflect.TypeOf(tableStruct)
	var sql = "create table if not exists `" + tableName + "`(id bigint primary key auto_increment,"
	size := f.NumField()
	for i := 0; i < size; i++ {
		field := f.Field(i)
		name := field.Name
		typeName := field.Type.Name()
		if name == "id" {
			continue
		}
		if typeName == "string" {
			if i == size-1 {
				sql = sql + name + " varchar(255));"
			} else {
				sql = sql + name + " varchar(255),"
			}
			continue
		}
		if typeName == "int" {
			if i == size-1 {
				sql = sql + name + " int default 0);"
			} else {
				sql = sql + name + " int default 0,"
			}
			continue
		}
		if typeName == "bool" {
			if i == size-1 {
				sql = sql + name + " int default 0);"
			} else {
				sql = sql + name + " int default 0,"
			}
			continue
		}
	}
	_, err := h.db.Exec(sql)
	if err != nil {
		log.Println(err)
	}
}
