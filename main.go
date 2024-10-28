package main

import (
	"github.com/gin-gonic/gin"
	"github.com/lyn0904/commongo/common/mysql"
	"github.com/lyn0904/commongo/common/redis"
	"github.com/lyn0904/commongo/common/sqlite"
	"github.com/lyn0904/commongo/common/web"
)

type User struct {
	id       int64
	userName string
	password string
	ty       int
	admin    bool
}

func main() {
	sqlite := sqlite.NewSqlite("./hell.db")
	sqlite.CreateTable("user", User{})
	sqlite.Db.Exec("delete from user where id=?", 1)
	sqlite.Db.Exec("insert into user (userName,password) values (?,?)", "嘎嘎", "321")
	redis.NewRedisClient("localhost:6379", "", 0)
	mysqlHelper := mysql.NewMysqlHelper("root", "123456", "localhost:3306", "gotest")
	mysqlHelper.CreateTable("user", User{})
	web := web.NewWeb("8088")
	web.AddPostRequestHandler("poet", func(context *gin.Context) {
		file, err := context.FormFile("file")
		if err != nil {
			web.ReturnFail(context, "上传失败", err.Error())
			return
		}
		context.SaveUploadedFile(file, "./"+file.Filename)
		web.ReturnSuccess(context, "成功", nil)
	})
	web.GetEngine().Use(func(context *gin.Context) {
		context.Next()
	})
	web.Run()

	//common.Blocking()
}
