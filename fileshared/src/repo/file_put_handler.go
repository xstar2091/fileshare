package repo

import (
	"errors"
	"fileshared/src/conf"
	"fileshared/src/utils"
	"io/fs"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

type FileWriter struct {
	IsFileExist bool
	LogInfo     *utils.LogInfo
	Next        FileShare
}

func (obj *FileWriter) Work(w http.ResponseWriter, r *http.Request) {
	if r.Method != "PUT" {
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
	if !obj.checkWritable(&fullPath, w) {
		return
	}
	obj.beforeWrite(&fullPath)
	if obj.LogInfo.ErrorCode != 0 {
		utils.ResponseError(http.StatusInternalServerError, obj.LogInfo, w)
		return
	}
	content, err := ioutil.ReadAll(r.Body)
	if err != nil {
		obj.LogInfo.LogError("read body from request failed:%s", err.Error())
		utils.ResponseError(http.StatusInternalServerError, obj.LogInfo, w)
		return
	}
	err = ioutil.WriteFile(fullPath, content, 644)
	if err != nil {
		obj.LogInfo.LogError("write file failed, error:%s", err.Error())
		utils.ResponseError(http.StatusInternalServerError, obj.LogInfo, w)
		return
	}
	if obj.IsFileExist {
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusCreated)
	}
}

func (obj *FileWriter) checkWritable(fullPath *string, w http.ResponseWriter) bool {
	fileInfo, err := os.Stat(*fullPath)
	if err != nil {
		if errors.Is(err, fs.ErrNotExist) {
			// 文件不存在，可写
			obj.IsFileExist = false
			return true
		}
		obj.LogInfo.LogError("stat file failed, error:%s", err.Error())
		utils.ResponseError(http.StatusInternalServerError, obj.LogInfo, w)
		return false
	}

	if fileInfo.IsDir() {
		obj.LogInfo.LogError("path is a directory, need file only")
		utils.ResponseError(http.StatusBadRequest, obj.LogInfo, w)
		return false
	}

	obj.IsFileExist = true
	return true
}

// beforeWrite 写文件之前的准备工作，确保目录存在
func (obj *FileWriter) beforeWrite(fullPath *string) {
	createNewParentPath := false
	parentPath := filepath.Dir(*fullPath)
	defer func() {
		if createNewParentPath {
			obj.LogInfo.LogInfo("parent path not found, create new")
			obj.LogInfo.LogInfo("parent path:%s", parentPath)
			err := os.MkdirAll(parentPath, os.ModeDir)
			if err != nil {
				obj.LogInfo.LogError("create parent path failed:%s", err.Error())
			}
		}
	}()
	_, err := os.Stat(parentPath)
	if err != nil {
		createNewParentPath = true
		return
	}
}
