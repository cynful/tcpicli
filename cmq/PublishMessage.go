// 2018-02-06 18:00:04.356667 -0800 PST m=+28.475846370
// Code generated by go generate; DO NOT EDIT.
// This file was generated by robots

package cmq

import (
	"encoding/json"
	"github.com/tencentcloudplatform/tcpicli/core"
)

type PublishMessageResp struct {
	Code      int64  `json:"code"`
	Message   string `json:"message"`
	MsgID     string `json:"msgId"`
	RequestID string `json:"requestId"`
}

// Implement https://cloud.tencent.com/document/api/406/7411
func (c *CmqClient) PublishMessage(options ...string) (*PublishMessageResp, error) {
	resp, err := c.DoAction("PublishMessage", options...)
	if err != nil {
		return nil, err
	}
	var s PublishMessageResp
	err = json.Unmarshal(resp, &s)
	return &s, err
}

func PublishMessage(options ...string) (*PublishMessageResp, error) {
	return DefaultClient.PublishMessage(options...)
}

func (r *PublishMessageResp) String(args ...interface{}) (string, error) {
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
