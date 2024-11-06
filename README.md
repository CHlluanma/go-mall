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

## 编码指南
- uber go style guide: https://github.com/uber-go/guide
- protobuf style guide: https://protobuf.dev/programming-guides/style/
- MDN HTTP response status codes：https://developer.mozilla.org/zh-CN/docs/Web/HTTP/Status
- 约定式提交：https://www.conventionalcommits.org/zh-hans/v1.0.0/
- 语义化版本：https://semver.org/lang/zh-CN/