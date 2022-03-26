package utils

import (
    "encoding/json"
    "fmt"
    "log"
)

type LogInfo struct {
    Method string `json:"method"`
    Path string `json:"path"`
    ErrorCode int `json:"error_code"`
    ErrorInfo []string `json:"error"`
    MessageInfo []string `json:"message"`
    CostInfo map[string]int64 `json:"cost"`
}

func NewLogInfo() *LogInfo {
    return &LogInfo{
        Method:      "",
        Path:        "",
        ErrorCode:   0,
        ErrorInfo:   make([]string, 0),
        MessageInfo: make([]string, 0),
        CostInfo:    make(map[string]int64, 0),
    }
}

func (obj *LogInfo) LogError(format string, args ...interface{}) {
    obj.ErrorCode = -1
    obj.ErrorInfo = append(obj.ErrorInfo, fmt.Sprintf(format, args...))
}

func (obj *LogInfo) LogInfo(format string, args ...interface{}) {
    obj.MessageInfo = append(obj.MessageInfo, fmt.Sprintf(format, args...))
}

func (obj *LogInfo) WriteLog() {
    logLevel := ""
    if obj.ErrorCode == 0 {
        logLevel = "info"
    } else {
        logLevel = "error"
    }
    jsonStr, _ := json.Marshal(obj)
    log.Printf("|%s| %s", logLevel, jsonStr)
}
