package repo

import (
	"errors"
	"fileshared/src/conf"
	"fileshared/src/utils"
	"fmt"
	"io/fs"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

type FileReader struct {
	LogInfo *utils.LogInfo
	Next    FileShare
}

func (obj *FileReader) Work(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		if obj.Next != nil {
			obj.Next.Work(w, r)
		}
		return
	}
	// 若以路径分隔符结尾，则认为是获取目录信息，本节点不做处理
	if strings.HasSuffix(r.URL.Path, "/") {
		if obj.Next != nil {
			obj.Next.Work(w, r)
		}
		return
	}

	fullPath := filepath.Join(conf.GlobalConf.Repo.RootPath, r.URL.Path)
	if !obj.checkReadable(&fullPath, w) {
		return
	}
	content, err := ioutil.ReadFile(fullPath)
	if err != nil {
		obj.LogInfo.LogError("read file failed, error:%s", err.Error())
		utils.ResponseError(http.StatusInternalServerError, obj.LogInfo, w)
		return
	}
	w.Header().Set("Content-Type", "application/text")
	w.Header().Set("Content-Length", fmt.Sprintf("%d", len(content)))
	w.WriteHeader(http.StatusOK)
	_, err = w.Write(content)
	if err != nil {
		obj.LogInfo.LogError("write response body failed:%s", err.Error())
	}
}

func (obj *FileReader) checkReadable(fullPath *string, w http.ResponseWriter) bool {
	fileInfo, err := os.Stat(*fullPath)
	if err != nil {
		if errors.Is(err, fs.ErrNotExist) {
			obj.LogInfo.LogError("file not found")
			utils.ResponseError(http.StatusNotFound, obj.LogInfo, w)
		} else if errors.Is(err, fs.ErrPermission) {
			obj.LogInfo.LogError("file no permission for read")
			utils.ResponseError(http.StatusForbidden, obj.LogInfo, w)
		} else {
			obj.LogInfo.LogError("stat file failed, error:%s", err.Error())
			utils.ResponseError(http.StatusForbidden, obj.LogInfo, w)
		}
		return false
	}
	if fileInfo.IsDir() {
		obj.LogInfo.LogError("request path is a directory")
		utils.ResponseError(http.StatusNotFound, obj.LogInfo, w)
		return false
	}
	return true
}
