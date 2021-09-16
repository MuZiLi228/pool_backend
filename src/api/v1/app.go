package v1

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
	"pool_backend/src/api"
	"pool_backend/src/global"
	"pool_backend/src/model/request"

	"github.com/gin-gonic/gin"
)

var (
	reqPutApp    request.PutApp
	filenamePrex string
)

//AppV1 api
type AppV1 struct {
}


// PutApp 更新app
// @Summary 更新app apk暂时存放服务器
// @Description 更新app json版本信息,apk暂时存放服务器
// @Tags app
// @Accept application/json
// @Produce application/json
// @Param type header string false "ios|android"
// @Param file formData  file true "file 上传文件"
// @Success 200 {object} response.Resp
// @Router /v1/app/uploads [post]
func (appV1 *AppV1) PutApp(c *gin.Context) {

	appType := c.PostForm("type")
	if appType == "" {
		api.ResponseHTTPOK("req_err", "错误请求!", "", c)
		return
	}

	//存放public目录下 app包、json文件信息
	cur, _ := os.Getwd()
	if appType == "android" {
		filenamePrex = filepath.Join(cur, "/src/public/app/android")
	} else if appType == "ios" {
		filenamePrex = filepath.Join(cur, "/src/public/app/ios")
	} else {
		api.ResponseHTTPOK("not_exist", "非法请求!", nil, c)
		return
	}

	//获取到所有的文件
	form, _ := c.MultipartForm()
	//获取到所有的文件数组
	files := form.File["files"]
	//遍历数组进行处理
	for _, file := range files {
		//进行文件保存
		err := c.SaveUploadedFile(file, filenamePrex+"/"+file.Filename)
		if err != nil {
			global.Logger.Error("上传app文件 报错:", err)
			api.ResponseHTTPOK("sys_err", "系统出错!", nil, c)
			return
		}
	}

	api.ResponseHTTPOK("ok", "请求成功!", nil, c)
	return
}

// GetAppVersion 获取app版本信息
// @Summary 获取app版本信息
// @Description 获取app版本信息
// @Tags app
// @Accept application/json
// @Produce application/json
// @Param type query string false "ios|android"
// @Success 200 {object} response.Resp
// @Router /v1/app [get]
func (appV1 *AppV1) GetAppVersion(c *gin.Context) {

	appType, ok := c.GetQuery("type")
	if !ok {
		api.ResponseHTTPOK("req_err", "错误请求!", "", c)
	}
	cur, _ := os.Getwd()
	if appType == "android" {
		filenamePrex = filepath.Join(cur, "/src/public/app/android")
	} else if appType == "ios" {
		filenamePrex = filepath.Join(cur, "/src/public/app/ios")
	} else {
		api.ResponseHTTPOK("not_exist", "非法请求!", nil, c)
		return
	}
	// 打开json文件
	jsonFile, err := os.Open(filenamePrex + "/" + "app.json")
	// 最好要处理以下错误
	if err != nil {
		global.Logger.Error("打开json文件 报错:", err.Error())
		api.ResponseHTTPOK("sys_err", "系统出错!", nil, c)
		return
	}

	// 要记得关闭
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var data map[string]interface{}
	if err := json.Unmarshal(byteValue, &data); err != nil {
		global.Logger.Error("解析 app版本json 报错:", err.Error())
		api.ResponseHTTPOK("fail", "请求失败!", "", c)
		return
	}
	api.ResponseHTTPOK("ok", "请求成功!", data, c)
	return
}

// AppDownload app下载
// @Summary app下载链接
// @Description app下载
// @Tags app
// @Accept application/json
// @Produce application/json
// @Param type path string false "ios|android"
// @Param name path string false "apk name"
// @Success 200 {object} response.Resp
// @Router /v1/app/down/{type}/{name} [get]
func (appV1 *AppV1) AppDownload(c *gin.Context) {
	appType := c.Param("type")
	appName := c.Param("name")

	cur, _ := os.Getwd()
	if appType == "android" {
		filenamePrex = filepath.Join(cur, "/src/public/app/android")
	} else if appType == "ios" {
		filenamePrex = filepath.Join(cur, "/src/public/app/ios")
	} else {
		api.ResponseHTTPOK("not_exist", "非法请求!", nil, c)
		return
	}
	//打开文件
	filePath := filenamePrex + "/" + appName
	fileTmp, errByOpenFile := os.Open(filePath)
	defer fileTmp.Close()

	//获取文件的名称
	c.Header("Content-Type", "application/octet-stream")
	c.Header("Content-Disposition", "attachment; filename="+appName)
	c.Header("Content-Transfer-Encoding", "binary")
	c.Header("Cache-Control", "no-cache")
	if filePath == "" || appName == "" || errByOpenFile != nil {
		api.Response(404, "not_exist", "文件不存在!", nil, c)
		return
	}
	// c.Header("Content-Type", "application/octet-stream")
	// c.Header("Content-Disposition", "attachment; filename="+appName)
	// c.Header("Content-Transfer-Encoding", "binary")

	c.File(filePath)
	return

}
