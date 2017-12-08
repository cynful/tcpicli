// 2017-12-06 18:39:48.291933 -0800 PST m=+11.274237606
// Code generated by go generate; DO NOT EDIT.
// This file was generated by robots

package cvm

import (
	"encoding/json"
	"github.com/tencentcloudplatform/tcpicli/core"
)

type InquiryPriceResetInstanceResp struct {
	Response struct {
		Price struct {
			InstancePrice struct {
				ChargeUnit string  `json:"ChargeUnit"`
				UnitPrice  float64 `json:"UnitPrice"`
			} `json:"InstancePrice"`
		} `json:"Price"`
		Error     interface{} `json:"Error,omitempty"`
		RequestID string      `json:"RequestId"`
	} `json:"Response"`
}

// Implement https://cloud.tencent.com/document/api/213/9490
func InquiryPriceResetInstance(options ...string) (*InquiryPriceResetInstanceResp, error) {
	resp, err := DoAction("InquiryPriceResetInstance", options...)
	if err != nil {
		return nil, err
	}
	var s InquiryPriceResetInstanceResp
	err = json.Unmarshal(resp, &s)
	return &s, err
}

func (r *InquiryPriceResetInstanceResp) String(args ...interface{}) (string, error) {
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