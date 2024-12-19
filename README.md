# log-collector

日志收集器

## 1 window golang环境
```
下载并解压 https://go.dev/dl/go1.23.4.windows-386.zip
gowork自建目录
go为解压目录

配置环境变量
GOPATH：D:\gowork
GOROOT D:\go
%GOPATH%\bin
%GOROOT%\bin
```

## 2 linux golang环境
```
wget https://go.dev/dl/go1.23.4.linux-amd64.tar.gz
tar -zxvf go1.23.4.linux-amd64.tar.gz 
export PATH=$PATH:/home/ubuntu/software/go/bin
export GOPATH=$HOME/go
export PATH=$PATH:$GOPATH/bin
```

## 3 window 开发环境
### 安装参数校验工具
```
go install github.com/envoyproxy/protoc-gen-validate@v1.0.2

```


### 安装API文档生成工具
```
go install github.com/grpc-ecosystem/grpc-gateway/protoc-gen-openapiv2@v2.18.0

```

### 安装vkit grpc接口生成工具
```
go install github.com/vison888/protoc-gen-vkit@master

```

### 安装protoc
```
https://github.com/protocolbuffers/protobuf/releases/download/v3.20.3/protoc-3.20.3-win64.zip

```

### 安装protoc go
```
https://pkg.go.dev/google.golang.org/protobuf@v1.28.0
https://github.com/protocolbuffers/protobuf-go/releases/tag/v1.28.0
```
