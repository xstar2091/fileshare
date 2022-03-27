package handler

import (
    "encoding/json"
    "filesync/src/conf"
    "fmt"
    "io/fs"
    "io/ioutil"
)

type RecordUpdateHandler struct {
    SyncConf *conf.Conf
    RecordConf *conf.RecordConf
    Next SyncHandler
}

func (obj *RecordUpdateHandler) Work() {
    data, err := json.Marshal(obj.RecordConf.Record)
    if err != nil {
        fmt.Println("record map parsed to json failed")
        panic(err)
    }
    err = ioutil.WriteFile(obj.SyncConf.RecordFile, data, fs.ModePerm)
    if err != nil {
        fmt.Println("update record failed")
        panic(err)
    }
}
