package share

import (
	"fileshare_client/src/conf"
	"fileshare_client/src/utils"
	"os"
	"path/filepath"
)

type HttpPutDirectoryRecursiveHandler struct {
	MiddlePath string
	Next       FileShare
}

func (obj HttpPutDirectoryRecursiveHandler) Work(shareConf conf.ShareConf) bool {
	if shareConf.Op != "put" {
		utils.LogDebug("op is not put, HttpPutDirectoryRecursiveHandler is ignored")
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
		utils.LogDebug("local is not a directory, HttpPutDirectoryRecursiveHandler is ignored")
		return obj.Next.Work(shareConf)
	}
	if !shareConf.IsRecursive {
		utils.LogDebug("upload directory is not set recursively, HttpPutDirectoryRecursiveHandler is ignored")
		return obj.Next.Work(shareConf)
	}
	_, obj.MiddlePath = filepath.Split(shareConf.Local)
	return false
}
