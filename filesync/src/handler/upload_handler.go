package handler

import (
    "bufio"
    "errors"
    "filesync/src/conf"
    "fmt"
    "io"
    "os/exec"
    "strings"
)

type UploadHandler struct {
    SyncConf *conf.Conf
    RecordConf *conf.RecordConf
    Next SyncHandler
}

func (obj *UploadHandler) Work() {
    cmd := exec.Command("fsc", obj.SyncConf.UploadFile)
    stdout, err := cmd.StdoutPipe()
    if err != nil {
        fmt.Println("create upload stdout pipe failed")
        panic(err)
    }
    stderr, err := cmd.StderrPipe()
    if err != nil {
        fmt.Println("create upload stderr pipe failed")
        panic(err)
    }
    err = cmd.Start()
    if err != nil {
        fmt.Println("upload failed")
        panic(err)
    }

    needExit := false
    buf := bufio.NewReader(stdout)
    for {
        line, err := buf.ReadString('\n')
        line = strings.TrimSpace(line)
        fmt.Println(line)
        if err != nil {
            if errors.Is(err, io.EOF) {
                break
            } else {
                needExit = true
                fmt.Println("upload failed")
                panic(err)
            }
        }
    }
    buf = bufio.NewReader(stderr)
    for {
        line, err := buf.ReadString('\n')
        line = strings.TrimSpace(line)
        if len(line) > 0 {
            fmt.Println(line)
            needExit = true
        }
        if err != nil {
            if errors.Is(err, io.EOF) {
                break
            } else {
                needExit = true
                fmt.Println("print upload error failed")
                panic(err)
            }
        }
    }
    if needExit {
        return
    }
    obj.Next.Work()
}
