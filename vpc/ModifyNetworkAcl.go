// 2017-11-17 13:49:02.322879 -0800 PST m=+33.515043097
// Code generated by go generate; DO NOT EDIT.
// This file was generated by robots

package vpc

import (
	"encoding/json"
	"github.com/tencentcloudplatform/tcpicli/core"
)

type ModifyNetworkAclResp struct {
	Code     int64  `json:"code"`
	CodeDesc string `json:"codeDesc"`
	Message  string `json:"message"`
}

// Implement https://cloud.tencent.com/document/api/215/1443
func ModifyNetworkAcl(options ...string) (*ModifyNetworkAclResp, error) {
	resp, err := DoAction("ModifyNetworkAcl", options...)
	if err != nil {
		return nil, err
	}
	var s ModifyNetworkAclResp
	err = json.Unmarshal(resp, &s)
	return &s, err
}

func (r *ModifyNetworkAclResp) String(args ...interface{}) (string, error) {
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
