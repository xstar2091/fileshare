package repo

import (
	"fileshared/src/conf"
	"os"
	"testing"
)

// 单测上传文件接口的文件名前缀
var UnitTestTempFileNamePrefix = "unittest_temp_file"

func setup() {
	conf.GlobalConf = &conf.ServerConf{
		Port: 8097,
		Repo: conf.RepoConf{
			RootPath: ".",
			ShareList: []conf.ShareConf{
				{
					Path:      "/read",
					Rule:      "read",
					Recursive: true,
				},
				{
					Path:      "/write",
					Rule:      "write",
					Recursive: true,
				},
				{
					Path:      "/readwrite",
					Rule:      "readwrite",
					Recursive: true,
				},
				{
					Path:      "/read_norecursive",
					Rule:      "read",
					Recursive: false,
				},
				{
					Path:      "/write_norecursive",
					Rule:      "write",
					Recursive: false,
				},
				{
					Path:      "/readwrite_norecursive",
					Rule:      "readwrite",
					Recursive: false,
				},
			},
		},
	}
}

func TestMain(m *testing.M) {
	setup()
	code := m.Run()
	os.Exit(code)
}
