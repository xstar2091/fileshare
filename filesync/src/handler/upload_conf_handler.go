package handler

import (
    "encoding/json"
    "filesync/src/conf"
    "fmt"
    "io/fs"
    "io/ioutil"
    "path/filepath"
    "strings"
)

type UploadConfHandler struct {
    changedFileList []string
    middlePath string
    localRootPath string
    SyncConf *conf.Conf
    RecordConf *conf.RecordConf
    Next SyncHandler
}

func (obj *UploadConfHandler) Work() {
    obj.changedFileList = make([]string, 0)
    dir, file := filepath.Split(obj.SyncConf.Local.Path)
    obj.localRootPath = dir
    obj.middlePath = file
    obj.ReadChangeList(dir, file)
    if len(obj.changedFileList) == 0 {
        return
    }
    obj.CreateUploadConfFile()
    obj.Next.Work()
}

func (obj *UploadConfHandler) ReadChangeList(parentPath string, currentPath string) {
    fullPath := filepath.Join(parentPath, currentPath)
    fileInfoList, err := ioutil.ReadDir(fullPath)
    if err != nil {
        fmt.Println("walk dir failed:", fullPath)
        panic(err)
    }

    builder := &conf.PathBuilder{}
    for _, fileInfo := range fileInfoList {
        if obj.IsExclude(fileInfo.Name()) {
            continue
        }
        if fileInfo.IsDir() {
            obj.ReadChangeList(fullPath, fileInfo.Name())
            continue
        }
        filePath := builder.Clear().Append(fullPath).Append(fileInfo.Name()).FullPath()
        filePath = strings.ReplaceAll(filePath, "\\", "/")
        changeTime, ok := obj.RecordConf.Record[filePath]
        if !ok || changeTime != fileInfo.ModTime().UnixMilli() {
            obj.RecordConf.Record[filePath] = fileInfo.ModTime().UnixMilli()
            obj.changedFileList = append(obj.changedFileList, filePath)
        }
    }
}

func (obj *UploadConfHandler) IsExclude(name string) bool {
    ret := false
    for _, item := range obj.SyncConf.Local.Exclude {
        if name == item {
            ret = true
            break
        }
    }
    return ret
}

type UploadConf struct {
    Host          string      `json:"host"`
    RootDir       string      `json:"root_dir"`
    ShareConfList []ShareConf `json:"share"`
}

type ShareConf struct {
    Op          string `json:"op"`
    Remote      string `json:"remote"`
    Local       string `json:"local"`
    IsRecursive bool   `json:"recursive"`
}

func (obj *UploadConfHandler) CreateUploadConfFile() {
    uploadConf := UploadConf{
        Host:          obj.SyncConf.Remote.Host,
        RootDir:       obj.localRootPath,
        ShareConfList: make([]ShareConf, 0),
    }
    uploadConf.RootDir = strings.ReplaceAll(uploadConf.RootDir, "\\", "/")
    builder := &conf.PathBuilder{}
    for _, item := range obj.changedFileList {
        fileRelativePath := item[len(obj.localRootPath):]
        shareConf := ShareConf{
            Op:          "put",
            Remote:      fileRelativePath,
            Local:       builder.Clear().Append(fileRelativePath).FullPath(),
            IsRecursive: false,
        }
        uploadConf.ShareConfList = append(uploadConf.ShareConfList, shareConf)
    }

    data, err := json.Marshal(&uploadConf)
    if err != nil {
        fmt.Printf("create upload conf failed")
        panic(err)
    }
    err = ioutil.WriteFile(obj.SyncConf.UploadFile, data, fs.ModePerm)
    if err != nil {
        fmt.Printf("write upload conf failed")
        panic(err)
    }
}
