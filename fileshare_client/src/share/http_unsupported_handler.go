package share

import "fileshare_client/src/conf"

type HttpUnsupportedHandler struct {
	Next FileShare
}

func (obj HttpUnsupportedHandler) Work(_ conf.ShareConf) bool {
	return false
}
