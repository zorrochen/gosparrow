package proxy

import (
	"bytes"
	"context"
	"encoding/json"
	"gosparrow/base/log"
	"gosparrow/base/pub"
	"io/ioutil"
	"net"
	"net/http"
	"time"
)

func Get(url string) ([]byte, int) {
	return reqProxy(url, "GET", nil, nil, 0)
}

func PostJson(url string, req interface{}) ([]byte, int) {
	reqjson, _ := json.Marshal(req)
	return reqProxy(url, "POST", nil, reqjson, 0)
}

// ================= 初始化 =================
// 长连接的httpclient
var httpClient *http.Client

// 长连接配置
func init() {
	httpClient = http.DefaultClient
	httpClient.Transport = &http.Transport{
		Proxy: http.ProxyFromEnvironment,
		Dial: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
		}).Dial,
		TLSHandshakeTimeout: 10 * time.Second,
		DisableCompression:  true,
	}
}

// ================= req =================
func reqProxy(url string, method string, header map[string]string, data []byte, timeout int64) ([]byte, int) {
	// 请求初始化
	request, err := http.NewRequest(method, url, bytes.NewReader(data))
	if err != nil {
		return nil, -1
	}

	// 添加header
	for k, v := range header {
		request.Header.Add(k, v)
	}

	// 默认超时时间：3s
	if timeout == 0 {
		timeout = 3
	}

	// 添加超时
	ctx, _ := context.WithTimeout(context.TODO(), time.Duration(timeout)*time.Second)
	request = request.WithContext(ctx)

	// 添加请求打印
	log.Log.Debug("[proxy-req] %s, %v, %s", request.URL.String(), request.Header, pub.CopyHttpRequestBody(request))

	// 发送请求
	resp, err := httpClient.Do(request)
	if err != nil {
		return nil, -1
	}
	defer resp.Body.Close()

	// 读取响应body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, -1
	}

	// 添加返回打印
	log.Log.Debug("[proxy-return] %s, %s", request.URL.String(), string(body))

	// return
	return body, resp.StatusCode
}
