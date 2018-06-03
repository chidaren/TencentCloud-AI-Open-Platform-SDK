package tcloud

import (
	"bytes"
)

type FilterResponse struct {
	Code int                                   `json:"ret"`
	Msg  string                                `json:"msg"`
	Data struct{ Image string `json:"image"` } `json:"data"`
}

type FilterRequest struct {
	Filter int    `json:"filter"`
	Image  []byte `json:"image"`
}

func (client *client) Filter(req *FilterRequest) (res *FilterResponse, err error) {
	params := client.genParams(req)
	res = &FilterResponse{}
	err = client.sendRequest("https://api.ai.qq.com/fcgi-bin/ptu/ptu_imgfilter", bytes.NewReader(params), res)
	return
}
