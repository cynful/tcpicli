// 2017-11-13 08:12:07.976283 +0800 CST m=+19.630839998
// Code generated by go generate; DO NOT EDIT.
// This file was generated by robots

package vpc

import (
	"encoding/json"
	"github.com/tencentcloudplatform/tcpicli/core"
)

type DescribeVpcClassicLinkResp struct {
	Code       int64  `json:"code"`
	CodeDesc   string `json:"codeDesc"`
	Message    string `json:"message"`
	TotalCount int64  `json:"totalCount"`
}

// Implement https://cloud.tencent.com/document/api/215/1844
func DescribeVpcClassicLink(options ...string) (*DescribeVpcClassicLinkResp, error) {
	resp, err := DoAction("DescribeVpcClassicLink", options...)
	if err != nil {
		return nil, err
	}
	var s DescribeVpcClassicLinkResp
	err = json.Unmarshal(resp, &s)
	return &s, err
}

func (r *DescribeVpcClassicLinkResp) String(args ...interface{}) (string, error) {
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
