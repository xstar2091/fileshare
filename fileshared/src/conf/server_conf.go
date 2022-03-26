package conf

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

var GlobalConf = &ServerConf{}

type ServerConf struct {
	Port uint16   `json:"port"`
	Repo RepoConf `json:"repo"`
}

type RepoConf struct {
	// 分享根目录，最后一个字符不能是路径分隔符
	RootPath  string      `json:"root"`
	ShareList []ShareConf `json:"share"`
}

type ShareConf struct {
	// 分享目录，RootPath的相对目录，不可为空，必须以路径分隔符开始，最后一个字符不能是路径分隔符
	Path string `json:"path"`
	// 权限，只能是：read、write、readwrite中的任一个
	Rule string `json:"rule"`
	// 分享目录后，其子目录是否继承分享规则
	Recursive bool `json:"recursive"`
}

func InitConf() bool {
	flag.Parse()
	var confFile = ""
	if flag.NArg() == 1 {
		confFile = flag.Arg(0)
	} else {
		confFile = "./conf.json"
	}
	data, err := ioutil.ReadFile(confFile)
	if err != nil {
		log.Printf("read conf file failed, error:%s", err.Error())
		return false
	}
	return GlobalConf.Init(data)
}

func (obj *ServerConf) Init(data []byte) bool {
	err := json.Unmarshal(data, obj)
	if err != nil {
		log.Printf("parse conf faile failed, error:%s", err.Error())
		return false
	}
	obj.Repo.RootPath = obj.correctPath(obj.Repo.RootPath)
	for i := 0; i < len(obj.Repo.ShareList); i++ {
		obj.Repo.ShareList[i].Path = obj.correctPath(obj.Repo.ShareList[i].Path)
	}
	return obj.isValid()
}

func (obj *ServerConf) correctPath(path string) string {
	pathLen := len(path)
	if pathLen == 0 {
		return path
	}
	path = strings.ReplaceAll(path, "\\", "/")
	// root path 最后一个字符不能是'/'
	lastIndex := pathLen - 1
	if path[lastIndex] == '/' {
		path = path[0:lastIndex]
	}
	return path
}

func (obj *ServerConf) isValid() bool {
	index := -1
	ret := true
	for _, item := range obj.Repo.ShareList {
		index += 1
		if item.Path == "" {
			ret = false
			fmt.Printf("share path can not be empty or /, index=%d\n", index)
			break
		}
		if item.Rule != "read" && item.Rule != "write" && item.Rule != "readwrite" {
			ret = false
			fmt.Printf("share rule must be read, write or readwrite, index=%d\n", index)
			break
		}
	}
	return ret
}
