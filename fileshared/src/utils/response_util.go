package utils

import "net/http"

func ResponseError(httpStatusCode int, logInfo *LogInfo, w http.ResponseWriter) {
    logInfo.LogError("response %d", httpStatusCode)
    w.WriteHeader(httpStatusCode)
}

func ResponseData(content []byte, contentType string, w http.ResponseWriter, logInfo *LogInfo) {
    w.Header().Set("Content-Type", contentType)
    w.WriteHeader(http.StatusOK)
    _, err := w.Write(content)
    if err != nil {
        logInfo.LogError("write response body failed:%s", err.Error())
        ResponseError(http.StatusInternalServerError, logInfo, w)
    }
}
