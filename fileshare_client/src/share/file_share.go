package share

import "fileshare_client/src/conf"

type FileShare interface {
	Work(shareConf conf.ShareConf) bool
}

func CreateFileShare() FileShare {
	unsupported := HttpUnsupportedHandler{
		Next: nil,
	}
	httpGet := HttpGetHandler{
		Next: unsupported,
	}
	httpPutDirRecursive := HttpPutDirectoryRecursiveHandler{
		Next: httpGet,
	}
	httpPutDir := HttpPutDirectoryHandler{
		Next: httpPutDirRecursive,
	}
	httpPutFile := HttpPutFileHandler{
		Next: httpPutDir,
	}
	return httpPutFile
}
