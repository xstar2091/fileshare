package repo

import (
	"fileshared/src/handler"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// TestFilePutFileNotExist 文件不存在
func TestFilePutFileNotExist(t *testing.T) {
	uploadFileName := fmt.Sprintf("%s_1.txt", UnitTestTempFileNamePrefix)
	urlPath := fmt.Sprintf("/write/%s", uploadFileName)
	mux := http.NewServeMux()
	mux.Handle(urlPath, &handler.ServerHandler{})

	reader := strings.NewReader("TestFilePutFileNotExist")
	r, _ := http.NewRequest(http.MethodPut, urlPath, reader)
	w := httptest.NewRecorder()
	mux.ServeHTTP(w, r)
	resp := w.Result()

	expectedStatusCode := http.StatusCreated
	if resp.StatusCode != expectedStatusCode {
		t.Errorf("test status_code failed, actual:%d, expected:%d", resp.StatusCode, expectedStatusCode)
	}
}

// TestFilePutFileExist 文件存在
func TestFilePutFileExist(t *testing.T) {
	uploadFileName := fmt.Sprintf("%s_1.txt", UnitTestTempFileNamePrefix)
	urlPath := fmt.Sprintf("/write/%s", uploadFileName)
	mux := http.NewServeMux()
	mux.Handle(urlPath, &handler.ServerHandler{})

	reader := strings.NewReader("TestFilePutFileExist")
	r, _ := http.NewRequest(http.MethodPut, urlPath, reader)
	w := httptest.NewRecorder()
	mux.ServeHTTP(w, r)
	resp := w.Result()

	expectedStatusCode := http.StatusOK
	if resp.StatusCode != expectedStatusCode {
		t.Errorf("test status_code failed, actual:%d, expected:%d", resp.StatusCode, expectedStatusCode)
	}
}

// TestFilePutEndWithSeparator 文件存在，但以路径分隔符结尾
func TestFilePutEndWithSeparator(t *testing.T) {
	uploadFileName := fmt.Sprintf("%s_1.txt", UnitTestTempFileNamePrefix)
	urlPath := fmt.Sprintf("/write/%s/", uploadFileName)
	mux := http.NewServeMux()
	mux.Handle(urlPath, &handler.ServerHandler{})

	reader := strings.NewReader("TestFilePutEndWithSeparator")
	r, _ := http.NewRequest(http.MethodPut, urlPath, reader)
	w := httptest.NewRecorder()
	mux.ServeHTTP(w, r)
	resp := w.Result()

	expectedStatusCode := http.StatusBadRequest
	if resp.StatusCode != expectedStatusCode {
		t.Errorf("test status_code failed, actual:%d, expected:%d", resp.StatusCode, expectedStatusCode)
	}
}

// TestFilePutFileNotExistRecursive 文件不存在，递归
func TestFilePutFileNotExistRecursive(t *testing.T) {
	uploadFileName := fmt.Sprintf("%s_1.txt", UnitTestTempFileNamePrefix)
	urlPath := fmt.Sprintf("/write/w2/%s", uploadFileName)
	mux := http.NewServeMux()
	mux.Handle(urlPath, &handler.ServerHandler{})

	reader := strings.NewReader("TestFilePutFileNotExistRecursive")
	r, _ := http.NewRequest(http.MethodPut, urlPath, reader)
	w := httptest.NewRecorder()
	mux.ServeHTTP(w, r)
	resp := w.Result()

	expectedStatusCode := http.StatusCreated
	if resp.StatusCode != expectedStatusCode {
		t.Errorf("test status_code failed, actual:%d, expected:%d", resp.StatusCode, expectedStatusCode)
	}
}

// TestFilePutFileExist 文件存在，递归
func TestFilePutFileExistRecursive(t *testing.T) {
	uploadFileName := fmt.Sprintf("%s_1.txt", UnitTestTempFileNamePrefix)
	urlPath := fmt.Sprintf("/write/w2/%s", uploadFileName)
	mux := http.NewServeMux()
	mux.Handle(urlPath, &handler.ServerHandler{})

	reader := strings.NewReader("TestFilePutFileExistRecursive")
	r, _ := http.NewRequest(http.MethodPut, urlPath, reader)
	w := httptest.NewRecorder()
	mux.ServeHTTP(w, r)
	resp := w.Result()

	expectedStatusCode := http.StatusOK
	if resp.StatusCode != expectedStatusCode {
		t.Errorf("test status_code failed, actual:%d, expected:%d", resp.StatusCode, expectedStatusCode)
	}
}

// TestFilePutFileNotExistNoRecursive 文件不存在，不允许递归
func TestFilePutFileNotExistNoRecursive(t *testing.T) {
	uploadFileName := fmt.Sprintf("%s_1.txt", UnitTestTempFileNamePrefix)
	urlPath := fmt.Sprintf("/write_norecursive/w2/%s", uploadFileName)
	mux := http.NewServeMux()
	mux.Handle(urlPath, &handler.ServerHandler{})

	reader := strings.NewReader("TestFilePutFileNotExistNoRecursive")
	r, _ := http.NewRequest(http.MethodPut, urlPath, reader)
	w := httptest.NewRecorder()
	mux.ServeHTTP(w, r)
	resp := w.Result()

	expectedStatusCode := http.StatusForbidden
	if resp.StatusCode != expectedStatusCode {
		t.Errorf("test status_code failed, actual:%d, expected:%d", resp.StatusCode, expectedStatusCode)
	}
}

// TestFilePutFileNotExistNoRecursive 文件存在，不允许递归
func TestFilePutFileExistNoRecursive(t *testing.T) {
	urlPath := "/write_norecursive/w1/write_norecursive1.txt"
	mux := http.NewServeMux()
	mux.Handle(urlPath, &handler.ServerHandler{})

	reader := strings.NewReader("TestFilePutFileExistNoRecursive")
	r, _ := http.NewRequest(http.MethodPut, urlPath, reader)
	w := httptest.NewRecorder()
	mux.ServeHTTP(w, r)
	resp := w.Result()

	expectedStatusCode := http.StatusForbidden
	if resp.StatusCode != expectedStatusCode {
		t.Errorf("test status_code failed, actual:%d, expected:%d", resp.StatusCode, expectedStatusCode)
	}
}
