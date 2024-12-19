package server

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/vison888/go-vkit/errorsx/neterrors"
	"github.com/vison888/go-vkit/gate"
	"github.com/vison888/go-vkit/logger"
	"github.com/vison888/logcollector/app"
	"github.com/vison888/logcollector/handler"
)

func tokenCheckFunc(w http.ResponseWriter, r *http.Request) error {
	tokenStr := r.Header.Get("Authorization")
	if tokenStr == "" {
		return neterrors.Unauthorized("没有请求令牌，请求失败!")
	}
	if tokenStr != "ad8045ec-37a3-075b-1f83-53a6ebcae9c1" {
		return neterrors.Unauthorized("令牌错误，请求失败!")
	}
	return nil
}

func logFunc(f gate.HandlerFunc) gate.HandlerFunc {
	return func(ctx context.Context, req *gate.HttpRequest, resp *gate.HttpResponse) error {
		startTime := time.Now()
		err := f(ctx, req, resp)
		costTime := time.Since(startTime)
		body, _, _ := req.Read()
		var logText string
		if err != nil {
			logText = fmt.Sprintf("fail cost:[%v] url:[%v] req:[%v] resp:[%v]", costTime.Milliseconds(), req.Uri(), string(body), err.Error())
		} else {
			logText = fmt.Sprintf("success cost:[%v] url:[%v] req:[%v] resp:[%v]", costTime.Milliseconds(), req.Uri(), string(body), string(resp.Content()))
		}
		logger.Infof(logText)
		return err
	}
}

func Start() {
	h := gate.NewNativeHandler(
		gate.HttpAuthHandler(tokenCheckFunc),
		gate.HttpWrapHandler(logFunc),
	)
	err := h.RegisterApiEndpoint(handler.GetList(), handler.GetApiEndpoint())
	if err != nil {
		logger.Errorf("[main] RegisterApiEndpoint fail %s", err)
		panic(err)
	}
	http.HandleFunc("/api/", func(w http.ResponseWriter, r *http.Request) {
		h.Handle(w, r)
	})

	logger.Infof("[main] Listen port:%d", app.Cfg.Server.HttpPort)
	err = http.ListenAndServe(fmt.Sprintf(":%d", app.Cfg.Server.HttpPort), nil)
	if err != nil {
		logger.Errorf("[main] ListenAndServe fail %s", err)
		panic(err)
	}
}
