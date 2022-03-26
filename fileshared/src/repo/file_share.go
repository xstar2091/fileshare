package repo

import (
	"fileshared/src/utils"
	"net/http"
)

type FileShare interface {
	Work(w http.ResponseWriter, r *http.Request)
}

func CreateShare() (FileShare, *utils.LogInfo) {
	logInfo := utils.NewLogInfo()
	unsupported := &UnsupportedHandler{
		LogInfo: logInfo,
		Next:    nil,
	}
	writer := &FileWriter{
		IsFileExist: false,
		LogInfo:     logInfo,
		Next:        unsupported,
	}
	reader := &FileReader{
		LogInfo: logInfo,
		Next:    writer,
	}
	dirReader := &DirectoryReader{
		LogInfo: logInfo,
		Next:    reader,
	}
	validator := &Validator{
		LogInfo: logInfo,
		Next:    dirReader,
	}
	return validator, logInfo
}
