package conf

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type ClientConf struct {
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

var IsRecursive = false
var IsDebugMode = false

// CreateNewConf 创建一个新的分享配置文件，并返回是否创建成功
func CreateNewConf() (ClientConf, bool) {
	if !checkCommandLine() {
		return ClientConf{}, false
	}
	if flag.NArg() == 1 && !strings.HasPrefix(flag.Arg(0), "http://") {
		return createConfFromFile()
	}
	return createConfFromCommandLine()
}

func checkCommandLine() bool {
	if len(os.Args) == 1 {
		helpStr := `获取一个远程文件并打印到标准输出
fs http://host/path/remote_file

获取一个远程文件，保存到本地，并指定本地文件名
fs http://host/path/remote_file /local_path/local_file

获取一个远程文件，保存到本地，指定本地目录名，本地文件名与远程文件名保持一致
fs http://host/path/remote_file /local_path/

列出远程目录下的内容
fs http://host/path/remote_path

以非递归的方式下载远程目录
fs http://host/path/remote_path /path/local_path

以递归的方式下载远程目录
fs http://host/path/remote_path /path/local_path -r

上传本地文件，必须指定远程文件名
fs /path/local_file http://host/path/remote_file

以非递归的方式上传本地目录
fs /path/local_path http://host/path/remote_path

以递归方式上传本地目录
fs /path/local_path http://host/path/remote_path -r

指定分享配置文件
fs /path/share_conf.json
`
		_, _ = fmt.Fprint(os.Stderr, helpStr)
		return false
	}
	flag.BoolVar(&IsRecursive, "r", false, "是否递归上传下载")
	flag.Parse()
	if flag.NArg() == 0 {
		_, _ = fmt.Fprint(os.Stderr, "invalid command line")
		return false
	}
	return true
}

func createConfFromFile() (ClientConf, bool) {
	clientConf := ClientConf{}
	data, err := ioutil.ReadFile(flag.Arg(0))
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "load conf file failed, error:%s", err.Error())
		return clientConf, false
	}
	err = json.Unmarshal(data, &clientConf)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "parse conf file failed, error:%s", err.Error())
		return clientConf, false
	}
	builder := PathBuilder{}
	for i := 0; i < len(clientConf.ShareConfList); i++ {
		clientConf.ShareConfList[i].Remote = builder.Clear().Append(clientConf.Host).
			Append(clientConf.ShareConfList[i].Remote).FullPath()
		clientConf.ShareConfList[i].Local = builder.Clear().Append(clientConf.RootDir).
			Append(clientConf.ShareConfList[i].Local).FullPath()
	}
	return clientConf, true
}

func createConfFromCommandLine() (ClientConf, bool) {
	if strings.HasPrefix(flag.Arg(0), "http://") {
		return createDownloadConf()
	}
	return createUploadConf()
}

func createDownloadConf() (ClientConf, bool) {
	localPath := ""
	if flag.NArg() == 2 {
		localPath = flag.Arg(1)
	}
	clientConf := ClientConf{
		Host:    "",
		RootDir: "",
		ShareConfList: []ShareConf{
			{
				Op:          "get",
				Remote:      flag.Arg(0),
				Local:       localPath,
				IsRecursive: IsRecursive,
			},
		},
	}
	return clientConf, true
}

func createUploadConf() (ClientConf, bool) {
	if flag.NArg() != 2 {
		return ClientConf{}, false
	}
	clientConf := ClientConf{
		Host:    "",
		RootDir: "",
		ShareConfList: []ShareConf{
			{
				Op:          "put",
				Remote:      flag.Arg(1),
				Local:       flag.Arg(0),
				IsRecursive: IsRecursive,
			},
		},
	}
	return clientConf, true
}

func (obj ShareConf) ToString() string {
	return fmt.Sprintf("op=%s,remote=%s,local=%s,recursive=%t", obj.Op, obj.Remote, obj.Local, obj.IsRecursive)
}
