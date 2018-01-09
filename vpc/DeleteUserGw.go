// 2018-01-09 09:23:39.545038 -0800 PST m=+100.870890169
// Code generated by go generate; DO NOT EDIT.
// This file was generated by robots

package vpc

import (
	"encoding/json"
	"github.com/tencentcloudplatform/tcpicli/core"
)

type DeleteUserGwResp struct {
	Code     int64  `json:"code"`
	CodeDesc string `json:"codeDesc"`
	Message  string `json:"message"`
}

// Implement https://cloud.tencent.com/document/api/215/5117
func (c *VpcClient) DeleteUserGw(options ...string) (*DeleteUserGwResp, error) {
	resp, err := c.DoAction("DeleteUserGw", options...)
	if err != nil {
		return nil, err
	}
	var s DeleteUserGwResp
	err = json.Unmarshal(resp, &s)
	return &s, err
}

func DeleteUserGw(options ...string) (*DeleteUserGwResp, error) {
	return DefaultClient.DeleteUserGw(options...)
}

func (r *DeleteUserGwResp) String(args ...interface{}) (string, error) {
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
