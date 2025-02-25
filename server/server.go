package server

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
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

func alert(content string) {
	content = strings.ReplaceAll(content, "\\tat", "&nbsp;&nbsp;&nbsp;&nbsp;")
	content = strings.ReplaceAll(content, "\\n", "<br/>")

	// Create the alert payload
	alertData := map[string]interface{}{
		"title":   "Android出现宕机",
		"content": content,
		"time":    time.Now().Format(time.RFC3339),
	}

	logger.Infof("content=:%s", content)
	jsonData, err := json.Marshal(alertData)
	if err != nil {
		logger.Errorf("Failed to marshal alert data: %v", err)
		return
	}

	// Create the HTTP request
	req, err := http.NewRequest("POST", app.Cfg.Alert.AlertUrl, bytes.NewBuffer(jsonData))
	if err != nil {
		logger.Errorf("Failed to create request: %v", err)
		return
	}

	// Set headers
	req.Header.Set("Content-Type", "application/json")

	// Create HTTP client with timeout
	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	// Send the request
	resp, err := client.Do(req)
	if err != nil {
		logger.Errorf("Failed to send alert: %v", err)
		return
	}
	defer resp.Body.Close()

	// Check response status
	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		logger.Errorf("Alert failed with status %d: %s", resp.StatusCode, string(body))
		return
	}

	logger.Infof("Alert sent successfully")
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

	http.HandleFunc("/api/collector/log.crash", func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if e := r.Body.Close(); e != nil {
				return
			}
		}()
		bytes, err := io.ReadAll(r.Body)
		if err != nil {
			logger.Errorf("[main] err %s", err.Error())
			return
		}

		go alert(string(bytes))

		netErr := &neterrors.NetError{
			Msg:  "OK",
			Code: 0,
		}
		content, _ := json.Marshal(netErr)

		logger.Infof("read android crash msg bytes:%s", string(bytes))
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.WriteHeader(200)
		w.Header().Set("Content-Length", strconv.Itoa(len(content)))
		w.Write(content)
	})

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
