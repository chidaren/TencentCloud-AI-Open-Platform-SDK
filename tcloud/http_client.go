package tcloud

import (
	"net/http"
	"net"
	"time"
)


var httpClient *http.Client

func initClient(coonTimeout int, transportTimeout int, maxCoonPerHost int, maxIdleCoons int) {
	transport := &http.Transport{
		Dial: func(netP, address string) (net.Conn, error) {
			conn, err := net.DialTimeout(netP, address, time.Duration(coonTimeout) * time.Second)
			if err != nil {
				return nil, err
			}
			return conn, nil
		},
		MaxIdleConns: maxIdleCoons,
		MaxIdleConnsPerHost: maxCoonPerHost,
	}
	httpClient = &http.Client{Transport: transport, Timeout: time.Duration(transportTimeout) * time.Second}
}
