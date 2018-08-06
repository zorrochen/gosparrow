package proxy

import (
	"encoding/json"
	"errors"
	"fmt"
	"gosparrow/proxy"
	"net/http"
)

//================= Ccv2CircleInfo =================
type Ccv2CircleInfoReq struct {
	City  string `json:"city"`
	Count int64  `json:"count"`
	Page  int64  `json:"page"`
}

type Ccv2CircleInfoResp struct {
	ErrorCode int64  `json:"error_code"`
	ErrorMsg  string `json:"error_msg"`
	TestHyx   string `json:"test_hyx"`
}

//圈子信息
func Ccv2CircleInfo(req Ccv2CircleInfoReq) (*Ccv2CircleInfoResp, error) {
	rst := &Ccv2CircleInfoResp{}

	//request
	url := fmt.Sprintf("http://%s/ccv2/circle/info?%s", ReqAddr, proxy.Struct2Querystr(req))
	httpcode, body := proxy.Get(url)
	if httpcode != http.StatusOK {
		return nil, errors.New("request failed.")
	}

	//response
	err := json.Unmarshal(body, rst)
	if err != nil {
		return nil, err
	}

	return rst, nil
}

//================= circleBaseInfo =================
type circleBaseInfoReq struct {
	Content  string `json:"content"`
	FileLine string `json:"fileLine"`
	SrvName  string `json:"srvName"`
}

type circleBaseInfoResp struct {
	ErrCode string `json:"err_code"`
	ErrMsg  string `json:"err_msg"`
}

//圈子基础信息
func circleBaseInfo(req circleBaseInfoReq) (*circleBaseInfoResp, error) {
	rst := &circleBaseInfoResp{}

	//request
	url := fmt.Sprintf("http://%s/ccv2/circle/baseinfo", ReqAddr)
	httpcode, body := proxy.PostJson(url, req)
	if httpcode != http.StatusOK {
		return nil, errors.New("request failed.")
	}

	//response
	err := json.Unmarshal(body, rst)
	if err != nil {
		return nil, err
	}

	return rst, nil
}

//================= mycircles =================
type mycirclesReq struct {
	A int64 `json:"a"`
	B int64 `json:"b"`
}

type mycirclesResp struct {
	C int64    `json:"c"`
	D int64    `json:"d"`
	F []string `json:"f"`
}

//我的圈子信息
func mycircles(req mycirclesReq) (*mycirclesResp, error) {
	rst := &mycirclesResp{}

	//request
	url := fmt.Sprintf("http://%s/ccv2/circle/mycircles?%s", ReqAddr, proxy.Struct2Querystr(req))
	httpcode, body := proxy.Get(url)
	if httpcode != http.StatusOK {
		return nil, errors.New("request failed.")
	}

	//response
	err := json.Unmarshal(body, rst)
	if err != nil {
		return nil, err
	}

	return rst, nil
}

//================= circleAllInfo =================
type circleAllInfoReq struct {
	A int64 `json:"a"`
	B int64 `json:"b"`
}

type circleAllInfoResp struct {
	C int64    `json:"c"`
	D int64    `json:"d"`
	F []string `json:"f"`
}

//信息组装
func circleAllInfo(req circleAllInfoReq) (*circleAllInfoResp, error) {
	rst := &circleAllInfoResp{}

	//request
	url := fmt.Sprintf("http://%s/ccv2/circle/allinfo?%s", ReqAddr, proxy.Struct2Querystr(req))
	httpcode, body := proxy.Get(url)
	if httpcode != http.StatusOK {
		return nil, errors.New("request failed.")
	}

	//response
	err := json.Unmarshal(body, rst)
	if err != nil {
		return nil, err
	}

	return rst, nil
}
