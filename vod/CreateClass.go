// 2017-12-01 10:54:06.763181 -0800 PST m=+1.481380632
// Code generated by go generate; DO NOT EDIT.
// This file was generated by robots

package vod

import (
	"encoding/json"
	"github.com/tencentcloudplatform/tcpicli/core"
)

type CreateClassResp struct {
	Code       int64  `json:"code"`
	CodeDesc   string `json:"codeDesc"`
	Message    string `json:"message"`
	NewClassID string `json:"newClassId"`
}

// Implement https://cloud.tencent.com/document/api/266/7812
func CreateClass(options ...string) (*CreateClassResp, error) {
	resp, err := DoAction("CreateClass", options...)
	if err != nil {
		return nil, err
	}
	var s CreateClassResp
	err = json.Unmarshal(resp, &s)
	return &s, err
}

func (r *CreateClassResp) String(args ...interface{}) (string, error) {
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