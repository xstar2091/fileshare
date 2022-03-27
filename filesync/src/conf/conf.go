package conf

type Conf struct {
    ConfFile   string
    RecordFile string
    UploadFile string
    Local LocalConf `json:"local"`
    Remote RemoteConf `json:"remote"`
}

type LocalConf struct {
    Path string `json:"path"`
    Exclude []string `json:"exclude"`
}

type RemoteConf struct {
    Host string `json:"host"`
}
