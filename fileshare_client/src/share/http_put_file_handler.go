package share

import (
	"fileshare_client/src/conf"
	"fileshare_client/src/utils"
	"os"
)

type HttpPutFileHandler struct {
	Next FileShare
}

// Work 接口：fs /path/local_file http://host/path/remote_file
func (obj HttpPutFileHandler) Work(shareConf conf.ShareConf) bool {
	if shareConf.Op != "put" {
		utils.LogDebug("op is not put, HttpPutFileHandler is ignored")
		return obj.Next.Work(shareConf)
	}
	var err error = nil
	defer func() {
		if err != nil {
			utils.LogError("failed, error:%s", err.Error())
		}
	}()

	utils.LogInfo("%s", shareConf.Local)
	fileInfo, err := os.Stat(shareConf.Local)
	if err != nil {
		return false
	}
	if fileInfo.IsDir() {
		utils.LogDebug("local is a directory, HttpPutFileHandler is ignored")
		return obj.Next.Work(shareConf)
	}

	return utils.UploadFile(shareConf.Local, shareConf.Remote)
}
