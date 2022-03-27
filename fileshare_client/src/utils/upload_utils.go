package utils

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

func UploadFile(localFile string, remoteFile string) bool {
	var err error = nil
	defer func() {
		if err != nil {
			LogError("failed, error:%s", err.Error())
		}
	}()

	data, err := ioutil.ReadFile(localFile)
	if err != nil {
		return false
	}
	req, err := http.NewRequest(http.MethodPut, remoteFile, bytes.NewReader(data))
	if err != nil {
		return false
	}
	req.Header.Set("Content-Type", "application/text")
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return false
	}
	LogDebug("response status code %d", resp.StatusCode)
	if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
		err = fmt.Errorf("faield, remote response %d", resp.StatusCode)
		return false
	}

	return true
}
