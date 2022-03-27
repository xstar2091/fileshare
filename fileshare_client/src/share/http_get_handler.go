package share

import "fileshare_client/src/conf"

type HttpGetHandler struct {
	Next FileShare
}

func (obj HttpGetHandler) Work(shareConf conf.ShareConf) bool {
	return false
}
