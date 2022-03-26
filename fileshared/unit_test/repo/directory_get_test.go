package repo

import (
	"encoding/json"
	"fileshared/src/handler"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

type PathInfo struct {
	IsDir      bool   `json:"is_dir"`
	Mode       uint32 `json:"mode"`
	ModifyTime string `json:"modify_time"`
	Name       string `json:"name"`
	Size       int64  `json:"size"`
}

type DirectoryGetResponse struct {
	Path string     `json:"path"`
	Info []PathInfo `json:"info"`
}

// TestDirectoryGet 目录存在，url path不以分隔符结尾
func TestDirectoryGetPathNotEndWithSeparator(t *testing.T) {
	mux := http.NewServeMux()
	mux.Handle("/read", &handler.ServerHandler{})

	// 获取目录信息，路径不以分隔符结尾
	r, _ := http.NewRequest(http.MethodGet, "/read", nil)
	w := httptest.NewRecorder()
	mux.ServeHTTP(w, r)
	resp := w.Result()

	expectedStatusCode := http.StatusOK
	if resp.StatusCode != expectedStatusCode {
		t.Errorf("test status_code failed, actual:%d, expected:%d", resp.StatusCode, expectedStatusCode)
	}
	expectedContentType := "application/json"
	contentType := resp.Header.Get("Content-Type")
	if contentType != expectedContentType {
		t.Errorf("test header[\"Content-Type\"] failed, actual:%s, expected:%s", contentType, expectedContentType)
	}
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Errorf("read response body failed, error:%s", err.Error())
	}
	jsonObj := DirectoryGetResponse{}
	err = json.Unmarshal(data, &jsonObj)
	if err != nil {
		t.Errorf("parse response body from json failed, error:%s", err.Error())
	}
}

// TestDirectoryGet 目录存在，url path以分隔符结尾
func TestDirectoryGetPathEndWithSeparator(t *testing.T) {
	mux := http.NewServeMux()
	mux.Handle("/read/", &handler.ServerHandler{})

	// 获取目录信息，路径不以分隔符结尾
	r, _ := http.NewRequest(http.MethodGet, "/read/", nil)
	w := httptest.NewRecorder()
	mux.ServeHTTP(w, r)
	resp := w.Result()

	expectedStatusCode := http.StatusOK
	if resp.StatusCode != expectedStatusCode {
		t.Errorf("test status_code failed, actual:%d, expected:%d", resp.StatusCode, expectedStatusCode)
	}
	expectedContentType := "application/json"
	contentType := resp.Header.Get("Content-Type")
	if contentType != expectedContentType {
		t.Errorf("test header[\"Content-Type\"] failed, actual:%s, expected:%s", contentType, expectedContentType)
	}
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Errorf("read response body failed, error:%s", err.Error())
	}
	jsonObj := DirectoryGetResponse{}
	err = json.Unmarshal(data, &jsonObj)
	if err != nil {
		t.Errorf("parse response body from json failed, error:%s", err.Error())
	}
}

// TestDirectoryGet 空目录
func TestDirectoryGetEmptyPath(t *testing.T) {
	mux := http.NewServeMux()
	mux.Handle("/read/r2", &handler.ServerHandler{})

	// 获取目录信息，路径不以分隔符结尾
	r, _ := http.NewRequest(http.MethodGet, "/read/r2", nil)
	w := httptest.NewRecorder()
	mux.ServeHTTP(w, r)
	resp := w.Result()

	expectedStatusCode := http.StatusOK
	if resp.StatusCode != expectedStatusCode {
		t.Errorf("test status_code failed, actual:%d, expected:%d", resp.StatusCode, expectedStatusCode)
	}
	expectedContentType := "application/json"
	contentType := resp.Header.Get("Content-Type")
	if contentType != expectedContentType {
		t.Errorf("test header[\"Content-Type\"] failed, actual:%s, expected:%s", contentType, expectedContentType)
	}
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Errorf("read response body failed, error:%s", err.Error())
	}
	jsonObj := DirectoryGetResponse{}
	err = json.Unmarshal(data, &jsonObj)
	if err != nil {
		t.Errorf("parse response body from json failed, error:%s", err.Error())
	}
}

// TestDirectoryGet 目录不存在
func TestDirectoryGetPathNotExist(t *testing.T) {
	mux := http.NewServeMux()
	mux.Handle("/read/not_exist_directory", &handler.ServerHandler{})

	// 获取目录信息，路径不以分隔符结尾
	r, _ := http.NewRequest(http.MethodGet, "/read/not_exist_directory", nil)
	w := httptest.NewRecorder()
	mux.ServeHTTP(w, r)
	resp := w.Result()

	expectedStatusCode := http.StatusNotFound
	if resp.StatusCode != expectedStatusCode {
		t.Errorf("test status_code failed, actual:%d, expected:%d", resp.StatusCode, expectedStatusCode)
	}
}

// TestDirectoryGet 目录无读权限
func TestDirectoryGetPathNoPermission(t *testing.T) {
	mux := http.NewServeMux()
	mux.Handle("/write", &handler.ServerHandler{})

	// 获取目录信息，路径不以分隔符结尾
	r, _ := http.NewRequest(http.MethodGet, "/write", nil)
	w := httptest.NewRecorder()
	mux.ServeHTTP(w, r)
	resp := w.Result()

	expectedStatusCode := http.StatusForbidden
	if resp.StatusCode != expectedStatusCode {
		t.Errorf("test status_code failed, actual:%d, expected:%d", resp.StatusCode, expectedStatusCode)
	}
}

// TestDirectoryGet 递归读取目录
func TestDirectoryGetPathRecursive(t *testing.T) {
	mux := http.NewServeMux()
	mux.Handle("/read/r1", &handler.ServerHandler{})

	// 获取目录信息，路径不以分隔符结尾
	r, _ := http.NewRequest(http.MethodGet, "/read/r1", nil)
	w := httptest.NewRecorder()
	mux.ServeHTTP(w, r)
	resp := w.Result()

	expectedStatusCode := http.StatusOK
	if resp.StatusCode != expectedStatusCode {
		t.Errorf("test status_code failed, actual:%d, expected:%d", resp.StatusCode, expectedStatusCode)
	}
	expectedContentType := "application/json"
	contentType := resp.Header.Get("Content-Type")
	if contentType != expectedContentType {
		t.Errorf("test header[\"Content-Type\"] failed, actual:%s, expected:%s", contentType, expectedContentType)
	}
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Errorf("read response body failed, error:%s", err.Error())
	}
	jsonObj := DirectoryGetResponse{}
	err = json.Unmarshal(data, &jsonObj)
	if err != nil {
		t.Errorf("parse response body from json failed, error:%s", err.Error())
	}
}

// TestDirectoryGet 目录不允许递归读取
func TestDirectoryGetPathNoRecursive(t *testing.T) {
	mux := http.NewServeMux()
	mux.Handle("/read_norecursive/r1", &handler.ServerHandler{})

	// 获取目录信息，路径不以分隔符结尾
	r, _ := http.NewRequest(http.MethodGet, "/read_norecursive/r1", nil)
	w := httptest.NewRecorder()
	mux.ServeHTTP(w, r)
	resp := w.Result()

	expectedStatusCode := http.StatusForbidden
	if resp.StatusCode != expectedStatusCode {
		t.Errorf("test status_code failed, actual:%d, expected:%d", resp.StatusCode, expectedStatusCode)
	}
}
