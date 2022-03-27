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
        fmt.Println("create upload pipe failed")
        panic(err)
    }
    err = cmd.Start()
    if err != nil {
        fmt.Println("upload failed")
        panic(err)
    }

    buf := bufio.NewReader(stdout)
    for {
        line, err := buf.ReadString('\n')
        line = strings.TrimSpace(line)
        fmt.Println(line)
        if err != nil {
            if errors.Is(err, io.EOF) {
                break
            } else {
                fmt.Println("upload failed")
                panic(err)
            }
        }
    }
    obj.Next.Work()
}
