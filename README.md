# 说明

## 本地测试

```shell
go run main.go
```

## gin 框架文档

https://gin-gonic.com/zh-cn/docs/examples/

## VsCode & Go

https://learn.microsoft.com/zh-cn/azure/developer/go/configure-visual-studio-code

## 设置国内镜像

查看镜像配置

```
$ go env | grep GOPROXY
GOPROXY="https://goproxy.cn"
```

### 七牛 CDN
```
go env -w  GOPROXY=https://goproxy.cn,direct
```

### 阿里云
```
go env -w GOPROXY=https://mirrors.aliyun.com/goproxy/,direct
```

### 官方
```
go env -w  GOPROXY=https://goproxy.io,direct
```