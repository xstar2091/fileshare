package handler

import "filesync/src/conf"

type SyncHandler interface {
    Work()
}

func CreateSyncHandler() SyncHandler {
    syncConf := &conf.Conf{}
    recordConf := &conf.RecordConf{}
    recordUpdate := &RecordUpdateHandler{
        SyncConf:   syncConf,
        RecordConf: recordConf,
        Next:       nil,
    }
    upload := &UploadHandler{
        SyncConf:   syncConf,
        RecordConf: recordConf,
        Next:       recordUpdate,
    }
    uploadConf := &UploadConfHandler{
        SyncConf:   syncConf,
        RecordConf: recordConf,
        Next:       upload,
    }
    recordLoad := &RecordLoadHandler{
        SyncConf:   syncConf,
        RecordConf: recordConf,
        Next:       uploadConf,
    }
    confLoad := &ConfLoadHandler{
        SyncConf:   syncConf,
        RecordConf: recordConf,
        Next:       recordLoad,
    }
    return confLoad
}
