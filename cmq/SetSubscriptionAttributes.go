// 2018-02-06 18:00:06.053066 -0800 PST m=+30.172247524
// Code generated by go generate; DO NOT EDIT.
// This file was generated by robots

package cmq

import (
	"encoding/json"
	"github.com/tencentcloudplatform/tcpicli/core"
)

type SetSubscriptionAttributesResp struct {
	Code      int64  `json:"code"`
	Message   string `json:"message"`
	RequestID string `json:"requestId"`
}

// Implement https://cloud.tencent.com/document/api/406/7416
func (c *CmqClient) SetSubscriptionAttributes(options ...string) (*SetSubscriptionAttributesResp, error) {
	resp, err := c.DoAction("SetSubscriptionAttributes", options...)
	if err != nil {
		return nil, err
	}
	var s SetSubscriptionAttributesResp
	err = json.Unmarshal(resp, &s)
	return &s, err
}

func SetSubscriptionAttributes(options ...string) (*SetSubscriptionAttributesResp, error) {
	return DefaultClient.SetSubscriptionAttributes(options...)
}

func (r *SetSubscriptionAttributesResp) String(args ...interface{}) (string, error) {
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
