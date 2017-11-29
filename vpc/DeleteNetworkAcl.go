// 2017-11-17 13:40:47.647154 -0800 PST m=+31.426731586
// Code generated by go generate; DO NOT EDIT.
// This file was generated by robots

package vpc

import (
	"encoding/json"
	"github.com/tencentcloudplatform/tcpicli/core"
)

type DeleteNetworkAclResp struct {
	Code     int64  `json:"code"`
	CodeDesc string `json:"codeDesc"`
	Message  string `json:"message"`
}

// Implement https://cloud.tencent.com/document/api/215/1439
func DeleteNetworkAcl(options ...string) (*DeleteNetworkAclResp, error) {
	resp, err := DoAction("DeleteNetworkAcl", options...)
	if err != nil {
		return nil, err
	}
	var s DeleteNetworkAclResp
	err = json.Unmarshal(resp, &s)
	return &s, err
}

func (r *DeleteNetworkAclResp) String(args ...interface{}) (string, error) {
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
