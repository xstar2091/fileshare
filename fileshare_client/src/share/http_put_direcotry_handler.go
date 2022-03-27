package share

import (
	"fileshare_client/src/conf"
	"fileshare_client/src/utils"
	"io/ioutil"
	"os"
	"path/filepath"
)

type HttpPutDirectoryHandler struct {
	MiddlePath string
	Next       FileShare
}

func (obj HttpPutDirectoryHandler) Work(shareConf conf.ShareConf) bool {
	if shareConf.Op != "put" {
		utils.LogDebug("op is not put, HttpPutDirectoryHandler is ignored")
		return obj.Next.Work(shareConf)
	}
	var err error = nil
	defer func() {
		if err != nil {
			utils.LogError("failed, error:%s", err.Error())
		}
	}()

	fileInfo, err := os.Stat(shareConf.Local)
	if err != nil {
		return false
	}
	if !fileInfo.IsDir() {
		utils.LogDebug("local is not a directory, HttpPutDirectoryHandler is ignored")
		return obj.Next.Work(shareConf)
	}
	if shareConf.IsRecursive {
		utils.LogDebug("upload directory recursively, HttpPutDirectoryHandler is ignored")
		return obj.Next.Work(shareConf)
	}
	_, obj.MiddlePath = filepath.Split(shareConf.Local)
	fileInfoList, err := ioutil.ReadDir(shareConf.Local)
	if err != nil {
		return false
	}

	localBuilder := &conf.PathBuilder{}
	remoteBuilder := &conf.PathBuilder{}
	for _, item := range fileInfoList {
		if item.IsDir() {
			utils.LogDebug("ignore directory in HttpPutDirectoryHandler, dir:%s", item.Name())
			continue
		}
		localBuilder.Clear().Append(shareConf.Local).Append(item.Name())
		remoteBuilder.Clear().Append(shareConf.Remote).Append(obj.MiddlePath).Append(item.Name())
		utils.LogInfo("%s", localBuilder.FullPath())
		if !utils.UploadFile(localBuilder.FullPath(), remoteBuilder.FullPath()) {
			return false
		}
	}

	return true
}
