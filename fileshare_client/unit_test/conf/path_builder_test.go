package conf

import (
	"fileshare_client/src/conf"
	"testing"
)

// TestPathBuilderAppendUnix 测试Unix风格路径的构造
func TestPathBuilderAppendUnix(t *testing.T) {
	builder := &conf.PathBuilder{}
	expected := "/root/code/share/a.cpp"

	pathList := []string{
		"/root",
		"/code/share",
		"a.cpp",
	}
	for _, item := range pathList {
		builder.Append(item)
	}
	actual := builder.FullPath()
	if expected != actual {
		t.Errorf("build unix path error\nexpected:%s\n  actual:%s", expected, actual)
	}

	builder.Clear()
	pathList = []string{
		"/root/",
		"/code/share/",
		"/a.cpp",
	}
	for _, item := range pathList {
		builder.Append(item)
	}
	actual = builder.FullPath()
	if expected != actual {
		t.Errorf("build unix path error\nexpected:%s\n  actual:%s", expected, actual)
	}

	builder.Clear()
	pathList = []string{
		"/root/",
		"code/share/",
		"a.cpp",
	}
	for _, item := range pathList {
		builder.Append(item)
	}
	actual = builder.FullPath()
	if expected != actual {
		t.Errorf("build unix path error\nexpected:%s\n  actual:%s", expected, actual)
	}

	builder.Clear()
	pathList = []string{
		"/root",
		"code/share",
		"a.cpp",
	}
	for _, item := range pathList {
		builder.Append(item)
	}
	actual = builder.FullPath()
	if expected != actual {
		t.Errorf("build unix path error\nexpected:%s\n  actual:%s", expected, actual)
	}

	builder.Clear()
	pathList = []string{
		"",
		"/root/code/share",
		"a.cpp",
	}
	for _, item := range pathList {
		builder.Append(item)
	}
	actual = builder.FullPath()
	if expected != actual {
		t.Errorf("build unix path error\nexpected:%s\n  actual:%s", expected, actual)
	}

	builder.Clear()
	pathList = []string{
		"",
		"/root/code/share/",
		"/a.cpp",
	}
	for _, item := range pathList {
		builder.Append(item)
	}
	actual = builder.FullPath()
	if expected != actual {
		t.Errorf("build unix path error\nexpected:%s\n  actual:%s", expected, actual)
	}

	builder.Clear()
	pathList = []string{
		"",
		"/root/code/share/",
		"a.cpp",
	}
	for _, item := range pathList {
		builder.Append(item)
	}
	actual = builder.FullPath()
	if expected != actual {
		t.Errorf("build unix path error\nexpected:%s\n  actual:%s", expected, actual)
	}

	builder.Clear()
	pathList = []string{
		"",
		"/root/code/share",
		"/a.cpp",
	}
	for _, item := range pathList {
		builder.Append(item)
	}
	actual = builder.FullPath()
	if expected != actual {
		t.Errorf("build unix path error\nexpected:%s\n  actual:%s", expected, actual)
	}

	builder.Clear()
	pathList = []string{
		"/root/code/share",
		"",
		"a.cpp",
	}
	for _, item := range pathList {
		builder.Append(item)
	}
	actual = builder.FullPath()
	if expected != actual {
		t.Errorf("build unix path error\nexpected:%s\n  actual:%s", expected, actual)
	}

	builder.Clear()
	pathList = []string{
		"/root/code/share",
		"/",
		"a.cpp",
	}
	for _, item := range pathList {
		builder.Append(item)
	}
	actual = builder.FullPath()
	if expected != actual {
		t.Errorf("build unix path error\nexpected:%s\n  actual:%s", expected, actual)
	}
}

// TestPathBuilderAppendWindows 测试Windows风格路径的构造
func TestPathBuilderAppendWindows(t *testing.T) {
	builder := &conf.PathBuilder{}
	expected := "d:\\root/code/share/a.cpp"

	pathList := []string{
		"d:\\root",
		"/code/share",
		"a.cpp",
	}
	for _, item := range pathList {
		builder.Append(item)
	}
	actual := builder.FullPath()
	if expected != actual {
		t.Errorf("build unix path error\nexpected:%s\n  actual:%s", expected, actual)
	}

	expected = "d:\\root\\code/share/a.cpp"
	builder.Clear()
	pathList = []string{
		"d:\\root\\",
		"/code/share/",
		"/a.cpp",
	}
	for _, item := range pathList {
		builder.Append(item)
	}
	actual = builder.FullPath()
	if expected != actual {
		t.Errorf("build unix path error\nexpected:%s\n  actual:%s", expected, actual)
	}

	builder.Clear()
	pathList = []string{
		"d:\\root\\",
		"code/share/",
		"a.cpp",
	}
	for _, item := range pathList {
		builder.Append(item)
	}
	actual = builder.FullPath()
	if expected != actual {
		t.Errorf("build unix path error\nexpected:%s\n  actual:%s", expected, actual)
	}

	expected = "d:\\root/code/share/a.cpp"
	builder.Clear()
	pathList = []string{
		"d:\\root",
		"code/share",
		"a.cpp",
	}
	for _, item := range pathList {
		builder.Append(item)
	}
	actual = builder.FullPath()
	if expected != actual {
		t.Errorf("build unix path error\nexpected:%s\n  actual:%s", expected, actual)
	}

	builder.Clear()
	pathList = []string{
		"",
		"d:\\root/code/share",
		"a.cpp",
	}
	for _, item := range pathList {
		builder.Append(item)
	}
	actual = builder.FullPath()
	if expected != actual {
		t.Errorf("build unix path error\nexpected:%s\n  actual:%s", expected, actual)
	}

	builder.Clear()
	pathList = []string{
		"",
		"d:\\root/code/share/",
		"/a.cpp",
	}
	for _, item := range pathList {
		builder.Append(item)
	}
	actual = builder.FullPath()
	if expected != actual {
		t.Errorf("build unix path error\nexpected:%s\n  actual:%s", expected, actual)
	}

	builder.Clear()
	pathList = []string{
		"",
		"d:\\root/code/share/",
		"a.cpp",
	}
	for _, item := range pathList {
		builder.Append(item)
	}
	actual = builder.FullPath()
	if expected != actual {
		t.Errorf("build unix path error\nexpected:%s\n  actual:%s", expected, actual)
	}

	builder.Clear()
	pathList = []string{
		"",
		"d:\\root/code/share",
		"/a.cpp",
	}
	for _, item := range pathList {
		builder.Append(item)
	}
	actual = builder.FullPath()
	if expected != actual {
		t.Errorf("build unix path error\nexpected:%s\n  actual:%s", expected, actual)
	}

	builder.Clear()
	pathList = []string{
		"d:\\root/code/share",
		"",
		"a.cpp",
	}
	for _, item := range pathList {
		builder.Append(item)
	}
	actual = builder.FullPath()
	if expected != actual {
		t.Errorf("build unix path error\nexpected:%s\n  actual:%s", expected, actual)
	}

	builder.Clear()
	pathList = []string{
		"d:\\root/code/share",
		"/",
		"a.cpp",
	}
	for _, item := range pathList {
		builder.Append(item)
	}
	actual = builder.FullPath()
	if expected != actual {
		t.Errorf("build unix path error\nexpected:%s\n  actual:%s", expected, actual)
	}
}

// TestPathBuilderListCall 测试链式调用
func TestPathBuilderListCall(t *testing.T) {
	expected := "/root/code/share/a.cpp"
	builder := conf.PathBuilder{}
	builder.Append("/home/user")
	pathList := []string{
		"/root",
		"code/share",
		"a.cpp",
	}
	actual := builder.Clear().Append(pathList[0]).Append(pathList[1]).Append(pathList[2]).FullPath()
	if expected != actual {
		t.Errorf("list call failed\nexpected:%s\n  actual:%s", expected, actual)
	}

	expected = "/root/code/share/work/"
	pathList = []string{
		"/root",
		"code/share",
		"work/",
	}
	actual = builder.Clear().Append(pathList[0]).Append(pathList[1]).Append(pathList[2]).FullPath()
	if expected != actual {
		t.Errorf("list call failed\nexpected:%s\n  actual:%s", expected, actual)
	}
}
