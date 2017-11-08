// 2017-11-07 17:23:49.262937947 +0800 CST m=+0.575089703
// Code generated by go generate; DO NOT EDIT.
// This file was generated by robots

package ccs

import (
	"encoding/json"
	"github.com/tencentcloudplatform/tcpicli/core"
)

type CreateClusterServiceResp struct {
	Code     int64       `json:"code"`
	CodeDesc string      `json:"codeDesc"`
	Data     interface{} `json:"data"`
	Message  string      `json:"message"`
}

// Implement https://cloud.tencent.com/document/api/457/9436
func CreateClusterService(options ...string) (*CreateClusterServiceResp, error) {
	resp, err := DoAction("CreateClusterService", options...)
	if err != nil {
		return nil, err
	}
	var s CreateClusterServiceResp
	err = json.Unmarshal(resp, &s)
	return &s, err
}

func (r *CreateClusterServiceResp) String(args ...interface{}) (string, error) {
	var b []byte
	var err error
	if len(args) == 0 {
		b, err = json.MarshalIndent(r, "", "  ")
	} else if len(args) == 1 {
		if val, ok := args[0].(string); ok {
			if len(val) == 0 {
				b, err = json.MarshalIndent(r, "", "  ")
			} else {
				b, err = core.FmtOutput(val, r)
			}
		}
	}
	return string(b), err
}
