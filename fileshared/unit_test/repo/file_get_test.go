package repo

import (
	"fileshared/src/handler"
	"net/http"
	"net/http/httptest"
	"testing"
)

// TestFileGetFileExist 文件存在
func TestFileGetFileExist(t *testing.T) {
	mux := http.NewServeMux()
	mux.Handle("/read/read1.txt", &handler.ServerHandler{})

	// 获取目录信息，路径不以分隔符结尾
	r, _ := http.NewRequest(http.MethodGet, "/read/read1.txt", nil)
	w := httptest.NewRecorder()
	mux.ServeHTTP(w, r)
	resp := w.Result()

	expectedStatusCode := http.StatusOK
	if resp.StatusCode != expectedStatusCode {
		t.Errorf("test status_code failed, actual:%d, expected:%d", resp.StatusCode, expectedStatusCode)
	}
	expectedContentType := "application/text"
	contentType := resp.Header.Get("Content-Type")
	if contentType != expectedContentType {
		t.Errorf("test header[\"Content-Type\"] failed, actual:%s, expected:%s", contentType, expectedContentType)
	}
}

// TestFileGetFileNotExist 文件不存在
func TestFileGetFileNotExist(t *testing.T) {
	mux := http.NewServeMux()
	mux.Handle("/read/not_exist_file.txt", &handler.ServerHandler{})

	// 获取目录信息，路径不以分隔符结尾
	r, _ := http.NewRequest(http.MethodGet, "/read/not_exist_file.txt", nil)
	w := httptest.NewRecorder()
	mux.ServeHTTP(w, r)
	resp := w.Result()

	expectedStatusCode := http.StatusNotFound
	if resp.StatusCode != expectedStatusCode {
		t.Errorf("test status_code failed, actual:%d, expected:%d", resp.StatusCode, expectedStatusCode)
	}
}

// TestFileGetFileEndWithSeparator 文件存在，但以路径分隔符结尾，当作目录处理，返回404
func TestFileGetFileEndWithSeparator(t *testing.T) {
	mux := http.NewServeMux()
	mux.Handle("/read/not_exist_file.txt", &handler.ServerHandler{})

	// 获取目录信息，路径不以分隔符结尾
	r, _ := http.NewRequest(http.MethodGet, "/read/not_exist_file.txt/", nil)
	w := httptest.NewRecorder()
	mux.ServeHTTP(w, r)
	resp := w.Result()

	expectedStatusCode := http.StatusNotFound
	if resp.StatusCode != expectedStatusCode {
		t.Errorf("test status_code failed, actual:%d, expected:%d", resp.StatusCode, expectedStatusCode)
	}
}

// TestFileGetNoPermission 文件存在，但没有读权限
func TestFileGetNoPermission(t *testing.T) {
	mux := http.NewServeMux()
	mux.Handle("/write/write1.txt", &handler.ServerHandler{})

	// 获取目录信息，路径不以分隔符结尾
	r, _ := http.NewRequest(http.MethodGet, "/write/write1.txt", nil)
	w := httptest.NewRecorder()
	mux.ServeHTTP(w, r)
	resp := w.Result()

	expectedStatusCode := http.StatusForbidden
	if resp.StatusCode != expectedStatusCode {
		t.Errorf("test status_code failed, actual:%d, expected:%d", resp.StatusCode, expectedStatusCode)
	}
}

// TestFileGetRecursive 文件存在，递归读取
func TestFileGetRecursive(t *testing.T) {
	mux := http.NewServeMux()
	mux.Handle("/read/r1/read_r1.txt", &handler.ServerHandler{})

	// 获取目录信息，路径不以分隔符结尾
	r, _ := http.NewRequest(http.MethodGet, "/read/r1/read_r1.txt", nil)
	w := httptest.NewRecorder()
	mux.ServeHTTP(w, r)
	resp := w.Result()

	expectedStatusCode := http.StatusOK
	if resp.StatusCode != expectedStatusCode {
		t.Errorf("test status_code failed, actual:%d, expected:%d", resp.StatusCode, expectedStatusCode)
	}
	expectedContentType := "application/text"
	contentType := resp.Header.Get("Content-Type")
	if contentType != expectedContentType {
		t.Errorf("test header[\"Content-Type\"] failed, actual:%s, expected:%s", contentType, expectedContentType)
	}
}

// TestFileGetNoRecursive 文件存在，但不允许递归读取
func TestFileGetNoRecursive(t *testing.T) {
	mux := http.NewServeMux()
	mux.Handle("/read_norecursive/r1/read_norecursive_r1.txt", &handler.ServerHandler{})

	// 获取目录信息，路径不以分隔符结尾
	r, _ := http.NewRequest(http.MethodGet, "/read_norecursive/r1/read_norecursive_r1.txt", nil)
	w := httptest.NewRecorder()
	mux.ServeHTTP(w, r)
	resp := w.Result()

	expectedStatusCode := http.StatusForbidden
	if resp.StatusCode != expectedStatusCode {
		t.Errorf("test status_code failed, actual:%d, expected:%d", resp.StatusCode, expectedStatusCode)
	}
}
