package common

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
)

type WebGin struct {
	port   string
	engine *gin.Engine
}

func NewWebGin(port string) *WebGin {
	engine := gin.Default()
	webGin := WebGin{
		port:   port,
		engine: engine,
	}
	return &webGin
}

// Run 运行
func (w *WebGin) Run() {
	_ = w.engine.Run(":" + w.port)
}

func (w *WebGin) GetEngine() *gin.Engine {
	return w.engine
}

// AddPostRequestHandler 添加请求处理器
func (w *WebGin) AddPostRequestHandler(path string, handler gin.HandlerFunc) {
	w.engine.POST(path, handler)
}

// AddGetRequestHandler 添加请求处理器
func (w *WebGin) AddGetRequestHandler(path string, handler gin.HandlerFunc) {
	w.engine.GET(path, handler)
}

// ParserRequest 解析请求数据
func ParserRequest(context *gin.Context) map[string]interface{} {
	var resJSONObj map[string]any
	data, err := io.ReadAll(context.Request.Body)
	if err != nil {
		fmt.Println(err)
		ReturnFail(context, "参数解析失败", nil)
		return nil
	}
	err = json.Unmarshal(data, &resJSONObj)
	if err != nil {
		fmt.Println(err)
		ReturnFail(context, "参数解析失败", nil)
		return nil
	}
	return resJSONObj
}

// ReturnSuccess 返回成功
func ReturnSuccess(context *gin.Context, msg string, data any) {
	var jsonObj = make(map[string]any)
	jsonObj["code"] = 1
	jsonObj["msg"] = msg
	if data != nil {
		jsonObj["data"] = data
	}
	context.AbortWithStatusJSON(http.StatusOK, jsonObj)
}

// ReturnFail 返回失败
func ReturnFail(context *gin.Context, msg string, data any) {
	var jsonObj = make(map[string]any)
	jsonObj["code"] = 0
	jsonObj["msg"] = msg
	if data != nil {
		jsonObj["data"] = data
	}
	context.AbortWithStatusJSON(http.StatusOK, jsonObj)
}

// ReturnCustom 返回失败
func ReturnCustom(context *gin.Context, code int, msg string, data any) {
	var jsonObj = make(map[string]any)
	jsonObj["code"] = code
	jsonObj["msg"] = msg
	if data != nil {
		jsonObj["data"] = data
	}
	context.AbortWithStatusJSON(http.StatusOK, jsonObj)
}
