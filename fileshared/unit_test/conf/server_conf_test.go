package conf_test

import (
	"encoding/json"
	"fileshared/src/conf"
	"testing"
)

// TestInitPathSeparator 测试RootPath、ShareList path结尾符
func TestInitPathSeparator(t *testing.T) {
	tmpConf := conf.ServerConf{
		Port: 9087,
		Repo: conf.RepoConf{
			RootPath: "E:\\WindowsLib\\Desktop\\x",
			ShareList: []conf.ShareConf{
				{
					Path:      "/a",
					Rule:      "read",
					Recursive: false,
				},
				{
					Path:      "/b/",
					Rule:      "write",
					Recursive: true,
				},
				{
					Path:      "/c/",
					Rule:      "readwrite",
					Recursive: true,
				},
			},
		},
	}
	data, err := json.Marshal(tmpConf)
	if err != nil {
		t.Errorf("ServerConf serialized to json failed, error:%s", err.Error())
	}
	initResult := conf.GlobalConf.Init(data)
	if !initResult {
		t.Errorf("init conf, expected=true, actual=%t", initResult)
	}
	if tmpConf.Port != conf.GlobalConf.Port {
		t.Errorf("load port failed, expected=%d, actual=%d", tmpConf.Port, conf.GlobalConf.Port)
	}
	if len(conf.GlobalConf.Repo.RootPath) > 0 {
		if conf.GlobalConf.Repo.RootPath[len(conf.GlobalConf.Repo.RootPath)-1] == '/' {
			t.Errorf("RootPath can not end with '/', actual=%s", conf.GlobalConf.Repo.RootPath)
		}
	}
	for i := 0; i < len(conf.GlobalConf.Repo.ShareList); i++ {
		if tmpConf.Repo.ShareList[i].Rule != conf.GlobalConf.Repo.ShareList[i].Rule {
			t.Errorf("load share list rule failed, expected=%s, actual=%s",
				tmpConf.Repo.ShareList[i].Rule, conf.GlobalConf.Repo.ShareList[i].Rule)
		}
		if tmpConf.Repo.ShareList[i].Recursive != conf.GlobalConf.Repo.ShareList[i].Recursive {
			t.Errorf("load share list recursive failed, expected=%t, actual=%t",
				tmpConf.Repo.ShareList[i].Recursive, conf.GlobalConf.Repo.ShareList[i].Recursive)
		}
		rule := conf.GlobalConf.Repo.ShareList[i].Rule
		if rule != "read" && rule != "write" && rule != "readwrite" {
			t.Errorf("invalid rule, index=%d, expected=read or write or readwrite, actual=%s",
				i, rule)
		}
		path := conf.GlobalConf.Repo.ShareList[i].Path
		if len(path) == 0 {
			t.Errorf("share list path can not be empty, actual=%s", path)
		} else {
			if path[len(path)-1] == '/' {
				t.Errorf("share list path can not end with '/', actual=%s", path)
			}
		}
	}
}

// TestInitPathRootPathEmpty RootPath可以为空
func TestInitPathRootPathEmpty(t *testing.T) {
	tmpConf := conf.ServerConf{
		Port: 9087,
		Repo: conf.RepoConf{
			RootPath: "",
			ShareList: []conf.ShareConf{
				{
					Path:      "/a",
					Rule:      "read",
					Recursive: false,
				},
				{
					Path:      "/b",
					Rule:      "write",
					Recursive: true,
				},
				{
					Path:      "/c",
					Rule:      "readwrite",
					Recursive: true,
				},
			},
		},
	}
	data, err := json.Marshal(tmpConf)
	if err != nil {
		t.Errorf("ServerConf serialized to json failed, error:%s", err.Error())
	}
	initResult := conf.GlobalConf.Init(data)
	if !initResult {
		t.Errorf("init conf, expected=true, actual=%t", initResult)
	}
	if tmpConf.Port != conf.GlobalConf.Port {
		t.Errorf("load port failed, expected=%d, actual=%d", tmpConf.Port, conf.GlobalConf.Port)
	}
	if len(conf.GlobalConf.Repo.RootPath) > 0 {
		if conf.GlobalConf.Repo.RootPath[len(conf.GlobalConf.Repo.RootPath)-1] == '/' {
			t.Errorf("RootPath can not end with '/', actual=%s", conf.GlobalConf.Repo.RootPath)
		}
	}
	for i := 0; i < len(conf.GlobalConf.Repo.ShareList); i++ {
		if tmpConf.Repo.ShareList[i].Rule != conf.GlobalConf.Repo.ShareList[i].Rule {
			t.Errorf("load share list rule failed, expected=%s, actual=%s",
				tmpConf.Repo.ShareList[i].Rule, conf.GlobalConf.Repo.ShareList[i].Rule)
		}
		if tmpConf.Repo.ShareList[i].Recursive != conf.GlobalConf.Repo.ShareList[i].Recursive {
			t.Errorf("load share list recursive failed, expected=%t, actual=%t",
				tmpConf.Repo.ShareList[i].Recursive, conf.GlobalConf.Repo.ShareList[i].Recursive)
		}
		rule := conf.GlobalConf.Repo.ShareList[i].Rule
		if rule != "read" && rule != "write" && rule != "readwrite" {
			t.Errorf("invalid rule, index=%d, expected=read or write or readwrite, actual=%s",
				i, rule)
		}
		path := conf.GlobalConf.Repo.ShareList[i].Path
		if len(path) == 0 {
			t.Errorf("share list path can not be empty, actual=%s", path)
		} else {
			if path[len(path)-1] == '/' {
				t.Errorf("share list path can not end with '/', actual=%s", path)
			}
		}
	}
}

// TestInitPathSeparator4 RootPath可以为空，ShareList path不可为空
func TestInitRule(t *testing.T) {
	tmpConf := conf.ServerConf{
		Port: 9087,
		Repo: conf.RepoConf{
			RootPath: "",
			ShareList: []conf.ShareConf{
				{
					Path:      "/a",
					Rule:      "error",
					Recursive: false,
				},
				{
					Path:      "/b",
					Rule:      "write",
					Recursive: true,
				},
				{
					Path:      "/c",
					Rule:      "readwrite",
					Recursive: true,
				},
			},
		},
	}
	data, err := json.Marshal(tmpConf)
	if err != nil {
		t.Errorf("ServerConf serialized to json failed, error:%s", err.Error())
	}
	initResult := conf.GlobalConf.Init(data)
	if initResult {
		t.Errorf("init conf, expected=false, actual=%t", initResult)
	}
}

// TestInitError 测试初始化失败的情况
// ShareList path为空
// ShareList rule不是read、write、readwrite中的任何一个
func TestInitError(t *testing.T) {
	// 初始化失败：ShareList path不可为空
	tmpConf := conf.ServerConf{
		Port: 9087,
		Repo: conf.RepoConf{
			RootPath: "",
			ShareList: []conf.ShareConf{
				{
					Path:      "",
					Rule:      "read",
					Recursive: false,
				},
			},
		},
	}
	data, _ := json.Marshal(tmpConf)
	initResult := conf.GlobalConf.Init(data)
	if initResult {
		t.Errorf("init conf, expected=false, actual=%t", initResult)
	}

	// 初始化失败：ShareList rule只能是read、write或readwrite
	tmpConf = conf.ServerConf{
		Port: 9087,
		Repo: conf.RepoConf{
			RootPath: "",
			ShareList: []conf.ShareConf{
				{
					Path:      "/a",
					Rule:      "error",
					Recursive: false,
				},
			},
		},
	}
	data, _ = json.Marshal(tmpConf)
	initResult = conf.GlobalConf.Init(data)
	if initResult {
		t.Errorf("init conf, expected=false, actual=%t", initResult)
	}
}
