package repo

import (
	"encoding/json"
	"fileshared/src/conf"
	"fileshared/src/utils"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

type DirectoryReader struct {
	LogInfo *utils.LogInfo
	Next    FileShare
}

func (obj *DirectoryReader) Work(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		if obj.Next != nil {
			obj.Next.Work(w, r)
		}
		return
	}
	fullPath := filepath.Join(conf.GlobalConf.Repo.RootPath, r.URL.Path)
	fileInfo, err := os.Stat(fullPath)
	if err != nil {
		obj.LogInfo.LogError("path not found:%s", err.Error())
		utils.ResponseError(http.StatusNotFound, obj.LogInfo, w)
		return
	}
	if !fileInfo.IsDir() {
		// 若以路径分隔符结尾，则认为是获取目录信息，本节点直接返回错误信息
		if strings.HasSuffix(fullPath, "/") {
			obj.LogInfo.LogError("path not found")
			utils.ResponseError(http.StatusNotFound, obj.LogInfo, w)
			return
		}
		if obj.Next != nil {
			obj.Next.Work(w, r)
		}
		return
	}
	fileInfoList, err := ioutil.ReadDir(fullPath)
	if err != nil {
		obj.LogInfo.LogError("read directory failed:%s", err.Error())
		utils.ResponseError(http.StatusInternalServerError, obj.LogInfo, w)
		return
	}

	response := DirectoryGetResponse{
		Path: r.URL.Path,
		Info: make([]PathInfo, 0),
	}
	for _, item := range fileInfoList {
		info := PathInfo{
			IsDir:      item.IsDir(),
			Mode:       uint32(item.Mode()),
			ModifyTime: item.ModTime().Format("2006-01-02 15:04:05"),
			Name:       item.Name(),
			Size:       item.Size(),
		}
		response.Info = append(response.Info, info)
	}
	resp, err := json.Marshal(response)
	if err != nil {
		obj.LogInfo.LogError("directory info serialized to json failed:%s", err.Error())
		utils.ResponseError(http.StatusInternalServerError, obj.LogInfo, w)
		return
	}
	utils.ResponseData(resp, "application/json", w, obj.LogInfo)
}

type PathInfo struct {
	IsDir      bool   `json:"is_dir"`
	Mode       uint32 `json:"mode"`
	ModifyTime string `json:"modify_time"`
	Name       string `json:"name"`
	Size       int64  `json:"size"`
}

type DirectoryGetResponse struct {
	Path string     `json:"path"`
	Info []PathInfo `json:"info"`
}
