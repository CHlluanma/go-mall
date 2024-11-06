# go-mall
使用cloudwego的hertz和kitex搭建的gomall

学习自[gomall](https://github.com/cloudwego/biz-demo/tree/main/gomall)

## 技术栈
- `Kitex`
- `Hertz`
- `MySQL`
- `Redis`
- `Docker`
- `Consul`
- `OpenTelemetry`
- `cwgo`
- `NATS`
- `GORM`
- `Prometheus`
- `K8s`

## 工具安装
- `cwgo` 

```bash
# Go 1.15 及之前版本
GO111MODULE=on GOPROXY=https://goproxy.cn/,direct go get github.com/cloudwego/cwgo@latest

# Go 1.16 及以后版本
GOPROXY=https://goproxy.cn/,direct go install github.com/cloudwego/cwgo@latest
```