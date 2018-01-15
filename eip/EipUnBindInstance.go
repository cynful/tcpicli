// 2018-01-09 09:16:36.48717 -0800 PST m=+32.017954046
// Code generated by go generate; DO NOT EDIT.
// This file was generated by robots

package eip

import (
	"encoding/json"
	"github.com/tencentcloudplatform/tcpicli/core"
)

type EipUnBindInstanceResp struct {
	Code     int64  `json:"code"`
	CodeDesc string `json:"codeDesc"`
	Data     struct {
		RequestID int64 `json:"requestId"`
	} `json:"data"`
	Message string `json:"message"`
}

// Implement https://cloud.tencent.com/document/api/213/1376
func (c *EipClient) EipUnBindInstance(options ...string) (*EipUnBindInstanceResp, error) {
	resp, err := c.DoAction("EipUnBindInstance", options...)
	if err != nil {
		return nil, err
	}
	var s EipUnBindInstanceResp
	err = json.Unmarshal(resp, &s)
	return &s, err
}

func EipUnBindInstance(options ...string) (*EipUnBindInstanceResp, error) {
	return DefaultClient.EipUnBindInstance(options...)
}

func (r *EipUnBindInstanceResp) String(args ...interface{}) (string, error) {
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
