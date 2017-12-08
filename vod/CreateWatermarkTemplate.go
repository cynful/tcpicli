// 2017-12-01 14:30:16.059846 -0800 PST m=+13.181386388
// Code generated by go generate; DO NOT EDIT.
// This file was generated by robots

package vod

import (
	"encoding/json"
	"github.com/tencentcloudplatform/tcpicli/core"
)

type CreateWatermarkTemplateResp struct {
	Code       int64  `json:"code"`
	CodeDesc   string `json:"codeDesc"`
	Definition int64  `json:"definition"`
	Message    string `json:"message"`
}

// Implement https://cloud.tencent.com/document/api/266/11599
func CreateWatermarkTemplate(options ...string) (*CreateWatermarkTemplateResp, error) {
	resp, err := DoAction("CreateWatermarkTemplate", options...)
	if err != nil {
		return nil, err
	}
	var s CreateWatermarkTemplateResp
	err = json.Unmarshal(resp, &s)
	return &s, err
}

func (r *CreateWatermarkTemplateResp) String(args ...interface{}) (string, error) {
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