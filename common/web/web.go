package web

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
)

type Web struct {
	port   string
	engine *gin.Engine
}

func NewWeb(port string) *Web {
	engine := gin.Default()
	web := Web{
		port:   port,
		engine: engine,
	}
	return &web
}

// Run 运行
func (w *Web) Run() {
	_ = w.engine.Run(":" + w.port)
}

func (w *Web) GetEngine() *gin.Engine {
	return w.engine
}

// AddPostRequestHandler 添加请求处理器
func (w *Web) AddPostRequestHandler(path string, handler gin.HandlerFunc) {
	w.engine.POST(path, handler)
}

// AddGetRequestHandler 添加请求处理器
func (w *Web) AddGetRequestHandler(path string, handler gin.HandlerFunc) {
	w.engine.GET(path, handler)
}

// ParserRequest 解析请求数据
func (w *Web) ParserRequest(context *gin.Context) map[string]interface{} {
	var resJSONObj map[string]any
	data, err := io.ReadAll(context.Request.Body)
	if err != nil {
		fmt.Println(err)
		w.ReturnFail(context, "参数解析失败", nil)
		return nil
	}
	err = json.Unmarshal(data, &resJSONObj)
	if err != nil {
		fmt.Println(err)
		w.ReturnFail(context, "参数解析失败", nil)
		return nil
	}
	return resJSONObj
}

// ReturnSuccess 返回成功
func (w *Web) ReturnSuccess(context *gin.Context, msg string, data any) {
	var jsonObj = make(map[string]any)
	jsonObj["code"] = 1
	jsonObj["msg"] = msg
	if data != nil {
		jsonObj["data"] = data
	}
	context.AbortWithStatusJSON(http.StatusOK, jsonObj)
}

// ReturnFail 返回失败
func (w *Web) ReturnFail(context *gin.Context, msg string, data any) {
	var jsonObj = make(map[string]any)
	jsonObj["code"] = 0
	jsonObj["msg"] = msg
	if data != nil {
		jsonObj["data"] = data
	}
	context.AbortWithStatusJSON(http.StatusOK, jsonObj)
}

// ReturnCustom 返回失败
func (w *Web) ReturnCustom(context *gin.Context, code int, msg string, data any) {
	var jsonObj = make(map[string]any)
	jsonObj["code"] = code
	jsonObj["msg"] = msg
	if data != nil {
		jsonObj["data"] = data
	}
	context.AbortWithStatusJSON(http.StatusOK, jsonObj)
}
