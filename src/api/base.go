package api

import (
	"fmt"
	"net/http"
	"pool_backend/src/model/response"
	"pool_backend/src/util"
	"reflect"
	"time"

	"github.com/gin-gonic/gin"
)

var resp *response.Resp

// GetID 获取主键id
// @Summary 获取主键id
// @Description 获取主键id
// @Tags common
// @Accept application/json
// @Produce application/json
// @Success 200
// @Router /get_id [get]
func GetID(c *gin.Context) {
	for i := 0; i < 10; i++ {
		fmt.Println("id:", util.GenerateID(int64(i)).String())
	}
	id := util.GenerateID(1).String()

	ResponseHTTPOK("ok", "请求成功", id, c)
	return
}

// Health 健康检查
// @Summary 健康检查接口
// @Description 服务是否启动正常检查
// @Tags common
// @Accept application/json
// @Produce application/json
// @Param name query string false "用户名"
// @Success 200
// @Router /health [get]
func Health(c *gin.Context) {
	now := time.Now()
	id := util.GenerateID(1).String()
	data := map[string]interface{}{"Time": now, "id": id}
	ResponseHTTPOK("ok", "请求成功", data, c)
}

//Response api响应值
func Response(httpCode int, code string, msg string, data interface{}, c *gin.Context) {
	//反射判断interface是否为空值
	if reflect.TypeOf(data) != nil {
		c.JSON(httpCode, response.Resp{Code: code, Msg: msg, Data: data})
	} else {
		c.JSON(httpCode, response.Resp{Code: code, Msg: msg, Data: map[string]interface{}{}})
	}
}

//ResponseHTTPOK 成功响应
func ResponseHTTPOK(code string, msg string, data interface{}, c *gin.Context) {
	//反射判断interface是否为空值
	if reflect.TypeOf(data) != nil {
		c.JSON(http.StatusOK, response.Resp{Code: code, Msg: msg, Data: data})
	} else {
		c.JSON(http.StatusOK, response.Resp{Code: code, Msg: msg, Data: map[string]interface{}{}})
	}
}
