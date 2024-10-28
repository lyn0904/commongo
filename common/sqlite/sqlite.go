package sqlite

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"reflect"
)

type Sqlite struct {
	Db *sql.DB
}

func NewSqlite(dataSourceName string) Sqlite {
	sqlite := Sqlite{}
	open, err := sql.Open("sqlite3", dataSourceName)
	if err != nil {
		log.Fatal(err)
	}
	sqlite.Db = open
	return sqlite
}

func (s Sqlite) CreateTable(tableName string, tableStruct any) {
	f := reflect.TypeOf(tableStruct)
	var sql = "create table if not exists `" + tableName + "`(id integer not null primary key autoincrement,"
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
				sql = sql + name + " text)"
			} else {
				sql = sql + name + " text,"
			}
			continue
		}
		if typeName == "int" {
			if i == size-1 {
				sql = sql + name + " integer)"
			} else {
				sql = sql + name + " integer,"
			}
			continue
		}
		if typeName == "bool" {
			if i == size-1 {
				sql = sql + name + " integer)"
			} else {
				sql = sql + name + " integer,"
			}
			continue
		}
	}
	log.Println("sql:", sql)
	_, err := s.Db.Exec(sql)
	if err != nil {
		log.Println(err)
	}
}
