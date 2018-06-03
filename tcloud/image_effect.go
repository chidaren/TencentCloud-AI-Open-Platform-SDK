package tcloud

import (
	"bytes"
	"fmt"
)

type FaceFusionResponse struct {
	Code int                                   `json:"ret"`
	Msg  string                                `json:"msg"`
	Data struct{Image string `json:"image"`} `json:"data"`
}

type FaceFusionRequest struct {
	Model int `json:"model"`
	Image []byte `json:"image"`
}

func (client *client) FaceFusion(req *FaceFusionRequest) (res *FaceFusionResponse, err error) {
	params := client.genParams(req)
	fmt.Println(string(params))

	res = &FaceFusionResponse{}
	err = client.sendRequest(bytes.NewReader(params), res)
	return
}
