// 2018-02-08 13:05:42.021037 -0800 PST m=+15.484416013
// Code generated by go generate; DO NOT EDIT.
// This file was generated by robots

package cmq

import (
	"encoding/json"
	"github.com/tencentcloudplatform/tcpicli/core"
)

type BatchReceiveMessageResp struct {
	ClientRequestID float64 `json:"clientRequestId"`
	Code            float64 `json:"code"`
	Message         string  `json:"message"`
	MsgInfoList     []struct {
		DequeueCount         float64 `json:"dequeueCount"`
		EnqueueTime          float64 `json:"enqueueTime"`
		FirstDequeueTime     float64 `json:"firstDequeueTime"`
		MsgBody              string  `json:"msgBody"`
		MsgID                string  `json:"msgId"`
		NextVisibleTime      float64 `json:"nextVisibleTime"`
		OriginalDeQueueCount float64 `json:"originalDeQueueCount"`
		OriginalEnQueueTime  float64 `json:"originalEnQueueTime"`
		OriginalQueueID      string  `json:"originalQueueId"`
		ReceiptHandle        string  `json:"receiptHandle"`
	} `json:"msgInfoList"`
	RequestID string `json:"requestId"`
}

// Implement https://cloud.tencent.com/document/api/406/5924
func (c *CmqClient) BatchReceiveMessage(options ...string) (*BatchReceiveMessageResp, error) {
	resp, err := c.DoAction("BatchReceiveMessage", options...)
	if err != nil {
		return nil, err
	}
	var s BatchReceiveMessageResp
	err = json.Unmarshal(resp, &s)
	return &s, err
}

func BatchReceiveMessage(options ...string) (*BatchReceiveMessageResp, error) {
	return DefaultClient.BatchReceiveMessage(options...)
}

func (r *BatchReceiveMessageResp) String(args ...interface{}) (string, error) {
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
