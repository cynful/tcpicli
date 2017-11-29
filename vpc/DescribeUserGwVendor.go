// 2017-11-27 18:40:00.070603 -0800 PST m=+55.693368986
// Code generated by go generate; DO NOT EDIT.
// This file was generated by robots

package vpc

import (
	"encoding/json"
	"github.com/tencentcloudplatform/tcpicli/core"
)

type DescribeUserGwVendorResp struct {
	Code     int64  `json:"code"`
	CodeDesc string `json:"codeDesc"`
	Data     []struct {
		Platform   string `json:"platform"`
		Software   string `json:"software"`
		Vendorname string `json:"vendorname"`
	} `json:"data"`
	Message string `json:"message"`
}

// Implement https://cloud.tencent.com/document/api/215/5120
func DescribeUserGwVendor(options ...string) (*DescribeUserGwVendorResp, error) {
	resp, err := DoAction("DescribeUserGwVendor", options...)
	if err != nil {
		return nil, err
	}
	var s DescribeUserGwVendorResp
	err = json.Unmarshal(resp, &s)
	return &s, err
}

func (r *DescribeUserGwVendorResp) String(args ...interface{}) (string, error) {
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
