package handler

import (
    "fileshared/src/repo"
    "fileshared/src/utils"
    "net/http"
)

type ServerHandler struct {

}

func (obj *ServerHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    share, logInfo := repo.CreateShare()
    InitLog(r, logInfo)
    defer logInfo.WriteLog()
    share.Work(w, r)
}

func InitLog(r *http.Request, logInfo *utils.LogInfo) {
    logInfo.Method = r.Method
    logInfo.Path = r.URL.Path
}
