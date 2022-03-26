package repo

import (
	"fileshared/src/conf"
	"fileshared/src/utils"
	"net/http"
	"strings"
)

type Validator struct {
	LogInfo *utils.LogInfo
	Next    FileShare
}

func (obj *Validator) Work(w http.ResponseWriter, r *http.Request) {
	if !obj.checkMethod(&r.Method, w) {
		return
	}
	if !obj.checkUrlPath(&r.URL.Path, w) {
		return
	}
	if !obj.checkPermission(w, r) {
		return
	}
	if obj.Next != nil {
		obj.Next.Work(w, r)
	}
}

func (obj *Validator) checkMethod(method *string, w http.ResponseWriter) bool {
	if *method == "GET" || *method == "PUT" {
		return true
	}
	obj.LogInfo.LogError("not allowed method:%s", *method)
	utils.ResponseError(http.StatusMethodNotAllowed, obj.LogInfo, w)
	return false
}

// 路径中不允许包含..
func (obj *Validator) checkUrlPath(urlPath *string, w http.ResponseWriter) bool {
	if strings.Contains(*urlPath, "..") {
		obj.LogInfo.LogError("invalid path, contained \"..\"")
		utils.ResponseError(http.StatusBadRequest, obj.LogInfo, w)
		return false
	}
	return true
}

func (obj *Validator) checkPermission(w http.ResponseWriter, r *http.Request) bool {
	urlPath := &r.URL.Path
	sharedConf := obj.getSharedConf(urlPath)

	if !obj.checkPath(sharedConf, w) {
		return false
	}
	if !obj.checkRecursive(urlPath, sharedConf, w) {
		return false
	}
	if !obj.checkRule(&r.Method, sharedConf, w) {
		return false
	}

	return true
}

func (obj *Validator) getSharedConf(urlPath *string) *conf.ShareConf {
	pathLen := 0
	var shareConf *conf.ShareConf = nil
	for _, item := range conf.GlobalConf.Repo.ShareList {
		// url path 必须包含ShareList中的Path前缀
		if strings.HasPrefix(*urlPath, item.Path) {
			if len(item.Path) > pathLen {
				pathLen = len(item.Path)
				shareConf = &conf.ShareConf{
					Path:      item.Path,
					Rule:      item.Rule,
					Recursive: item.Recursive,
				}
			}
		}
	}
	return shareConf
}

// 配置文件中没有请求的路径
func (obj *Validator) checkPath(shareConf *conf.ShareConf, w http.ResponseWriter) bool {
	if shareConf == nil {
		obj.LogInfo.LogError("path not in the conf")
		utils.ResponseError(http.StatusForbidden, obj.LogInfo, w)
		return false
	}
	return true
}

// 不允许递归
func (obj *Validator) checkRecursive(urlPath *string, shareConf *conf.ShareConf, w http.ResponseWriter) bool {
	fileName := (*urlPath)[len(shareConf.Path):]
	if fileName == "/" {
		fileName = ""
	}
	if strings.Contains(fileName, "/") {
		if !shareConf.Recursive {
			obj.LogInfo.LogError("path found in the conf, but not allowed recursive")
			obj.LogInfo.LogError("conf path: {}, recursive:%d", shareConf.Path, shareConf.Recursive)
			utils.ResponseError(http.StatusForbidden, obj.LogInfo, w)
			return false
		}
	}
	return true
}

// 校验读写权限
func (obj *Validator) checkRule(method *string, shareConf *conf.ShareConf, w http.ResponseWriter) bool {
	if *method == "GET" {
		if shareConf.Rule == "read" || shareConf.Rule == "readwrite" {
			return true
		} else {
			obj.LogInfo.LogError("not readable path")
			obj.LogInfo.LogError("path:%s, rule:%s", shareConf.Path, shareConf.Rule)
			utils.ResponseError(http.StatusForbidden, obj.LogInfo, w)
		}
	}
	if *method == "PUT" {
		if shareConf.Rule == "write" || shareConf.Rule == "readwrite" {
			return true
		} else {
			obj.LogInfo.LogError("not writable path")
			obj.LogInfo.LogError("path:%s, rule:%s", shareConf.Path, shareConf.Rule)
			utils.ResponseError(http.StatusForbidden, obj.LogInfo, w)
		}
	}
	return false
}
