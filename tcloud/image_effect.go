package tcloud

import (
	"bytes"
)

type FaceFusionResponse struct {
	Code int                                   `json:"ret"`
	Msg  string                                `json:"msg"`
	Data struct{ Image string `json:"image"` } `json:"data"`
}

type FaceFusionRequest struct {
	Model int    `json:"model"`
	Image []byte `json:"image"`
}

func (client *client) FaceFusion(req *FaceFusionRequest) (res *FaceFusionResponse, err error) {
	params := client.genParams(req)
	res = &FaceFusionResponse{}
	err = client.sendRequest("https://api.ai.qq.com/fcgi-bin/ptu/ptu_facemerge", bytes.NewReader(params), res)
	return
}
