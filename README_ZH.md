# lifecycle

[![PkgGoDev](https://pkg.go.dev/badge/github.com/qmdx00/lifecycle)](https://pkg.go.dev/github.com/qmdx00/lifecycle)
![license](https://img.shields.io/github/license/qmdx00/lifecycle)
[![Go Report Card](https://goreportcard.com/badge/github.com/qmdx00/lifecycle)](https://goreportcard.com/report/github.com/qmdx00/lifecycle)
[![codecov](https://codecov.io/gh/qmdx00/lifecycle/branch/master/graph/badge.svg?token=MVJ5OIUYSK)](https://codecov.io/gh/qmdx00/lifecycle)
[![Build Status](https://app.travis-ci.com/qmdx00/lifecycle.svg?branch=master)](https://app.travis-ci.com/qmdx00/lifecycle)

[English](./README.md) | ä¸­æ

## ð ä»ç»

ä¸ä¸ªç®åçåºç¨çå½å¨æç®¡çå·¥å·ï¼æ¹ä¾¿æ¥å¥å¤ä¸ªæå¡å®ä¾ã

## ð ç¹æ§

- æ¯ææè½½å¤ä¸ª Serverã
- æ¹ä¾¿åºç¨åæå¡ä¹é´ä¼ éåæ°æ®ã
- ä¼éå°å¤çæå¡å¯å¨åç»æ­¢ã
- æä¾å¯æå±çæ¸çå½æ°é©å­ã

## ð§° å®è£

```
go get -u github.com/qmdx00/lifecycle
```

## ð  ä½¿ç¨

ç®åå°å®ç° `Server` interfaceï¼å¹¶æè½½å°åºç¨ä¸­å³å¯ã

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
    app.Cleanup(func() error {
        log.Println("do cleanup")
        return nil
    })

    if err := app.Run(); err != nil {
        log.Fatal(err)
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
    log.Printf("server %s start\n", info.Name())
    return e.srv.ListenAndServe()
}

func (e *EchoServer) Stop(ctx context.Context) error {
    info, _ := lifecycle.FromContext(ctx)
    log.Printf("server %s stop\n", info.Name())
    return e.srv.Shutdown(ctx)
}
```

## ð è®¸å¯è¯

Â© Wimi Yuan, 2021~time.Now <br>
Released under the [MIT License](./LICENSE).
