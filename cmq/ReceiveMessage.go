// 2018-02-06 12:11:24.62771 -0800 PST m=+3.231362046
// Code generated by go generate; DO NOT EDIT.
// This file was generated by robots

package cmq

import (
	"encoding/json"
	"github.com/tencentcloudplatform/tcpicli/core"
)

type ReceiveMessageResp struct {
	ClientRequestID      int64  `json:"clientRequestId"`
	Code                 int64  `json:"code"`
	DequeueCount         int64  `json:"dequeueCount"`
	EnqueueTime          int64  `json:"enqueueTime"`
	FirstDequeueTime     int64  `json:"firstDequeueTime"`
	Message              string `json:"message"`
	MsgBody              string `json:"msgBody"`
	MsgID                string `json:"msgId"`
	NextVisibleTime      int64  `json:"nextVisibleTime"`
	OriginalDeQueueCount int64  `json:"originalDeQueueCount"`
	OriginalEnQueueTime  int64  `json:"originalEnQueueTime"`
	OriginalQueueID      string `json:"originalQueueId"`
	ReceiptHandle        string `json:"receiptHandle"`
	RequestID            string `json:"requestId"`
}

// Implement https://cloud.tencent.com/document/api/406/5839
func (c *CmqClient) ReceiveMessage(options ...string) (*ReceiveMessageResp, error) {
	resp, err := c.DoAction("ReceiveMessage", options...)
	if err != nil {
		return nil, err
	}
	var s ReceiveMessageResp
	err = json.Unmarshal(resp, &s)
	return &s, err
}

func ReceiveMessage(options ...string) (*ReceiveMessageResp, error) {
	return DefaultClient.ReceiveMessage(options...)
}

func (r *ReceiveMessageResp) String(args ...interface{}) (string, error) {
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