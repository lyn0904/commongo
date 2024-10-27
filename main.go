package main

import (
	"github.com/gin-gonic/gin"
	"github.com/lyn0904/commongo/common/mysql"
	"github.com/lyn0904/commongo/common/redis"
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
	redis.NewRedisClient("localhost:6379", "", 0)
	mysqlHelper := mysql.NewMysqlHelper("root", "123456", "localhost:3306", "gotest")
	mysqlHelper.CreateTable("user", User{})
	webGin := web.NewWebGin("8088")
	webGin.AddPostRequestHandler("poet", func(context *gin.Context) {
		file, err := context.FormFile("file")
		if err != nil {
			web.ReturnFail(context, "上传失败", err.Error())
			return
		}
		context.SaveUploadedFile(file, "./"+file.Filename)
		web.ReturnSuccess(context, "成功", nil)
	})
	webGin.GetEngine().Use(func(context *gin.Context) {
		context.Next()
	})
	webGin.Run()

	//common.Blocking()
}
