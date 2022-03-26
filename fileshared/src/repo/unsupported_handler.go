package repo

import (
	"fileshared/src/utils"
	"net/http"
)

type UnsupportedHandler struct {
	LogInfo *utils.LogInfo
	Next    FileShare
}

func (obj *UnsupportedHandler) Work(w http.ResponseWriter, _ *http.Request) {
	obj.LogInfo.LogError("no supported operation node found")
	utils.ResponseError(http.StatusBadRequest, obj.LogInfo, w)
}
