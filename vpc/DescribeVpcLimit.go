// 2017-11-13 08:19:07.372602 +0800 CST m=+22.812442771
// Code generated by go generate; DO NOT EDIT.
// This file was generated by robots

package vpc

import (
	"encoding/json"
	"github.com/tencentcloudplatform/tcpicli/core"
)

type DescribeVpcLimitResp struct {
	Code     int64  `json:"code"`
	CodeDesc string `json:"codeDesc"`
	Data     struct {
		Limit struct {
			One int64 `json:"1"`
		} `json:"limit"`
	} `json:"data"`
	Message string `json:"message"`
}

// Implement https://cloud.tencent.com/document/api/215/1844
func DescribeVpcLimit(options ...string) (*DescribeVpcLimitResp, error) {
	resp, err := DoAction("DescribeVpcLimit", options...)
	if err != nil {
		return nil, err
	}
	var s DescribeVpcLimitResp
	err = json.Unmarshal(resp, &s)
	return &s, err
}

func (r *DescribeVpcLimitResp) String(args ...interface{}) (string, error) {
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
