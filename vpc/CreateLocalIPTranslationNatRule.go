// 2017-11-28 11:44:01.218285 -0800 PST m=+260.185719305
// Code generated by go generate; DO NOT EDIT.
// This file was generated by robots

package vpc

import (
	"encoding/json"
	"github.com/tencentcloudplatform/tcpicli/core"
)

type CreateLocalIPTranslationNatRuleResp struct {
	Code     int64  `json:"code"`
	CodeDesc string `json:"codeDesc"`
	Data     struct {
		TaskID int64 `json:"taskId"`
	} `json:"data"`
	Message string `json:"message"`
}

// Implement https://cloud.tencent.com/document/api/215/5185
func CreateLocalIPTranslationNatRule(options ...string) (*CreateLocalIPTranslationNatRuleResp, error) {
	resp, err := DoAction("CreateLocalIPTranslationNatRule", options...)
	if err != nil {
		return nil, err
	}
	var s CreateLocalIPTranslationNatRuleResp
	err = json.Unmarshal(resp, &s)
	return &s, err
}

func (r *CreateLocalIPTranslationNatRuleResp) String(args ...interface{}) (string, error) {
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
