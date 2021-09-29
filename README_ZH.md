# lifecycle

[![PkgGoDev](https://pkg.go.dev/badge/github.com/qmdx00/lifecycle)](https://pkg.go.dev/github.com/qmdx00/lifecycle)
![license](https://img.shields.io/github/license/qmdx00/lifecycle)
[![Go Report Card](https://goreportcard.com/badge/github.com/qmdx00/lifecycle)](https://goreportcard.com/report/github.com/qmdx00/lifecycle)
[![codecov](https://codecov.io/gh/qmdx00/lifecycle/branch/master/graph/badge.svg?token=MVJ5OIUYSK)](https://codecov.io/gh/qmdx00/lifecycle)
[![Build Status](https://app.travis-ci.com/qmdx00/lifecycle.svg?branch=master)](https://app.travis-ci.com/qmdx00/lifecycle)

[English](./README.md) | 中文

## 📖 介绍

一个简单的应用生命周期管理工具，方便接入多个服务实例。

## 🚀 特性

- 支持挂载多个 Server。
- 方便应用和服务之间传递元数据。
- 优雅地处理服务启动和终止。
- 提供可拓展的清理函数钩子。

## 🧰 安装

```
go get -u github.com/qmdx00/lifecycle
```

## 🛠 使用

简单地实现 `Server` interface，并挂载到应用中即可。

```go
package main

import (
    "context"
    "github.com/qmdx00/lifecycle"
    "log"
    "net/http"
)

func main() {
    app := lifecycle.NewApp(
        lifecycle.WithName("test"),
        lifecycle.WithVersion("v1.0"),
    )

    app.Attach("echo", NewEchoServer())
    app.Cleanup(func() {
        log.Println("do cleanup")
    })

    if err := app.Run(); err != nil {
        panic(err)
    }
}

func NewEchoServer() lifecycle.Server {
    handler := http.NewServeMux()
    handler.HandleFunc("/echo", func(writer http.ResponseWriter, request *http.Request) {
        _, _ = writer.Write([]byte("hello world"))
    })

    return &EchoServer{
        srv: &http.Server{
            Addr:    ":3000",
            Handler: handler,
        },
    }
}

type EchoServer struct {
    srv *http.Server
}

func (e *EchoServer) Run(ctx context.Context) error {
    info, _ := lifecycle.FromContext(ctx)
    log.Println(info.Name(), info.Version())
    return e.srv.ListenAndServe()
}

func (e *EchoServer) Stop(ctx context.Context) error {
    return e.srv.Shutdown(ctx)
}
```

## 📄 许可证

© Wimi Yuan, 2021~time.Now <br>
Released under the [MIT License](./LICENSE).
