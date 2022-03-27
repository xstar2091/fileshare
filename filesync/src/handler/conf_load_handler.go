package handler

import (
    "encoding/json"
    "filesync/src/conf"
    "fmt"
    "io/ioutil"
)

type ConfLoadHandler struct {
    SyncConf *conf.Conf
    RecordConf *conf.RecordConf
    Next SyncHandler
}

func (obj *ConfLoadHandler) Work() {
    builder := &conf.PathBuilder{}
    obj.SyncConf.ConfFile = builder.Clear().Append(conf.FileSyncDir).Append(conf.SyncConfFileName).FullPath()
    obj.SyncConf.RecordFile = builder.Clear().Append(conf.FileSyncDir).Append(conf.RecordFileName).FullPath()
    obj.SyncConf.UploadFile = builder.Clear().Append(conf.FileSyncDir).Append(conf.UploadConfFileName).FullPath()

    fmt.Printf("load conf from %s\n", obj.SyncConf.ConfFile)
    data, err := ioutil.ReadFile(obj.SyncConf.ConfFile)
    if err != nil {
        panic(err)
    }

    err = json.Unmarshal(data, obj.SyncConf)
    if err != nil {
        fmt.Println("parse conf file to json failed")
        panic(err)
    }

    obj.Next.Work()
}
