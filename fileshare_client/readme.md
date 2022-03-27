# 命令行使用介绍

## 下载

### 下载文件

获取一个远程文件并打印到标准输出
```
fs http://host/path/remote_file
```

获取一个远程文件，保存到本地，并指定本地文件名
```
fs http://host/path/remote_file /local_path/local_file
```

获取一个远程文件，保存到本地，指定本地目录名，本地文件名与远程文件名保持一致。注意本地路径必须以路径分隔符结尾。若没有最后的路径分隔符，
则当作本地文件处理
```
fs http://host/path/remote_file /local_path/
```

### 列出远程目录下的内容

本接口只列出远程目录下第一层的子文件和子目录，不会递归获取更下一层的内容。
```
fs http://host/path/remote_path
```

### 下载远程目录

以非递归的方式下载远程目录
```
fs http://host/path/remote_path /path/local_path
```

以递归的方式下载远程目录
```
fs http://host/path/remote_path /path/local_path -r
```

## 上传

上传本地文件，指定远程文件名，且此接口必须指定远程文件名。
```
fs /path/local_file http://host/path/remote_file
```

以非递归的方式上传本地目录
```
fs /path/local_path http://host/path/remote_path
```

以递归方式上传本地目录
```
fs /path/local_path http://host/path/remote_path -r
```

# 配置文件使用介绍

配置文件示例

```json
{
    "host": "http://127.0.0.1:9087",
    "root_dir": "E:\\WindowsLib\\Desktop\\client",
    "share": [
        {
            "op": "get",
            "remote": "/read/r1/read_r1.txt",
            "local": ""
        },
        {
            "op": "get",
            "remote": "/read/r1/read_r1.txt",
            "local": "/download/r1/read_r1.txt"
        },
        {
            "op": "ls",
            "remote": "/read",
            "local": ""
        },
        {
            "op": "get",
            "remote": "/read",
            "local": "/download/download_dir"
        },
        {
            "op": "get",
            "remote": "/read",
            "local": "/download/download_dir_recursive"
        },
        {
            "op": "put",
            "remote": "/write/w2/upload.txt",
            "local": "/upload.txt"
        },
        {
            "op": "put",
            "remote": "/write/w2",
            "local": "/upload"
        },
        {
            "op": "put",
            "remote": "/write/w2",
            "local": "/upload",
            "recursive": true
        }
    ]
}
```

## 上传

### 上传文件

remote说明
* 远程文件
* 不存在则创建
* 存在则覆盖
* 不可是目录

local说明
* 本地文件

配置文件示例
```json
{
    "host": "http://127.0.0.1:9087",
    "root_dir": "E:\\WindowsLib\\Desktop\\client",
    "share": [
        {
            "op": "put",
            "remote": "/write/w2/upload.txt",
            "local": "/upload.txt"
        }
    ]
}
```

### 上传目录

remote说明
* 远程目录
* 不存在则创建
* 存在则覆盖
* 不可是文件

local说明
* 本地目录

远程文件的实际存放路径为
${RemoteRootPath}/${Remote}/${LocalRelativePath}

若$.share.local以路径分隔符结尾，则上传的相对目录不包含local目录，否则包含
* local值为/upload/，远程文件为/$remote_path/upload/relative_path/local_file
* local值为/upload，远程文件为/$remote_path/relative_path/local_file
* relative_path是相对于$root_dir/$local的相对目录
* 只上传指定目录下的文件，目录不会上传

配置文件示例
```json
{
    "host": "http://127.0.0.1:9087",
    "root_dir": "E:\\WindowsLib\\Desktop\\client",
    "share": [
        {
            "op": "put",
            "remote": "/write/w2",
            "local": "/upload"
        }
    ]
}
```

若要递归上传本地目录，加入recursive字段即可，示例如下
```json
{
    "host": "http://127.0.0.1:9087",
    "root_dir": "E:\\WindowsLib\\Desktop\\client",
    "share": [
        {
            "op": "put",
            "remote": "/write/w2",
            "local": "/upload",
            "recursive": true
        }
    ]
}
```
