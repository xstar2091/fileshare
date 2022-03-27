package main

import (
	"fileshare_client/src/conf"
	"fileshare_client/src/share"
	"fileshare_client/src/utils"
	"os"
)

func main() {
	clientConf, ok := conf.CreateNewConf()
	if !ok {
		os.Exit(1)
	}
	ret := true
	fs := share.CreateFileShare()
	for _, item := range clientConf.ShareConfList {
		utils.LogDebug("share info:%s", item.ToString())
		if !fs.Work(item) {
			ret = false
		}
	}
	if !ret {
		os.Exit(2)
	}
}
