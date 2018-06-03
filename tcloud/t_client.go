package tcloud

import (
	"net/http"
	"fmt"
	"io/ioutil"
	"encoding/json"
	"io"
	"strings"
	"reflect"
	"strconv"
	"net/url"
	"sort"
	"time"
)

type client struct {
	request
	key string
}

func (client *client) genParams(paramStruct interface{}) []byte {
	baseRequest := &request{client.AppID, time.Now().Unix(), GetRandString(10)}

	allParams := []interface{}{baseRequest, paramStruct}
	paramsLen := len(allParams)
	finalParamsSlice := make([]string, 0, paramsLen+1+3)

	for i := 0; i < paramsLen; i++ {
		tmpStructValue := reflect.ValueOf(allParams[i]).Elem()
		tmpStructLen := tmpStructValue.NumField()

		for k := 0; k < tmpStructLen; k ++ {
			key := tmpStructValue.Type().Field(k).Tag.Get("json")
			value := ""

			switch tmpStructValue.Field(k).Kind() {
			case reflect.Int64, reflect.Int, reflect.Int32:
				value = strconv.FormatInt(tmpStructValue.Field(k).Int(), 10)
			case reflect.Slice:
				value = string(tmpStructValue.Field(k).Bytes())
			default:
				value = tmpStructValue.Field(k).String()
			}
			finalParamsSlice = append(finalParamsSlice, fmt.Sprintf("%s=%s", key, url.QueryEscape(value)))
		}
	}
	sortedParams := sort.StringSlice(finalParamsSlice)
	sortedParams.Sort()
	sortedParams = append(sortedParams, fmt.Sprintf("%s=%s", "app_key", client.key))
	reqParamsStr := strings.Join(sortedParams, "&")
	sign := strings.ToUpper(MD5(reqParamsStr))

	reqParamsStr += fmt.Sprintf("&%s=%s", "sign", sign)
	return []byte(reqParamsStr)
}

//func (client *client) genParams(paramStruct interface{}) []byte {
//	params := paramStruct.(*struct {
//		request
//		FaceFusionRequest
//	})
//	params.NonceStr = model.GetRandString(10)
//	params.Timestamp = time.Now().Unix()
//
//	paramsV := reflect.ValueOf(params).Elem()
//	paramsLen := paramsV.NumField()
//	paramsSlice := make([]string, 0, paramsLen+1)
//
//	for i := 0; i < paramsLen; i++ {
//		t := paramsV.Field(i).NumField()
//		for k := 0; k < t; k ++ {
//			key := paramsV.Field(i).Type().Field(k).Tag.Get("json")
//			value := ""
//			switch paramsV.Field(i).Field(k).Kind() {
//			case reflect.Int64, reflect.Int, reflect.Int32:
//				value = strconv.FormatInt(paramsV.Field(i).Field(k).Int(), 10)
//			case reflect.Slice:
//				value = string(paramsV.Field(i).Field(k).Bytes())
//			default:
//				value = paramsV.Field(i).Field(k).String()
//			}
//			paramsSlice = append(paramsSlice, fmt.Sprintf("%s=%s", key, url.QueryEscape(value)))
//		}
//	}
//	sortedParams := sort.StringSlice(paramsSlice)
//	sortedParams.Sort()
//	sortedParams = append(sortedParams, fmt.Sprintf("%s=%s", "app_key", client.key))
//	reqParamsStr := strings.Join(sortedParams, "&")
//
//	sign := strings.ToUpper(model.MD5(reqParamsStr))
//
//	reqParamsStr += fmt.Sprintf("&%s=%s", "sign", sign)
//	return []byte(reqParamsStr)
//}

func (client *client) sendRequest(urlStr string, data io.Reader, res interface{}) (err error) {
	resp, err := httpClient.Post(urlStr, "application/x-www-form-urlencoded", data)
	if err != nil {
		err = fmt.Errorf("call.failed, err: %v", err)
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		err = fmt.Errorf("unusual response, code: %d", resp.StatusCode)
		return
	}

	resData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		err = fmt.Errorf("read.res.failed, err: %v", err)
	}

	err = json.Unmarshal(resData, res)
	if err != nil {
		err = fmt.Errorf("json.parse.failed, data: %s, err: %v", resData, err)
	}

	return
}

type request struct {
	AppID     int    `json:"app_id"`
	Timestamp int64  `json:"time_stamp"`
	NonceStr  string `json:"nonce_str"`
}

func NewClient(appId int, key string, coonTimeout int, transportTimeout int, maxCoonPerHost int, maxIdleCoons int) *client {
	initClient(coonTimeout, transportTimeout, maxCoonPerHost, maxIdleCoons)
	return &client{request{AppID: appId}, key}
}
