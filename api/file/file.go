package file

import (
	"encoding/base64"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/tigercandy/prado/global"
	"github.com/tigercandy/prado/internal/models/common/response"
	"io/ioutil"
	"os"
	"strings"
)

func UploadFile(c *gin.Context) {
	var (
		err      error
		prefix   string
		tag      string
		fileType string
		savePath string
		protocol string = "http"
		reqHost  string
	)

	tag, _ = c.GetPostForm("type")
	fileType = c.DefaultQuery("file_type", "images")

	if fileType != "images" && fileType != "files" {
		response.Error(c, -1, "仅支持图片和文件的上传!")
		return
	}

	if strings.HasPrefix(c.Request.Header.Get("Origin"), "https") {
		protocol = "https"
	}

	reqHostList := strings.Split(c.Request.Host, ":")
	if len(reqHostList) > 1 && reqHostList[1] == "80" {
		reqHost = reqHostList[0]
	} else {
		reqHost = c.Request.Host
	}

	if global.App.Config.Domain.GetHost {
		prefix = fmt.Sprintf("%s://%s/", protocol, reqHost)
	} else {
		prefix = fmt.Sprintf("%s://%s", prefix, global.App.Config.Domain.Url)
		if !strings.HasSuffix(global.App.Config.Domain.Url, "/") {
			prefix = protocol + "/"
		}
	}

	if tag == "" {
		tag = "1"
	}

	savePath = global.App.Config.File.SaveDir + "/" + fileType + "/"
	_, err = os.Stat(savePath)
	if os.IsNotExist(err) {
		err = os.MkdirAll(savePath, 0755)
		if err != nil {
			response.Error(c, -1, "创建图片上传目录失败!")
			return
		}
	}

	guid := strings.ReplaceAll(uuid.New().String(), "-", "")

	switch tag {
	case "1":
		files, err := c.FormFile("file")
		if err != nil {
			response.Error(c, -1, "图片不能为空")
			return
		}
		singleFile := savePath + guid + "-" + files.Filename
		_ = c.SaveUploadedFile(files, singleFile)
		response.Success(c, prefix+singleFile, "上传成功!")
		return
	case "2":
		files := c.Request.MultipartForm.File["file"]
		multiPartFile := make([]string, len(files))
		for _, file := range files {
			guid = strings.ReplaceAll(uuid.New().String(), "-", "")
			multiPartFileName := savePath + guid + "-" + file.Filename
			_ = c.SaveUploadedFile(file, multiPartFileName)
			multiPartFile = append(multiPartFile, prefix+multiPartFileName)
		}
		response.Success(c, multiPartFile, "上传成功!")
		return
	case "3":
		files, _ := c.GetPostForm("file")
		d, _ := base64.StdEncoding.DecodeString(files)
		_ = ioutil.WriteFile(savePath+guid+".jpg", d, 0666)
		response.Success(c, prefix+savePath+guid+".jpg", "上传成功!")
		return
	default:
		response.Error(c, 200, "上传文件标识不正确!")
		return
	}
}
