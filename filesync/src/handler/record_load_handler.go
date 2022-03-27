package handler

import (
    "encoding/json"
    "errors"
    "filesync/src/conf"
    "fmt"
    "io/fs"
    "io/ioutil"
)

type RecordLoadHandler struct {
    SyncConf *conf.Conf
    RecordConf *conf.RecordConf
    Next SyncHandler
}

func (obj *RecordLoadHandler) Work() {
    fmt.Printf("load record from %s\n", obj.SyncConf.RecordFile)
    data, err := ioutil.ReadFile(obj.SyncConf.RecordFile)
    if err == nil {
        obj.RecordConf.Record = make(map[string]int64, 0)
        err = json.Unmarshal(data, &obj.RecordConf.Record)
        if err != nil {
            fmt.Println("parse record to json failed")
            panic(err)
        }
    } else if errors.Is(err, fs.ErrNotExist) {
        obj.RecordConf.Record = make(map[string]int64, 0)
    } else {
        fmt.Printf("load record failed")
        panic(err)
    }

    obj.Next.Work()
}
